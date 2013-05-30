package lib

import (
	"os"
	"strconv"
)

func ParsePerm(s string) (os.FileMode, error) {
	p, err := strconv.ParseInt(s, 8, 32)
	return os.FileMode(p), err
}
