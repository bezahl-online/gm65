package param

import "flag"

var Disconnected *bool

func init() {
	Disconnected = flag.Bool("disconnected", false, "do not report missing physical device (testmode)")
	flag.Parse()
}
