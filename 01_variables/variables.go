package variables

import "fmt"

var packageLevel string

func variables() {
	// variables can be declared at the package or the function level

	var functionLevel string

	fmt.Println(packageLevel, functionLevel)
}

func initializers() {
	// variables are 'initialized' when they are given a value
	// variables can be declared and initialized at the same time
	// if an initializer is present, the type will be inferred
	// use ':=' notation as shorthand assignment
	// shorthand assignment is only available at the function level

	var i int = 1
	var j = 2
	k := 3
	fmt.Println(i, j, k)
}

func multipleValues() {
	var (
		firstName string = "John"
		lastName         = "Doe"
		age              = 30
	)

	fmt.Println(firstName, lastName, age)
}

func zeroValues() {
	// variables declared without initialization are given a 'zero value'
	// the zero value is:
	// '0' for numeric types
	// 'false' for boolean types
	// '""' (the empty string) for strings

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func constants() {
	// constants are variables that remain unchanged
	// they are usually declared at the package level
	const fileExtension string = ".csv"
	const databasePort int = 8000

	const pi = 3.14159265358979
	// pi = 1 + 2 // will not compile

	// a constant may be used in an expression but will remain unchanged
	fmt.Println(pi+1, pi)
}

func _iota() {
	// iota is used to generate a sequence of constants, it starts at 0

	const (
		Pending    = iota // 0
		Processing        // 1
		Shipped           // 2
		Delivered         // 3
		Cancelled         // 4
	)

}
