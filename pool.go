package stream

import "sync"

var bufPool = &sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 8)
		return &buf
	},
}
