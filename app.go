package main

import "fmt"

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

	/*
		for [condition |  ( init; condition; increment ) | Range] {
		   statement(s);
		}
	*/

	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i, " - ", i*i)
		}

		var i, sumi int = 0, 0
		for i < 10 {
			i++
			sumi++
		}
		fmt.Println(sumi)

		numbers := [5]int{1, 2, 3, 4, 5}
		var sumElements int = 0

		for index, element := range numbers {
			fmt.Println("There is", element, "at position", index, "of the array")
			sumElements += element
		}
		fmt.Println(sumElements)
	*/

	// using a function
	/*
		var maxVal = max(4, 6)
		fmt.Println(maxVal)
	*/

	// strings
	/*
		var myName = "Daniel Cormier"

		// len of string
		var lenMyName = len(myName)

		fmt.Println("My name is", myName)
		fmt.Println("my name is about", lenMyName, "characters long")
	*/

	// arrays

	// declaring arrays
	/*
		var variable_name [SIZE] variable_type
	*/
	/*
		var numbers [5]int

		numbers[0] = 2

		fmt.Println(numbers)

		for i := 1; i < 5; i++ {
			numbers[i] = i * i
		}
		fmt.Println(numbers)

		// Initializing Arrays
		var someNumber = [5]int{0, 1, 10, 11, 12}
		fmt.Println(someNumber)

		for i := 0; i < 5; i++ {
			fmt.Println(numbers[i] + someNumber[0])
		}

		// n-Dimensional array
		var twoDArray = [2][2]int{{1, 2}, {3, 4}}

		fmt.Println(twoDArray)

		for row := 0; row < 2; row++ {
			for col := 0; col < 2; col++ {
				fmt.Println("row", row, ", col", col, ":", twoDArray[row][col])
			}
		}

		printArrayI([]int{2,3,4,5,6})
	*/

	// structure
	/*
		type struct_variable_type struct {
			member definition;
			member definition;
			...
			member definition;
		 }
	*/

	/*
		var micsProf Profile

		micsProf.name = "John Doe"
		micsProf.age = 32
		micsProf.height = 6.4

		printProfile(micsProf)

		// map = key => value pair
		// var_name := map[key_dt]value_dt {key:val, ..., keyn:valn}
		sid := map[string]string{"name": "Predentition", "age": "23 guest"}

		fmt.Println("Sids name is", sid["name"])
		fmt.Println("Sid is", sid["age"], "years old")
		// var var_name map[key_type]value_type
		var myMap map[string]string
		myMap = make(map[string]string)

		myMap["name"] = "John Doe"
		myMap["age"] = "34 years"

		fmt.Println("Sids name is", myMap["name"])
		fmt.Println("Sid is", myMap["age"])

		key, exists := sid["age"]

		if exists {
			fmt.Println(key, "exists")
		} else {
			myMap["age"] = "40"
		}

		for key := range sid {
			fmt.Println(key, "=>", sid[key])
		}

		// delete from the map
		delete(sid, "age")

		fmt.Println(sid)

		fmt.Println("The factorial of 5 is", factorial(5))
	*/

	// type casting
	// type_name(expression)

	/*
		var totalScore, numberOfTest int = 10, 97

		var average float32 = float32(totalScore / numberOfTest)
		fmt.Println("Average:", average)

		var average1 float32 = float32(totalScore) / float32(numberOfTest)
		fmt.Println("Average1:", average1)
	*/

	// interface

	pitbull := English{name: "John Doe"}
	squash := Arab{name: "Salman Doe"}

	pitbull.sayHello()
	squash.sayHello()

}

/*
	type interface_name interface {
		method_name() return_type
		...
		method_nameX() return_type
	}
*/

// Human insterface
type Human interface {
	sayHello()
}

// Call the object that implements the Human interface
func objectSayHello(human Human) {
	human.sayHello()
}

// English object to use Human inteterface
type English struct {
	name string
}

// Arab object to use Human inteterface
type Arab struct {
	name string
}

// Implementing the interface for the English man
func (patrick English) sayHello() {
	fmt.Println("Hello," + patrick.name + ".")
}

func (hassan Arab) sayHello() {
	fmt.Println("Halu," + hassan.name + ".")
}

// Factorial : this is a dimple factorial program using recursion
// func factorial(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}

// 	return n * factorial(n-1)

// }

// // Profile this is a profile struct
// type Profile struct {
// 	name   string
// 	age    int
// 	height float32
// }

// func printProfile(prof Profile) {
// 	fmt.Println("My name is", prof.name, ", I am", prof.age, "years old.")
// 	fmt.Println("I am", prof.height)
// }

/*
	func function_name( [parameter list] ) [return_types]
	{
	   body of the function
	}
*/

// func max(num1, num2 int) int {
// local variable declaration
// 	var result int

// 	if num1 >= num2 {
// 		result = num1
// 	} else {
// 		result = num2
// 	}
// 	return result
// }

// func swap(a, b int) (int, int) {
// 	return b, a
// }

// the above function is the sa,me as the below code
/*
	func swap(a int, b int) (int, int) {
		return b,a
	}
*/

// Prints the content of a given integer array
// func printArrayI(intArr []int) {
// 	for index, element := range intArr {
// 		fmt.Println(index, "has the value", element)
// 	}
// }
