package gm65

import (
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
		//fmt.Printf("\n%b %d\n", data, data)
		assert.Equal(t, set, (data&set)&^clear)
	}
}

// func TestReadCode(t *testing.T) {
// 	var code []byte
// 	var err error
// 	scanner.DisableAllBarcode()
// 	scanner.EnableQRCode()
// 	scanner.EnableEAN13()
// 	code,err = scanner.read()
// 	if err==nil {
// 		fmt.Println("code: "+string(code))
// 	}
// }
