package main

import (
	"bufio"
	"fmt"

	"github.com/lorenzotinfena/goji/collections"
)

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

func (io *IO) Flush() { io.w.Flush() }

// #endregion

// #region USEFUL FUNCTIONS

func (io *IO) ScanSliceInt(length int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = io.ScanInt()
	}
	return res
}

func (io *IO) ScanSlicePairInt(length int) []collections.Pair[int, int] {
	res := make([]collections.Pair[int, int], length)
	for i := 0; i < length; i++ {
		res[i] = collections.MakePair(io.ScanInt(), io.ScanInt())
	}
	return res
}

func (io *IO) PrintlnSliceInt(s []int) {
	tmp := len(s) - 1
	if tmp < 0 {
		io.Println()
	} else {
		for i := 0; i < len(s)-1; i++ {
			io.Print(s[i], " ")
		}
		io.Println(s[tmp])
	}
}

func (io *IO) PrintlnYes() {
	io.Println("Yes")
}

func (io *IO) PrintlnNo() {
	io.Println("No")
}

// #endregion
