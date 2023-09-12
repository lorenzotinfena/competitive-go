//go:build debug

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	io := IO{
		w: bufio.NewWriter(os.Stdout),
		r: bufio.NewReader(f),
	}
	defer io.Flush()
	solve(&io)
}
func (io *IO) Print(x ...any)   { fmt.Fprint(io.w, x...); io.Flush() }
func (io *IO) Println(x ...any) { fmt.Fprintln(io.w, x...); io.Flush() }
