package stream

import (
	"fmt"
)

type FormatUnsupportedTypeError string

func (e FormatUnsupportedTypeError) Error() string {
	return fmt.Sprintf("stream: unsupported type: %s", string(e))
}

type FormatLengthError uint8

func (e FormatLengthError) Error() string {
	return fmt.Sprintf("stream: length must be %d", e)
}
