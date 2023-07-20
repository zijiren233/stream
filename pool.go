package stream

import (
	"bufio"
	"sync"
)

var bufBytesPool = &sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 8)
		return &buf
	},
}

var bufReaderPool = &sync.Pool{
	New: func() interface{} {
		return bufio.NewReader(nil)
	},
}

var bufWriterPool = &sync.Pool{
	New: func() interface{} {
		return bufio.NewWriter(nil)
	},
}
