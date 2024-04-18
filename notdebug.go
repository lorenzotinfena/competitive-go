//go:build notdebug

package main

import (
	"runtime/debug"
	"syscall"
)

type IO struct {
	rBuffer      []byte
	rBufferIndex int
	rBufferEnd   int
	wBuffer      []byte
	wBufferSize  uint
}

const ioBufferSize = 1 << 15

func (io *IO) nextByte() byte {
	io.rBufferIndex++
	if io.rBufferIndex == io.rBufferEnd {
		io.rBufferEnd, _ = syscall.Read(syscall.Stdin, io.rBuffer)
		io.rBufferIndex = 0
	}
	return io.rBuffer[io.rBufferIndex]
}

func main() {
	io := IO{
		rBufferIndex: -1,
		rBufferEnd:   0,
		rBuffer:      make([]byte, ioBufferSize),
		wBuffer:      make([]byte, ioBufferSize*2),
	}
	if optimizationLevel == OptimizeTimeLimit {
		debug.SetGCPercent(-1)
	} else if optimizationLevel == OptimizeMemoryLimit {
		debug.SetGCPercent(1)
	}
	solve(&io)
	io.Flush()
}

func (io *IO) writeLargeBytes(b []byte) {
	if len(b) > ioBufferSize {
		syscall.Write(syscall.Stdout, b)
		return
	}
	for i := 0; i < len(b); i++ {
		io.wBuffer[io.wBufferSize] = b[i]
		io.wBufferSize++
	}
	if io.wBufferSize > ioBufferSize {
		syscall.Write(syscall.Stdout, io.wBuffer[:io.wBufferSize])
		io.wBufferSize = 0
	}
}

func (io *IO) writeSmallBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		io.wBuffer[io.wBufferSize] = b[i]
		io.wBufferSize++
	}
	if io.wBufferSize > ioBufferSize {
		syscall.Write(syscall.Stdout, io.wBuffer[:io.wBufferSize])
		io.wBufferSize = 0
	}
}

func (io *IO) PrintChar(b byte) {
	io.wBuffer[io.wBufferSize] = b
	io.wBufferSize++
	if io.wBufferSize > ioBufferSize {
		syscall.Write(syscall.Stdout, io.wBuffer[:io.wBufferSize])
		io.wBufferSize = 0
	}
}

func (io *IO) Flush() {
	syscall.Write(syscall.Stdout, io.wBuffer[:io.wBufferSize])
	io.wBufferSize = 0
}

func (io *IO) PrintNewLine() {
	io.wBuffer[io.wBufferSize] = '\n'
	io.wBufferSize++
	if io.wBufferSize > ioBufferSize {
		syscall.Write(syscall.Stdout, io.wBuffer[:io.wBufferSize])
		io.wBufferSize = 0
	}
}

func (io *IO) Print(x any) {
	switch x := x.(type) {
	case string:
		io.writeLargeBytes([]byte(x))
	case []byte:
		io.writeLargeBytes(x)
	case int:
		io.writeSmallBytes(Itoa(x))
	case int8:
		io.writeSmallBytes(Itoa(x))
	case int16:
		io.writeSmallBytes(Itoa(x))
	case int32:
		io.writeSmallBytes(Itoa(x))
	case int64:
		io.writeSmallBytes(Itoa(x))
	case uint:
		io.writeSmallBytes(Uitoa(x))
	case uint8:
		io.writeSmallBytes(Uitoa(x))
	case uint16:
		io.writeSmallBytes(Uitoa(x))
	case uint32:
		io.writeSmallBytes(Uitoa(x))
	case uint64:
		io.writeSmallBytes(Uitoa(x))
	}
}

func (io *IO) Println(x any) {
	switch x := x.(type) {
	case string:
		io.writeLargeBytes([]byte(x))
	case []byte:
		io.writeLargeBytes(x)
	case int:
		io.writeSmallBytes(Itoa(x))
	case int8:
		io.writeSmallBytes(Itoa(x))
	case int16:
		io.writeSmallBytes(Itoa(x))
	case int32:
		io.writeSmallBytes(Itoa(x))
	case int64:
		io.writeSmallBytes(Itoa(x))
	case uint:
		io.writeSmallBytes(Uitoa(x))
	case uint8:
		io.writeSmallBytes(Uitoa(x))
	case uint16:
		io.writeSmallBytes(Uitoa(x))
	case uint32:
		io.writeSmallBytes(Uitoa(x))
	case uint64:
		io.writeSmallBytes(Uitoa(x))
	}
	io.PrintNewLine()
}
