//go:build debug

package main

import (
	"bufio"
	"fmt"
	"os"
)

type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func (io *IO) nextByte() byte {
	tmp, _ := io.r.ReadByte()
	return tmp
}

func main() {
	f, _ := os.Open("input.txt")
	io := IO{
		w: bufio.NewWriter(os.Stdout),
		r: bufio.NewReader(f),
	}
	solve(&io)
}

func (io *IO) PrintNewLine() {
	fmt.Fprintln(io.w)
	io.Flush()
}

func (io *IO) PrintChar(b byte) {
	fmt.Fprint(io.w, string([]byte{b}))
	io.Flush()
}

func (io *IO) Print(x ...any) {
	for i := 0; i < len(x); i++ {
		switch x[i].(type) {
		case []byte:
			x[i] = string(x[i].([]byte))
		}
	}

	fmt.Fprint(io.w, x...)
	io.Flush()
}

func (io *IO) Println(x ...any) {
	for i := 0; i < len(x); i++ {
		switch x[i].(type) {
		case []byte:
			x[i] = string(x[i].([]byte))
		}
	}

	fmt.Fprintln(io.w, x...)
	io.Flush()
}

func (io *IO) Flush() { io.w.Flush() }
