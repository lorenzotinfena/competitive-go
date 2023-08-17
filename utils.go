package main

import (
	"bufio"
	"fmt"
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

// Variations:
// Slint(): empty slice
// Slint(n): slice with size n
// Slint(n, c): Slice with size n and capacity c
func Slint(args ...int) []int {
	switch len(args) {
	case 0:
		return []int{}
	case 1:
		return make([]int, args[0])
	case 2:
		return make([]int, args[0], args[1])
	}
	panic("Sli invalid args")
}

// #endregion
