package demo

import (
	"io"
	"os"
)

type BytePool struct {
	pool chan []byte
	l    int // 实际长度
	c    int // 容量
}

func NewBytePool(maxSize int, l int, c int) *BytePool {
	return &BytePool{pool: make(chan []byte, maxSize), l: l, c: c}
}

func (bp *BytePool) Get() (b []byte) {
	select {
	case b = <-bp.pool:
	default:
		if bp.c > 0 {
			b = make([]byte, bp.l, bp.c)
		} else {
			b = make([]byte, bp.l)
		}
	}
	return
}

func (bp *BytePool) Put(b []byte) {
	select {
	case bp.pool <- b:
	default:

	}
}

func readFile(b []byte) {
	f, _ := os.Open("data")
	for {
		n, err := io.ReadFull(f, b)

		if n == 0 || err == io.EOF {
			break
		}
	}
}
