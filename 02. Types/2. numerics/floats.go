package floats

import (
	"fmt"
	"math"
)

var f1 float32 // approximately 7 signigicant digits of precision
var f2 float64 // approximately 15-17 signigicant digits of precision

func Demo() {
	var x float32 = 13.25
	fmt.Printf("The binary representation of 'x' is %032b\n", math.Float32bits(x)) // 01000001010101000000000000000000

	// ┌───┬──────────┬───────────────────────┐
	// │ S │ Exponent │       Fraction        │
	// │ 1 │    8     │          23           │
	// └───┴──────────┴───────────────────────┘

	// 13.25 -> 13 = 1101
	// .25 -> 01
	// 13.25 = 1101.01
	//
	// normalize: shift the point left until one 1 remains to its left.  // 1101.01 = 1.10101 x 2^3   (moved 3 places -> exponent 3)
	//  // sign     = 0                 positive
	// exponent = 3 + 127 = 130     bias 127 so negatives fit: 10000010
	// fraction = 10101             the leading 1 is implied, so it is dropped. pad right to 23 bits
	//
	// | 0 | 10000010 | 10101000000000000000000 |

}
