package integers

import "fmt"

var a int8  // [-128, 127]
var b uint8 // [0, 255]

var c int16  // [-32768, 32767]
var d uint16 // [0, 65535]

var e int32  // [-2147483648, 2147483647]
var f4 uint32 // [0, 4294967295]

var g int64  // [-9223372036854775808 to 9223372036854775807]
var h uint64 // [0, 18446744073709551615]

func Demo() {
	x := 300
	fmt.Printf("The type of 'x' is %T\n", x) // int -> the natural integer size for this system architecture (32-bit, 64-bit)
	fmt.Println("The value of 'x' is", x)

	// integer overflow
	var y = uint8(x)
	fmt.Printf("'x' converted into a uint8 becomes %d\n", y)

	fmt.Printf("The binary representation of '300' is %b\n", x)
	fmt.Printf("The binary representation of '44' is %b\n", y)

	// a byte has 8 bits [2^7=128][2^6=64][2^5=32][2^4=16][2^3=8][2^2=4][2^1=2][2^0=1]
	// 300 requires 9 bits = 1 0010 1100
	// uint8 truncates and keeps the first 8 bits (0010 1100) = 44

	// For signed integers, the last bit is the sign bit:
	// 0 = positive, 1 = negative.
	//
	// The remaining 7 bits give 2^7 = 128 possible magnitudes.
	// You can determine the negative bit representation of a number using Two's Complement:
	// Flip every bit then add 1
	// Ex: 00111001 = 57 -> -57 becomes 11000111
}
