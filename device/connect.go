package device

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/term"
)

var connecting bool = false

// Connect manages the connection of the scanner
func (s *Scanner) Connect() {
	s.connected = false
	if s.port != nil {
		s.port.Close()
	}
	s.Config.SerialPort = "/dev/ttyACM0"
	if len(os.Getenv("GM65_PORT_NAME")) > 3 {
		s.Config.SerialPort = os.Getenv("GM65_PORT_NAME")
	}
	var pause delay = delay{
		dur: 1 * time.Second,
	}
	var err error = fmt.Errorf("gm65 not connected")
	fmt.Printf("connecting to scanner device on '%s'\n", s.Config.SerialPort)
	for {
		s.port, err = term.Open(s.Config.SerialPort)
		if err != nil {
			fmt.Printf("\n*** Error while connection to gm65 scanner:"+
				" %s\nRetrying after %d seconds\n", err.Error(), pause.getSeconds())
			if pause.getSeconds() < 30 {
				pause.double()
			}
			pause.wait()
		} else {
			// s.port.SetReadTimeout(time.Second)
			fmt.Println("device successfully connected")
			s.connected = true
			s.port.SetRaw()
			s.Configure()
			go s.Listen()
			break
		}
	}
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
