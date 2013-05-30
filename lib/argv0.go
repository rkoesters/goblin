package lib

import "os"

var argv0 string

func init() {
	argv0 = os.Args[0]
}
