package gm65

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0xe0}
	model, err := scanner.ProductModel()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, model, data)
		}
	}
}

func TestHardwareVersion(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0xe1}
	model, err := scanner.HardwareVersion()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, model, data)
		}
	}
}

func TestSoftwareVersion(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0xe2}
	model, err := scanner.SoftwareVersion()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, model, data)
		}
	}
}

func TestSoftwareDate(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0xe3}
	var zone2 [2]byte = [2]byte{0x00, 0xe4}
	var zone3 [2]byte = [2]byte{0x00, 0xe5}
	date, err := scanner.SoftwareDate()
	if assert.NoError(t, err) {
		year, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			month, _ := scanner.readZone(zone2)
			day, _ := scanner.readZone(zone3)
			testDate := fmt.Sprintf("%4d%02d%02d", int(year)+2000, month, day)
			assert.Equal(t, date, testDate)
		}
	}
}

func TestLight(t *testing.T) {
	err := scanner.Light(true, false)
	if assert.NoError(t, err) {
		data, _ := scanner.readZone([2]byte{0x0, 0x0})
		if assert.Equal(t, byte(0x08), data&0x08) {
			scanner.Light(false, false)
			data, _ := scanner.readZone([2]byte{0x0, 0x0})
			if assert.Equal(t, byte(0x00), data&0x0c) {
				scanner.Light(false, true)
				data, _ := scanner.readZone([2]byte{0x0, 0x0})
				fmt.Printf("%08b --> %08b", data, data&0x0c)
				if assert.Equal(t, byte(0x04), data&0x0c) {
					data, _ := scanner.readZone([2]byte{0x0, 0x0})
					assert.Equal(t, byte(0x04), data&0x04)
					scanner.Light(true, false)
				}

			}
		}
	}
}

func TestReadInterval(t *testing.T) {
	var interval byte = 15
	err := scanner.ReadInterval(interval)
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x04})
		if assert.NoError(t, err) {
			assert.Equal(t, interval, data)
		}
	}
}
func TestSingleReadTime(t *testing.T) {
	var readTime byte = 5
	err := scanner.SingleReadTime(readTime)
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x06})
		if assert.NoError(t, err) {
			assert.Equal(t, readTime, data)
		}
	}
}

func TestSensorMode(t *testing.T) {
	err := scanner.SensorMode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			assert.Equal(t, byte(0x03), data&0x03)
		}
	}
}

func TestManualMode(t *testing.T) {
	err := scanner.ManualMode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			assert.Equal(t, byte(0x00), data&0x03)
		}
	}
}

func TestContinuousMode(t *testing.T) {
	err := scanner.ContinuousMode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			assert.Equal(t, byte(0x02), data&0x03)
		}
	}
}
func TestCommandMode(t *testing.T) {
	err := scanner.ContinuousMode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			assert.Equal(t, byte(0x02), data&0x03)
		}
	}
}

func TestOpenLEDOnSuccess(t *testing.T) {
	err := scanner.OpenLEDOnSuccess(true)
	if assert.NoError(t, err) {
		data, _ := scanner.readZone([2]byte{0x0, 0x0})
		if assert.Equal(t, byte(0x80), data&0x80) {
			err := scanner.OpenLEDOnSuccess(false)
			if assert.NoError(t, err) {
				data, _ := scanner.readZone([2]byte{0x0, 0x0})
				assert.Equal(t, byte(0x00), data&0x80)

			}
		}
	}
}

func TestMute(t *testing.T) {
	err := scanner.Mute(false)
	if assert.NoError(t, err) {
		data, err := scanner.readZone([2]byte{0x0, 0x0})
		if assert.NoError(t, err) {
			if assert.Equal(t, byte(0x40), data&0x40) {
				err := scanner.Mute(true)
				if assert.NoError(t, err) {
					data, err := scanner.readZone([2]byte{0x0, 0x0})
					if assert.NoError(t, err) {
						assert.Equal(t, byte(0x0), data&0x40)
					}
				}
			}
		}
	}
}

func TestReadSoundFreq(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x0a}
	var freq int16 = 2000
	var sfn byte = byte(freq / 20)
	err := scanner.ReadSoundFreq(freq)
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, sfn, data)
		}
	}
}

func TestReadSoundDuration(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x0b}
	var duration byte = 255
	err := scanner.ReadSoundDuration(duration)
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, duration, data)
		}
	}
}

func TestDisableAllBarcode(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x2c}
	err := scanner.DisableAllBarcode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(0), data&0x06)
		}
	}
}

func TestEnableEAN13(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x2e}
	err := scanner.EnableEAN13()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(1), data&0x01)
		}
	}
}

func TestEnableEAN8(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x2f}
	err := scanner.EnableEAN8()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(1), data&0x01)
		}
	}
}

func TestEnableCode39(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x36}
	err := scanner.EnableCode39()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(1), data&0x01)
		}
	}
}

func TestEnableQRCode(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x3f}
	err := scanner.EnableQRCode()
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(1), data&0x01)
		}
	}
}

func TestRotateRead360(t *testing.T) {
	var zone [2]byte = [2]byte{0x00, 0x2c}
	err := scanner.RotateRead360(true)
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone)
		if assert.NoError(t, err) {
			assert.Equal(t, byte(1), data&0x01)
		}
	}
}

func TestAutoSleep(t *testing.T) {
	var zone1 [2]byte = [2]byte{0x00, 0x07}
	var zone2 [2]byte = [2]byte{0x00, 0x08}
	var timeMills uint16 = 30
	err := scanner.AutoSleep(false, 0)
	if assert.NoError(t, err) {
		data, err := scanner.readZone(zone1)
		data2, err := scanner.readZone(zone2)
		if assert.NoError(t, err) {
			if assert.Equal(t, byte(0), data&0x80) {
				if assert.Equal(t, byte(0), data2) {
					err := scanner.AutoSleep(true, timeMills)
					if assert.NoError(t, err) {
						data, err := scanner.readZone(zone1)
						data2, err := scanner.readZone(zone2)
						if assert.NoError(t, err) {
							if assert.Equal(t, byte(0x80), data&0x80) {
								assert.Equal(t, byte(30), data2)
							}
						}
					}
				}
			}
		}
	}
}
