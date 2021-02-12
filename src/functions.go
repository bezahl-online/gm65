package gm65

import (
	"encoding/binary"
	"fmt"
)

// Light sets the light
func (g *Scanner) Light(on bool, std bool) error {
	var set, clear, st byte = 0x08, 0x0c, 0x04
	var zone [2]byte = [2]byte{0, 0}
	if std {
		clear = set
		set = st
	} else if on {
		clear = 0
		st = 0
	} else {
		set = 0
	}
	return g.writeZoneBit(zone, set, clear)
}

// ProductModel return product model id
func (g *Scanner) ProductModel() (byte, error) {
	return g.readZone([2]byte{0, 0xe0})
}

// HardwareVersion return hardware version
func (g *Scanner) HardwareVersion() (byte, error) {
	return g.readZone([2]byte{0, 0xe1})
}

// SoftwareVersion return hardware version
func (g *Scanner) SoftwareVersion() (byte, error) {
	return g.readZone([2]byte{0, 0xe2})
}

// SoftwareDate return software date
func (g *Scanner) SoftwareDate() (string, error) {
	year, err := g.readZone([2]byte{0, 0xe3})
	if err == nil {
		month, err := g.readZone([2]byte{0, 0xe4})
		if err == nil {
			day, err := g.readZone([2]byte{0, 0xe5})
			if err == nil {
				return fmt.Sprintf("%4d%02d%02d", int(year)+2000, month, day), nil
			}
		}
	}
	return "", err
}

// ReadInterval sets the read interval
func (g *Scanner) ReadInterval(interval byte) error {
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

// SingleReadTime sets the time for a single read
// The longest time before first successful reading.
// After this time, module will be into no read time.
func (g *Scanner) SingleReadTime(readTime byte) error {
	return g.writeZoneByte([2]byte{0, 0x06}, readTime)
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
