// Package srand is a library for generating random strings, integers, floating point numbers.
package srand

import (
	"math/rand/v2"
	"strconv"
	"time"
)

// nolint
const (
	RNum  = 1 // Only number
	RUpper = 2 // Only capital letters
	RLower = 4 // Only lowercase letters
	RAll   = 7 // Numbers, upper and lower case letters
)

var (
	refSlices = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	kinds     = [][]byte{refSlices[0:10], refSlices[10:36], refSlices[0:36], refSlices[36:62], refSlices[36:], refSlices[10:62], refSlices[0:62]}
)

func init() {
	rand.IntN(int(time.Now().UnixNano()))
}

// String generate random strings of any length of multiple types, default length is 6 if size is empty
// example: String(R_ALL), String(R_ALL, 16), String(R_NUM|R_LOWER, 16)
func String(kind int, size ...int) string {
	return string(Bytes(kind, size...))
}

// Bytes generate random strings of any length of multiple types, default length is 6 if bytesLen is empty
// example: Bytes(R_ALL), Bytes(R_ALL, 16), Bytes(R_NUM|R_LOWER, 16)
func Bytes(kind int, bytesLen ...int) []byte {
	if kind > 7 || kind < 1 {
		kind = RAll
	}

	length := 6 // default length 6
	if len(bytesLen) > 0 {
		length = bytesLen[0]
		if length < 1 {
			length = 6
		}
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = kinds[kind-1][rand.IntN(len(kinds[kind-1]))]
	}

	return result
}

// Int generate random numbers of specified range size,
// compatible with Int(), Int(max), Int(min, max), Int(max, min) 4 ways, min<=random number<=max
func Int(rangeSize ...int) int {
	switch len(rangeSize) {
	case 0:
		return rand.IntN(101) // default 0~100
	case 1:
		return rand.IntN(rangeSize[0] + 1)
	default:
		if rangeSize[0] > rangeSize[1] {
			rangeSize[0], rangeSize[1] = rangeSize[1], rangeSize[0]
		}
		return rand.IntN(rangeSize[1]-rangeSize[0]+1) + rangeSize[0]
	}
}

// Float64 generates a random floating point amount the specified range size,
// Four types of passing references are supported, example: Float64(dpLength), Float64(dpLength, max),
// Float64(dpLength, min, max), Float64(dpLength, max, min), min<=random numbers<=max
func Float64(dpLength int, rangeSize ...int) float64 {
	dp := 0.0
	if dpLength > 0 {
		dpMax := 1
		for i := 0; i < dpLength; i++ {
			dpMax *= 10
		}
		dp = float64(rand.IntN(dpMax)) / float64(dpMax)
	}

	switch len(rangeSize) {
	case 0:
		return float64(rand.IntN(100)) + dp // default 0~100
	case 1:
		return float64(rand.IntN(rangeSize[0])) + dp
	default:
		if rangeSize[0] > rangeSize[1] {
			rangeSize[0], rangeSize[1] = rangeSize[1], rangeSize[0]
		}
		return float64(rand.IntN(rangeSize[1]-rangeSize[0])+rangeSize[0]) + dp
	}
}

// NewID Generate a milliseconds+random number ID.
func NewID() int64 {
	ns := time.Now().UnixMilli() * 1000000
	return ns + rand.Int64N(1000000)
}

// NewStringID Generate a string ID, the hexadecimal form of NewID(), total 16 bytes.
func NewStringID() string {
	return strconv.FormatInt(NewID(), 16)
}

// NewSeriesID Generate a datetime+random string ID,
// datetime is microsecond precision, 20  bytes, random is 6 bytes, total 26 bytes.
// example: 20060102150405000000123456
func NewSeriesID() string {
	// Declare a buffer, only 26 bytes are needed
	var buf [27]byte
	t := time.Now()

	// Format datetime with microsecond precision, and store in the buffer
	copy(buf[:], t.Format("20060102150405.000000"))

	// Generate a 6-digit random number and append it to the buffer.
	random := rand.IntN(1000000)
	buf[21] = '0' + byte(random/100000%10)
	buf[22] = '0' + byte(random/10000%10)
	buf[23] = '0' + byte(random/1000%10)
	buf[24] = '0' + byte(random/100%10)
	buf[25] = '0' + byte(random/10%10)
	buf[26] = '0' + byte(random%10)

	// Return the final string without the dot
	return string(buf[:14]) + string(buf[15:])
}
