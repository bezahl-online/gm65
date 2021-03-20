package device

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var connecting bool = false

// Connect manages the connection of the scanner
func Connect(s *Scanner) {
	if connecting {
		return
	}
	connecting = true
	var serialPortName string = "/dev/ttyACM0"
	if len(os.Getenv("GM65_PORT_NAME")) > 3 {
		serialPortName = os.Getenv("GM65_PORT_NAME")
	}
	var mutex *sync.RWMutex = (*s).lock
	if mutex == nil {
		mutex = &sync.RWMutex{}
	}
	*s = Scanner{
		lock:   mutex,
		Config: Config{SerialPort: serialPortName, Baud: 9600},
		port:   nil,
		command: &command{
			Head:     [2]byte{},
			Function: 0,
			Length:   0,
			Address:  [2]byte{},
			Data:     0,
			CRC:      [2]byte{},
		},
		connected: false,
	}
	var pause delay = delay{
		dur: 1 * time.Second,
	}

	if s.connected {
		(*s.port).Close()
	}

	// connect to gm65 scanner on /dev/ttyACM0
	var err error = fmt.Errorf("gm65 not connected")
	fmt.Printf("connecting to scanner device on '%s'\n", serialPortName)
	for err != nil {
		err := s.Open()
		if err != nil {
			fmt.Printf("\n*** Error while connection to gm65 scanner:"+
				" %s\nRetrying after %d seconds\n", err.Error(), pause.getSeconds())
			if pause.getSeconds() < 300 {
				pause.double()
			}
			pause.wait()
		} else {
			fmt.Println("device successfully connected")
			s.connected = true
			s.Configure()
			go s.Listen()
			break
		}
	}
	connecting = false
}

type delay struct {
	dur time.Duration
}

func (w *delay) getSeconds() int {
	return int((*w).dur.Seconds())
}

func (w *delay) wait() {
	time.Sleep(w.dur)
}

func (w *delay) double() {
	(*w).dur *= 2
}
