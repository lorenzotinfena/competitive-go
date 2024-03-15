package main

import (
	"strconv"
)

type OPTIMIZATION_LEVEL int

const (
	Balanced OPTIMIZATION_LEVEL = iota
	OptimizeTimeLimit
	OptimizeMemoryLimit
)

// #region IO STUFF

func (io *IO) ScanLine() []byte {
	x := []byte{}
	var b byte
	for {
		b = io.nextByte()
		if b <= '\n' {
			break
		}
		x = append(x, b)
	}

	return x
}

func (io *IO) ScanInt() int {
	var x int
	var b byte
	for {
		b = io.nextByte()
		if b >= '-' {
			if b >= '0' {
				x = int(b - '0')
			} else {
				b = io.nextByte()
				x = -int(b - '0')
			}
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '-' {
			if b >= '0' {
				x = int8(b - '0')
			} else {
				b = io.nextByte()
				x = -int8(b - '0')
			}
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '-' {
			if b >= '0' {
				x = int16(b - '0')
			} else {
				b = io.nextByte()
				x = -int16(b - '0')
			}
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '-' {
			if b >= '0' {
				x = int32(b - '0')
			} else {
				b = io.nextByte()
				x = -int32(b - '0')
			}
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '0' {
			x = uint(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '0' {
			x = uint8(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '0' {
			x = uint16(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '0' {
			x = uint32(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
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
		b = io.nextByte()
		if b >= '!' {
			x = []byte{b}
			break
		}
	}

	for {
		b = io.nextByte()
		if b < '!' {
			break
		}
		x = append(x, b)
	}

	return x
}

func (io *IO) ScanFloat32() float32 {
	var x float32
	var b byte
	negative := false
	for {
		b = io.nextByte()
		if b >= '-' {
			if b < '0' {
				negative = true
				b = io.nextByte()
			}
			x = float32(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
		if b <= '.' {
			break
		}
		x = (x * 10) + float32(b-'0')
	}
	if b == '.' {
		tmp := float32(1)
		for {
			b = io.nextByte()
			if b < '0' {
				break
			}
			tmp *= 10
			x = (x * 10) + float32(b-'0')
		}
		x /= tmp
	}
	if negative {
		x = -x
	}
	return x
}

func (io *IO) ScanFloat64() float64 {
	var x float64
	var b byte
	negative := false
	for {
		b = io.nextByte()
		if b >= '-' {
			if b < '0' {
				negative = true
				b = io.nextByte()
			}
			x = float64(b - '0')
			break
		}
	}

	for {
		b = io.nextByte()
		if b <= '.' {
			break
		}
		x = (x * 10) + float64(b-'0')
	}
	if b == '.' {
		tmp := float64(1)
		for {
			b = io.nextByte()
			if b < '0' {
				break
			}
			tmp *= 10
			x = (x * 10) + float64(b-'0')
		}
		x /= tmp
	}
	if negative {
		x = -x
	}
	return x
}

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
		io.PrintNewLine()
	} else {
		for i := 0; i < len(s)-1; i++ {
			io.Print(s[i])
			io.PrintChar(' ')
		}
		io.Println(s[tmp])
	}
}

func (io *IO) PrintlnSliceUInt(s []uint) {
	tmp := len(s) - 1
	if len(s) == 0 {
		io.PrintNewLine()
	} else {
		for i := 0; i < len(s)-1; i++ {
			io.Print(s[i])
			io.PrintChar(' ')
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

func Itoa[T int | int8 | int16 | int32 | int64](value T) []byte {
	if value == 0 {
		return []byte{'0'}
	}
	digits := 0
	value1 := value
	for {
		digits++
		value1 /= 10
		if value1 == 0 {
			break
		}
	}
	var res []byte
	if value > 0 {
		res = make([]byte, digits)
		digits--
	} else {
		value = -value
		res = make([]byte, digits+1)
		res[0] = '-'
	}
	for {
		res[digits] = byte('0' + value%10)
		value /= 10
		if value == 0 {
			break
		}
		digits--
	}
	return res
}

func Uitoa[T uint | uint8 | uint16 | uint32 | uint64](value T) []byte {
	if value == 0 {
		return []byte{'0'}
	}
	digits := -1
	value1 := value
	for {
		digits++
		value1 /= 10
		if value1 == 0 {
			break
		}
	}
	res := make([]byte, digits+1)
	for {
		res[digits] = byte('0' + value%10)
		value /= 10
		if value == 0 {
			break
		}
		digits--
	}
	return res
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
