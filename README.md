[![Go Report Card](https://goreportcard.com/badge/github.com/zijiren233/stream)](https://goreportcard.com/report/github.com/zijiren233/stream)
[![GoDoc](https://godoc.org/github.com/zijiren233/stream?status.svg)](https://godoc.org/github.com/zijiren233/stream)
    
# Stream

- stream is a simple and fast binary stream reader and writer for golang.

# Usage

- This is ping example

```go
package main

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/zijiren233/stream"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
	Data        []byte
}

type IPHeader struct {
	Version   uint8   // 4 bits
	HeaderLen uint8   // 4 bits
	TOS       uint8   // 8 bits
	TotalLen  uint16  // 16 bits
	ID        uint16  // 16 bits
	Flag      uint8   // 3 bits
	Offset    uint16  // 13 bits
	TTL       uint8   // 8 bits
	Protocol  uint8   // 8 bits
	Checksum  uint16  // 16 bits
	Src       [4]byte // 32 bits
	Dst       [4]byte // 32 bits
}

var (
	size = 0
)

func main() {
	c, err := net.DialTimeout("ip:icmp", "www.baidu.com", time.Second*3)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	icmp := ICMP{
		Type:        8,
		Code:        0,
		Checksum:    0,
		Identifier:  1,
		SequenceNum: 1,
		Data:        make([]byte, size),
	}

	buf := bytes.NewBuffer(make([]byte, 0, 8+size))
	w := stream.NewWriter(buf, stream.BigEndian)
	if err := w.Write(&icmp).Error(); err != nil {
		panic(err)
	}
	stream.BigEndian.WriteU16(buf.Bytes()[2:4], checkSun(buf.Bytes()))
	_, err = c.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	resp := make([]byte, 1024)
	n, err := c.Read(resp)
	if err != nil {
		panic(err)
	}
	ipHeader := IPHeader{}
	r := stream.NewReader(bytes.NewReader(resp[:n]), stream.BigEndian)
	if err := r.
		Bits(4, &ipHeader.Version).
		Bits(4, &ipHeader.HeaderLen).
		U8(&ipHeader.TOS).
		U16(&ipHeader.TotalLen).
		U16(&ipHeader.ID).
		Bits(3, &ipHeader.Flag).
		Bits(13, &ipHeader.Offset).
		U8(&ipHeader.TTL).
		U8(&ipHeader.Protocol).
		U16(&ipHeader.Checksum).
		Bytes(ipHeader.Src[:]).
		Bytes(ipHeader.Dst[:]).
		Error(); err != nil {
		panic(err)
	}

	fmt.Printf("ipHeader: %+v\n", ipHeader)
}

func checkSun(data []byte) uint16 {
	var (
		sum uint32
	)
	if len(data)%2 != 0 {
		data = data[:len(data)-1]
		sum += uint32(data[len(data)-1])
	}
	r := stream.NewReader(bytes.NewReader(data), stream.BigEndian)
	var v uint16
	for {
		if err := r.U16(&v).Error(); err != nil {
			break
		}
		sum += uint32(v)
	}

	for sum>>16 > 0 {
		sum = (sum >> 16) + (sum & 0xff00)
	}

	return uint16(^sum)
}
```