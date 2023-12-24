package main

import (
	"bufio"
	"fmt"
	"strconv"
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

func (io *IO) ScanString() (x []byte) { fmt.Fscan(io.r, &x); return }

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

func (io *IO) ScanSlicePairInt(length int) []II {
	res := make([]II, length)
	for i := 0; i < length; i++ {
		res[i] = II{io.ScanInt(), io.ScanInt()}
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

func Atoi(value string) int {
	tmp, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return tmp
}

func Itoa(value int) string {
	tmp := strconv.Itoa(value)
	return tmp
}

type II struct {
	First  int
	Second int
}

type III struct {
	First  int
	Second int
	Third  int
}

type IIII struct {
	First  int
	Second int
	Third  int
	Fourth int
}

// #endregion
