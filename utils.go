package main

import (
	"bufio"
	"fmt"
	"os"
)

// When debug is assumed that input.txt file exists
func main() {
	io := IO{
		w: bufio.NewWriter(os.Stdout),
	}
	for _, arg := range os.Args {
		if arg == "∰" { // this should be passed only when debugging
			f, _ := os.Open("input.txt")
			io.r = bufio.NewReader(f)
			goto here
		}
	}
	io.r = bufio.NewReader(os.Stdin)
here:
	defer io.Flush()
	solve(&io)
}

// #region IO STUFF

type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func (io *IO) ScanInt8() (x int8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt16() (x int16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt32() (x int32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt() (x int)     { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanUInt8() (x uint8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt16() (x uint16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt32() (x uint32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt() (x uint)     { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanFloat32() (x float32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanFloat64() (x float64) { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanString() (x string) { fmt.Fscan(io.r, &x); return }

func (io *IO) Print(x ...any)   { fmt.Fprint(io.w, x...) }
func (io *IO) PrintLn(x ...any) { fmt.Fprintln(io.w, x...) }

func (io *IO) Flush() { io.w.Flush() }

// #endregion

// #region USEFUL FUNCTIONS

// Variations:
// Sli(): empty slice
// Sli(n): slice with size n
// Sli(n, c): Slice with size n and capacity c
func Sli[T any](args ...int) []T {
	switch len(args) {
	case 0:
		return []T{}
	case 1:
		return make([]T, args[0])
	case 2:
		return make([]T, args[0], args[1])
	}
	panic("Sli invalid args")
}

// #endregion
