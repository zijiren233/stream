package stream

import "errors"

var (
	ErrAlreadyClosed = errors.New("stream: already closed")
)
