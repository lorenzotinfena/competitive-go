package main

import (
	"bufio"
	"strconv"
)

// #region IO STUFF

type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func (io *IO) ScanInt() int {
	var x int
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '-' {
			if b >= '0' {
				x = int(b - '0')
			} else {
				b, _ = io.r.ReadByte()
				x = -int(b - '0')
			}
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + int(b-'0')
	}

	return x
}

func (io *IO) ScanInt8() int8 {
	var x int8
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '-' {
			if b >= '0' {
				x = int8(b - '0')
			} else {
				b, _ = io.r.ReadByte()
				x = -int8(b - '0')
			}
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + int8(b-'0')
	}

	return x
}

func (io *IO) ScanInt16() int16 {
	var x int16
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '-' {
			if b >= '0' {
				x = int16(b - '0')
			} else {
				b, _ = io.r.ReadByte()
				x = -int16(b - '0')
			}
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + int16(b-'0')
	}

	return x
}

func (io *IO) ScanInt32() int32 {
	var x int32
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '-' {
			if b >= '0' {
				x = int32(b - '0')
			} else {
				b, _ = io.r.ReadByte()
				x = -int32(b - '0')
			}
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + int32(b-'0')
	}

	return x
}

func (io *IO) ScanUInt() uint {
	var x uint
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '0' {
			x = uint(b - '0')
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + uint(b-'0')
	}

	return x
}

func (io *IO) ScanUInt8() uint8 {
	var x uint8
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '0' {
			x = uint8(b - '0')
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + uint8(b-'0')
	}

	return x
}

func (io *IO) ScanUInt16() uint16 {
	var x uint16
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '0' {
			x = uint16(b - '0')
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + uint16(b-'0')
	}

	return x
}

func (io *IO) ScanUInt32() uint32 {
	var x uint32
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '0' {
			x = uint32(b - '0')
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '0' {
			break
		}
		x = (x * 10) + uint32(b-'0')
	}

	return x
}

func (io *IO) ScanString() []byte {
	var x []byte
	var b byte
	for {
		b, _ = io.r.ReadByte()
		if b >= '!' {
			x = []byte{b}
			break
		}
	}

	for {
		b, _ = io.r.ReadByte()
		if b < '!' {
			break
		}
		x = append(x, b)
	}

	return x
}

func (io *IO) ScanFloat32() float32 {
	x, _ := strconv.ParseFloat(string(io.ScanString()), 32)
	return float32(x)
}

func (io *IO) ScanFloat64() float64 {
	x, _ := strconv.ParseFloat(string(io.ScanString()), 32)
	return x
}

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

func (io *IO) PrintlnSlice(s []any) {
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
	tmp, _ := strconv.Atoi(value)
	return tmp
}

func Itoa(value int) string {
	return strconv.Itoa(value)
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
