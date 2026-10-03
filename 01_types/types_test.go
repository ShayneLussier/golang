package types

import "fmt"

func Example_predeclaredDataTypes() {
	var a bool = true

	var b int = -45
	var c uint = 16 // non-negative
	// special integer type 'byte' is equivalent to uint8
	var d float64 = 13.25

	var e rune = 'T' // a rune is a single character, must use single quotes.
	// Equivalent to int32, returns the character's Unicode code point by default in decimal format.
	var f string = "Hello, World!" // strings are immutable (cannot be altered)

	fmt.Println(a, b, c, d, string(e), f)
	// Output: true -45 16 13.25 T Hello, World!
}

func Example_zeroValue() {
	// When variables are declared but not initialized they carry the data types' zero value
	// initializing a variable with it's zero value makes it clear that it is intended
	var (
		a bool    // false
		b int     // 0
		c uint    // 0
		d float64 // 0
		e string  // ""
	)

	fmt.Println(a, b, c, d, e)
	// Output: false 0 0 0
}

func Example_operators1() {
	var a, b = 10, 5
	// the result of integer operations also return an 'int' type
	// to get a floating-point result, we must cast the int to a float64
	// dividing a int by 0 will panic

	var add = a + b
	var substract = a - b
	var multiply = a * b
	var divide = a / b
	var remainder = a % b

	fmt.Println(add, substract, multiply, divide, remainder)
	// Output: 15 5 50 2 0
}

func Example_operators2() {
	var x = 10
	// you can combine an operator with '=' the modify a variable (+=, -=, *=, /=, %=)

	x *= 2

	fmt.Println(x)
	// Output: 20
}

func Example_operators3() {
	var x, y int = 1, 2
	// relational operators
	var equal = x == y
	var notEqual = x != y

	fmt.Println(equal, notEqual)
	// Output: false true

}

func Example_operators4() {
	var x, y int = 1, 2
	// comparison operators (>, >=, <, <=)
	var smaller = x <= y

	fmt.Println(smaller)
	// Output: true
}

func Example_typeConversion() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)

	fmt.Println(sum1, sum2)
	// Output: 40.2 40
}

func Example_declaration() {
	// most verbose way
	var a int = 10

	// if the type on the right of '=' is the expected type you can ommit it
	var b = 20

	// multiple variables at the same time. can be different types
	var c, d = 30, "Pizza"

	// in a declaration list
	var (
		w    int
		x    int = 40
		y, z     = 50, "Ice Cream"
	)

	// using shorthand notation
	// only available inside functions
	// most common declaration style
	age := 35

	fmt.Println(a, b, c, d, w, x, y, z, age)
	// Output: 10 20 30 Pizza 0 40 50 Ice Cream 35
}

func Example_constants() {
	// constants are set when the code is compiled and can never change
	// can be type or untyped
	const pi = 3.14159
	// pi = 3 // error: can't change a constant

	const secondsPerDay = 60 * 60 * 24 // ok: math the compiler can do
	// const now = time.Now() // error: only known when the program runs

	fmt.Println(pi, secondsPerDay)
	// Output: 3.14159 86400
}
