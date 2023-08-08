//go:build !debug

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := IO{
		w: bufio.NewWriter(os.Stdout),
		r: bufio.NewReader(os.Stdin),
	}
	defer io.Flush()
	solve(&io)
}
func (io *IO) Print(x ...any)   { fmt.Fprint(io.w, x...) }
func (io *IO) PrintLn(x ...any) { fmt.Fprintln(io.w, x...) }
