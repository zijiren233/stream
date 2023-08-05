package stream

import (
	"errors"
	"fmt"
)

var ErrUnsupportedType = errors.New("unsupported type")

type FormatUnsupportedTypeError string

func (e FormatUnsupportedTypeError) Error() string {
	return fmt.Sprintf("stream: unsupported type: %s", string(e))
}
