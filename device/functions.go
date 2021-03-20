package device

import (
	"encoding/binary"
	"fmt"
)

// LightOn sets the aim (red light) to ON
func (g *Scanner) LightOn() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x08, 0x04)
}

// LightOff sets the aim (red light) to Off
func (g *Scanner) LightOff() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x00, 0x0c)
}

// LightStd sets the aim (red light) to Std
func (g *Scanner) LightStd() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x04, 0x08)
}

// AimOn sets the aim (red light) to ON
func (g *Scanner) AimOn() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x20, 0x10)
}

// AimOff sets the aim (red light) to OFF
func (g *Scanner) AimOff() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x00, 0x30)
}

// AimStd sets the aim (red light) to STD
func (g *Scanner) AimStd() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x10, 0x20)
}

// ProductModel return product model id
func (g *Scanner) ProductModel() (byte, error) {
	return g._readZone([2]byte{0, 0xe0})
}

// HardwareVersion return hardware version
func (g *Scanner) HardwareVersion() (byte, error) {
	return g._readZone([2]byte{0, 0xe1})
}

// SoftwareVersion return hardware version
func (g *Scanner) SoftwareVersion() (byte, error) {
	return g._readZone([2]byte{0, 0xe2})
}

// SoftwareDate return software date
func (g *Scanner) SoftwareDate() (string, error) {
	year, err := g._readZone([2]byte{0, 0xe3})
	if err == nil {
		month, err := g._readZone([2]byte{0, 0xe4})
		if err == nil {
			day, err := g._readZone([2]byte{0, 0xe5})
			if err == nil {
				return fmt.Sprintf("%4d%02d%02d", int(year)+2000, month, day), nil
			}
		}
	}
	return "", err
}

// ReadInterval sets the read interval
func (g *Scanner) ImageStabilisationTimeout(interval byte) error {
	return g.writeZoneByte([2]byte{0, 0x04}, interval)
}

// ReadSoundFreq sets the successfully read sound frequency
func (g *Scanner) ReadSoundFreq(frequency int16) error {
	if frequency < 1 || frequency > 5100 {
		return fmt.Errorf("frequency need to be between 20 and 5100 Hz")
	}
	var freq byte = byte(frequency / 20)
	return g.writeZoneByte([2]byte{0, 0x0a}, freq)
}

// ReadSoundDuration sets the duration in ms
// of the successfully read sound
func (g *Scanner) ReadSoundDuration(duration byte) error {
	return g.writeZoneByte([2]byte{0, 0x0b}, duration)
}

// AutoSleep sets the light
func (g *Scanner) AutoSleep(on bool, timeMills uint16) error {
	var b []byte = []byte{0, 0}
	binary.BigEndian.PutUint16(b, uint16(timeMills))
	err := g.writeZoneByte([2]byte{0, 0x08}, b[1])
	if err != nil {
		return err
	}
	var autoSleep byte = 0x80
	if on {
		b[0] |= autoSleep
	} else {
		b[0] &= ^autoSleep
	}
	return g.writeZoneByte([2]byte{0, 0x07}, b[0])
}

// SingleReadTime timeout for a single read attempt
func (g *Scanner) SingleReadTime(readTime byte) error {
	return g.writeZoneByte([2]byte{0, 0x06}, readTime)
}

// ReadIntervalTime set time interval of consecutive
// reading attemps
func (g *Scanner) ReadIntervalTime(readTime byte) error {
	return g.writeZoneByte([2]byte{0, 0x05}, readTime)
}

// SensorMode set scanner to sensor mode
func (g *Scanner) SensorMode() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x03, 0x00)
}

// ManualMode set scanner to manual mode
func (g *Scanner) ManualMode() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x00, 0x03)
}

// ContinuousMode set scanner to continuous mode
func (g *Scanner) ContinuousMode() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x02, 0x01)
}

// CommandMode set scanner to command mode
func (g *Scanner) CommandMode() error {
	return g.writeZoneBit([2]byte{0, 0}, 0x01, 0x02)
}

// OpenLEDOnSuccess set scanner to
// Open LED when successfully read
func (g *Scanner) OpenLEDOnSuccess(on bool) error {
	var set, clear byte = 0x80, 0x00
	if !on {
		set = 0
		clear = 0x80
	}
	return g.writeZoneBit([2]byte{0, 0}, set, clear)
}

// Mute set scanner mute
func (g *Scanner) Mute(on bool) error {
	var mute byte = 0x40
	var clear byte = 0
	if on {
		clear = mute
	}
	return g.writeZoneBit([2]byte{0, 0}, mute, clear)
}

// RotateRead360 set if scanner is allowed to scan 360 deg
func (g *Scanner) RotateRead360(on bool) error {
	var set byte = 0x01
	var clear byte = 0
	if !on {
		clear = set
		set = 0
	}
	return g.writeZoneBit([2]byte{0, 0x2c}, set, clear)
}

// ReadSuccessFullSound set if scanner should peep
func (g *Scanner) ReadSuccessFullSound(on bool) error {
	var set byte = 0x04
	var clear byte = 0
	if !on {
		clear = set
		set = 0
	}
	return g.writeZoneBit([2]byte{0, 0x0e}, set, clear)
}

// DisableAllBarcode disable any barcodes
func (g *Scanner) DisableAllBarcode() error {
	return g.writeZoneBit([2]byte{0, 0x2c}, 0x00, 0x06)
}

// EnableEAN13 allow scanner to read EAN13
func (g *Scanner) EnableEAN13() error {
	return g.writeZoneBit([2]byte{0, 0x2e}, 0x01, 0x00)
}

// EnableEAN8 allow scanner to read EAN8
func (g *Scanner) EnableEAN8() error {
	return g.writeZoneBit([2]byte{0, 0x2f}, 0x01, 0x00)
}

// EnableCode39 allow scanner to read Code39
func (g *Scanner) EnableCode39() error {
	return g.writeZoneBit([2]byte{0, 0x36}, 0x01, 0x00)
}

// EnableQRCode allow scanner to read QR codes
func (g *Scanner) EnableQRCode() error {
	return g.writeZoneBit([2]byte{0, 0x3f}, 0x01, 0x00)
}

func (g *Scanner) Configure() {
	fmt.Println("configure")
	g.DisableAllBarcode()
	g.EnableEAN13()
	g.LightOn()
	g.AimOn()
	g.ReadIntervalTime(1)
	g.ReadSuccessFullSound(false)
}
