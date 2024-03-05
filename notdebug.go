//go:build !debug

package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

type IO struct {
	rBuffer      []byte
	rBufferIndex uint
	w            *bufio.Writer
}

const rBufferSize = 4096

func (io *IO) nextByte() byte {
	if io.rBufferIndex == rBufferSize-1 {
		syscall.Read(syscall.Stdin, io.rBuffer)
		io.rBufferIndex = 0
	} else {
		io.rBufferIndex++
	}
	return io.rBuffer[io.rBufferIndex]
}

func main() {
	io := IO{
		rBufferIndex: rBufferSize - 1,
		rBuffer:      make([]byte, rBufferSize),
		w:            bufio.NewWriter(os.Stdout),
	}
	defer io.Flush()
	solve(&io)
}

func (io *IO) Print(x ...any) {
	for i := 0; i < len(x); i++ {
		switch x[i].(type) {
		case []byte:
			x[i] = string(x[i].([]byte))
		}
	}

	fmt.Fprint(io.w, x...)
}

func (io *IO) Println(x ...any) {
	for i := 0; i < len(x); i++ {
		switch x[i].(type) {
		case []byte:
			x[i] = string(x[i].([]byte))
		}
	}

	fmt.Fprintln(io.w, x...)
}
