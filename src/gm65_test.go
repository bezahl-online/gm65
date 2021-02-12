package gm65

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var scanner Scanner

func init() {
	scanner = Scanner{
		Config: Config{
			SerialPort: "/dev/ttyACM0",
			Baud:       9600,
		},
	}

	err := scanner.Open()
	if err != nil {
		log.Fatal(err)
	}
}
func TestWriteZoneBit(t *testing.T) {

	var data byte
	var set byte = 0x30
	var clear byte = 0x00
	err := scanner.writeZoneBit([2]byte{0x0, 0}, set, clear)
	if assert.NoError(t, err) {
		data, err = scanner.readZone([2]byte{0x0, 0})
		fmt.Printf("\n%b %d\n", data, data)
		assert.Equal(t, set, (data&set)&^clear)
	}
}

func TestSetLight(t *testing.T) {
	err := scanner.SetLight(false, false)
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			if assert.Equal(t, byte(0x30), data) {
				err := scanner.SetLight(true, false)
				if assert.NoError(t, err) {
					data, err := scanner.readZone([2]byte{0x0, 0x0})
					if assert.NoError(t, err) {
						if assert.Equal(t, byte(0x08), data&0x08) {
							err := scanner.SetLight(false, true)
							if assert.NoError(t, err) {
								data, err := scanner.readZone([2]byte{0x0, 0x0})
								if assert.NoError(t, err) {
									assert.Equal(t, byte(0x04), data&0x04)
									scanner.SetLight(true,false)
								}
							}
						}
					}
				}
			}
		}
	}
}

func TestSetReadInterval(t *testing.T) {
	var interval byte =15
	err:= scanner.SetReadInterval(interval)
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x04})
		if assert.NoError(t,err) {
			assert.Equal(t,interval,data)
		}
	}
}
func TestSetSensorMode(t *testing.T) {
	err:= scanner.SetSensorMode()
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			assert.Equal(t,byte(0x03),data&0x03)
		}
	}
}

func TestSetManualMode(t *testing.T) {
	err:= scanner.SetManualMode()
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			assert.Equal(t,byte(0x00),data&0x03)
		}
	}
}

func TestSetContinuousMode(t *testing.T) {
	err:= scanner.SetContinuousMode()
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			assert.Equal(t,byte(0x02),data&0x03)
		}
	}
}
func TestSetCommandMode(t *testing.T) {
	err:= scanner.SetContinuousMode()
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			assert.Equal(t,byte(0x02),data&0x03)
		}
	}
}
func TestSetOpenLEDOnSuccess(t *testing.T) {
	err:= scanner.SetOpenLEDOnSuccess()
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			assert.Equal(t,byte(0x03),data&0x03)
		}
	}
}

func TestSetMute(t *testing.T) {
	err:= scanner.SetMute(false)
	if assert.NoError(t,err) {
		data,err:= scanner.readZone([2]byte{0x0,0x0})
		if assert.NoError(t,err) {
			if assert.Equal(t,byte(0x40),data&0x40) {
				err:= scanner.SetMute(true)
				if assert.NoError(t,err) {
					data,err:= scanner.readZone([2]byte{0x0,0x0})
					if assert.NoError(t,err) {
						assert.Equal(t,byte(0x0),data&0x40)
					}
				}
			}
		}
	}
}
