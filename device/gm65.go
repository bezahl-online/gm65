package device

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/snksoft/crc"

	"go.bug.st/serial"
)

const read byte = 0x07
const send byte = 0x08
const save byte = 0x09

type command struct {
	Head     [2]byte
	Function byte
	Length   byte
	Address  [2]byte
	Data     byte
	CRC      [2]byte
}

// convert command structure to byte array
func (c command) getBytes() []byte {
	var b []byte = []byte{
		c.Head[0],
		c.Head[1],
		c.Function,
		c.Length,
		c.Address[0],
		c.Address[1],
		c.Data,
		c.CRC[0],
		c.CRC[1],
	}
	return b
}

// Config is the GM65 config structure
type Config struct {
	SerialPort string
	Baud       uint
}

type Barcode struct {
	sync.RWMutex
	Value string
}

// Scanner is the class
type Scanner struct {
	lock       *sync.RWMutex
	Config     Config
	port       *serial.Port
	connected  bool
	command    *command
	listening  bool
	latestCode Barcode
}

// Open comm port to gm65
func (g *Scanner) Open() error {
	mode := &serial.Mode{
		BaudRate: int(g.Config.Baud),
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	var err error
	port, err := serial.Open(g.Config.SerialPort, mode)
	g.port = &port
	g.connected = true
	return err

}

// TODO: listen to multiple scanner
func (g *Scanner) Listen() {
	if g.listening {
		return
	}
	g.listening = true
	for {
		if !g.connected {
			g.listening = false
			break
		}
		result, err := g.Read()
		if err == nil {
			g.latestCode.Lock()
			g.latestCode.Value = string(result)
			g.latestCode.Unlock()
		} else {
			g.listening = false
			break
		}
	}
}
func (g *Scanner) SetCode(code string) {
	g.latestCode.Lock()
	defer g.latestCode.Unlock()
	g.latestCode.Value = code
}
func (g *Scanner) GetLatest() string {
	g.latestCode.RLock()
	defer g.latestCode.RUnlock()
	code := g.latestCode.Value
	g.latestCode.Value = ""
	return code
}

// calculate crc and append to command structure
func (g *Scanner) crc() []byte {
	var c *command = (g.command)
	var data []byte = c.getBytes()[2:7]
	crcCode := crc.CalculateCRC(crc.XMODEM, data)
	var b []byte = []byte{0, 0}
	binary.BigEndian.PutUint16(b, uint16(crcCode))
	c.CRC = [2]byte{b[0], b[1]}
	return c.getBytes()
}

// add header bytes to command structure
func (g *Scanner) head() {
	var head [2]byte = [2]byte{0x7e, 0x00}
	(*g.command).Head = head
}

func (g *Scanner) _write(c *command) error {
	if err := g.reconnectIfLost(); err != nil {
		return err
	}
	g.command = c
	// add header bytes
	g.head()
	// default error message
	_, err := (*g.port).Write(g.crc())
	if err != nil {
		return err
	}
	return nil
}

func (g *Scanner) reconnectIfLost() error {
	if !g.connected {
		go Connect(g)
		time.Sleep(100 * time.Millisecond)
		if g.connected {
			// seems to be connected again
			return nil
		}
		return fmt.Errorf("lost connection to scanner device")
	}
	return nil
}

var errNotConnected error = fmt.Errorf("scanner device not connected")

// listen to gm65 on comm port
func (g *Scanner) _read() ([]byte, error) {
	if err := g.reconnectIfLost(); err != nil {
		return nil, err
	}
	buf := make([]byte, 128)
	var err error
	n, err := (*g.port).Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

// ReadTimeout listens to gm65 on comm port
func (g *Scanner) Read() ([]byte, error) {
	defer g.lock.Unlock()
	g.lock.Lock()
	return g._read()
}

func (g *Scanner) readZone(zone [2]byte) (byte, error) {
	defer g.lock.Unlock()
	g.lock.Lock()
	return g._readZone(zone)
}

// _readZone reads the data in the given zone
func (g *Scanner) _readZone(zone [2]byte) (byte, error) {
	err := g._write(&command{
		Function: read,
		Length:   1,
		Address:  zone,
		Data:     1,
		CRC:      [2]byte{},
	})
	buf, err := g._read()
	var data byte
	if err == nil && buf != nil && len(buf) == 7 {
		data = buf[4]
	}
	return data, err
}

// writeZoneBit writes single bits into given
// zone via logical OR and leaves other bits intact
func (g *Scanner) writeZoneBit(zone [2]byte, set byte, clear byte) error {
	defer g.lock.Unlock()
	g.lock.Lock()
	return g._writeZoneBit(zone, set, clear)
}

func (g *Scanner) _writeZoneBit(zone [2]byte, set byte, clear byte) error {
	data, err := g._readZone(zone)
	if err != nil {
		return err
	}
	//fmt.Printf("\nbefore: %08b\n", data)
	data |= set
	//fmt.Printf("\nset:    %08b %08b\n", data, set)
	data &= ^clear
	//fmt.Printf("\nclear:  %08b %08b\n", data, clear)
	return g._writeZoneByte(zone, data)
}

// writeZoneByte writes a byte to the zone
func (g *Scanner) writeZoneByte(zone [2]byte, data byte) error {
	defer g.lock.Unlock()
	g.lock.Lock()
	return g._writeZoneByte(zone, data)
}

func (g *Scanner) _writeZoneByte(zone [2]byte, data byte) error {
	err := g._write(&command{
		Function: send,
		Length:   1,
		Address:  zone,
		Data:     data,
		CRC:      [2]byte{},
	})
	buf, err := g._read()
	if err != nil {
		return err
	}
	if buf == nil || len(buf) != 7 {
		return fmt.Errorf("wrong data received")
	}
	return nil
}
