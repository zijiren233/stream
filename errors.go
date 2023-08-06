package stream

import (
	"fmt"
)

type FormatUnsupportedTypeError string

func (e FormatUnsupportedTypeError) Error() string {
	return fmt.Sprintf("stream: unsupported type: %s", string(e))
}

type ErrShortLength uint8

func (e ErrShortLength) Error() string {
	return fmt.Sprintf("stream: length must more then %d", e)
}

type ErrLongLength uint

func (e ErrLongLength) Error() string {
	return fmt.Sprintf("stream: length must less then %d", e)
}
