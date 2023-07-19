package pool

import "sync"

var BufPool = &sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 8)
		return &buf
	},
}
