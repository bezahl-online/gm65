package device

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/snksoft/crc"
	"github.com/tarm/serial"
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
	Baud       int
}

// Scanner is the class
type Scanner struct {
	lock 	sync.RWMutex
	Config  Config
	port    *serial.Port
	command *command
}

// Open comm port to gm65
func (g *Scanner) Open() error {
	c := &serial.Config{Name: g.Config.SerialPort, Baud: g.Config.Baud}
	var err error
	g.port, err = serial.OpenPort(c)
	return err
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

func (g *Scanner) write(c *command) error {
	if g.port==nil {
		return errNotConnected
	}
	g.command = c
	// add header bytes
	g.head()
	// default error message
	var err = fmt.Errorf("open serial port first")
	//check if comm port is opened
	if g.port != nil {
		// write to gm65	
		_, err = g.port.Write(g.crc())
		err = reconnectIfLost(err, g)
	}
	return err
}

func reconnectIfLost(err error, g *Scanner) error {
	if err != nil {
		go Connect(g)
		if err.Error() == "EOF" {
			err = fmt.Errorf("lost connection to scanner device")
		}
	}
	return err
}

var errNotConnected error = fmt.Errorf("scanner device not connected")
// listen to gm65 on comm port with timeout
func (g *Scanner) readWithTimeout(timeout time.Duration) ([]byte, error) {
	if g.port==nil {
		return nil,errNotConnected
	}
	// FIXME: set timeout
	buf := make([]byte, 128)
	var err error
	n, err := g.port.Read(buf)
	err = reconnectIfLost(err, g)
	return buf[:n], err
}

// listen to gm65 on comm port
func (g *Scanner) Read() ([]byte, error) {
	return g.readWithTimeout(0)
}

// readZone reads the data in the given zone
func (g *Scanner) readZone(zone [2]byte) (byte, error) {
	err := g.write(&command{
		Function: read,
		Length:   1,
		Address:  zone,
		Data:     1,
		CRC:      [2]byte{},
	})
	buf, err := g.Read()
	var data byte
	if err == nil && buf != nil && len(buf) == 7 {
		data = buf[4]
	}
	return data, err
}

// writeZoneBit writes single bits into given
// zone via logical OR and leaves other bits intact
func (g *Scanner) writeZoneBit(zone [2]byte, set byte, clear byte) error {
	data, err := g.readZone(zone)
	if err != nil {
		return err
	}
	//fmt.Printf("\nbefore: %08b\n", data)
	data |= set
	//fmt.Printf("\nset:    %08b %08b\n", data, set)
	data &= ^clear
	//fmt.Printf("\nclear:  %08b %08b\n", data, clear)
	return g.writeZoneByte(zone, data)
}

func (g *Scanner) writeZoneByte(zone [2]byte, data byte) error {
	defer g.lock.Unlock()
	g.lock.Lock()
	err := g.write(&command{
		Function: send,
		Length:   1,
		Address:  zone,
		Data:     data,
		CRC:      [2]byte{},
	})
	buf, err := g.Read()
	if err != nil {
		return err
	}
	if buf == nil || len(buf) != 7 {
		return fmt.Errorf("wrong data received")
	}
	return nil
}
