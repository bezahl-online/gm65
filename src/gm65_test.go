package gm65

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGm65 (t *testing.T) {
	var scanner Scanner = Scanner{
		Config: Config{
			SerialPort: "/dev/ttyACM0",
			Baud:       9600,
		},
	}

	err := scanner.Open()
	if err != nil {
		log.Fatal(err)
	}

	var data byte
	var set byte = 0x20
	err =scanner.writeZoneBit([2]byte{0x0,0},set)

	data, err =scanner.readZone([2]byte{0x0,0})
	fmt.Printf("\n%b %d\n",data,data)
	assert.Equal(t,set,data)
}