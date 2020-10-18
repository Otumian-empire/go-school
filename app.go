package main

func main() {
	// fmt.Println("Helloworld go - tutorialspoint")

	// fmt.Println("tokens-vars-data")
	// Single line comment

	/*
		multiline
		comment
	*/

	// Identifiers
	/*
		These are names used to identify variables, functions, ...
		identifier = letter { letter | unicode_digit }.
		@, $, and % is not allowed in identifiers


		Keywords:
		break 	default 	func 	interface 	select
		case 	defer 	Go 	map 	Struct
		chan 	else 	Goto 	package 	Switch
		const 	fallthrough 	if 	range 	Type
		continue 	for 	import 	return 	Var
	*/

	// Datatypes
	/*
		data types refer to an extensive system used for declaring variables or
			functions of different types

		1) Boolean types
			They are boolean types and consists of the two predefined constants: (a) true (b) false

		2) Numeric types
			They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.

		3) String types
			A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.

		4) Derived types
			They include (a) Pointer types, (b) Array types, (c) Structure types, (d) Union types and (e) Function types f) Slice types g) Interface types h) Map types i) Channel Types

		Array types and structure types are collectively referred to as aggregate types.
	*/

	// Variable Definition in Go
	// var variable_list optional_data_type;

	/*
		var age int
		var name string

		age = 354
		name = "96"

		fmt.Println(age, name)

		var xCord, yCord float64
		xCord = 12.33
		yCord = 34.55

		fmt.Println("The coordinatinates are ", xCord, " and ", yCord)
	*/

	// A dynamic type variable declaration requires the compiler to interpret the type
	// of the variable based on the value passed to it

	/*
		goVersion := 23.34
		fmt.Println(goVersion)

		var a, b, c = 1, 2, 3
		fmt.Println(a, b, c)

		// constant
		const PI float32 = 3.1432
		var radius float32 = 5.56
		var areaCircle float32 = PI * radius * radius

		fmt.Println("A circle of radius,", radius, "has an area of", areaCircle)

		// operators
		var num1, num2, res int = 2, 3, 0
		res = num1 + num2 // +, -, ..
		fmt.Println(res)

		fmt.Println(res == (num1*num1 - 4))
		fmt.Println(res > (num1*num1-4) && true)

		res += 1 - num2
		fmt.Println(res)
	*/

	// conditions
	/*
		var PASSMARK float32 = 50.0

		var myScore float32 = 56.5

		if myScore >= PASSMARK {
			fmt.Println("Suhnil passed")
		} else {
			fmt.Println("Suhnil did not pass")
		}
	*/

	/*
		if myScore >= 69.0 {
			fmt.Println("Grade A")
		} else if myScore >= 59.0 {
			fmt.Println("Grade B")
		} else if myScore >= 50.0 {
			fmt.Println("Grade C")
		} else {
			fmt.Println("Grade F")
		}
	*/

	/*
		switch {
		case myScore >= 69.0:
			fmt.Println("Grade A")
		case myScore >= 59.0:
			fmt.Println("Grade B")
		case myScore >= 50.0:
			fmt.Println("Grade C")
		default:
			fmt.Println("Grade F")

		}
	*/

}
