package lib

import "log"

func InitLog() {
	log.SetFlags(0)
	log.SetPrefix(argv0 + ": ")
}
