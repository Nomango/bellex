package packet

import "errors"

var (
	errPrefix = "error:"
	errSuffix = ";"
)

func NewError(msg string) error {
	return errors.New(errPrefix + msg + errSuffix)
}
