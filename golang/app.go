package main

import (
	"fmt"
	"reflect"

	"golang.dojo/utility"
)

func main() {
	fmt.Print("\033[H\033[2J")
	utility.PrintHeader("GO CLI APP")

	demoStrings()
	demoNumbers()
	demoBooleans()
	demoErrors()
	demoVariableTypesAndOnlyExplicityCoversions()
	demoArithmeticOperators()
	demoComparisonOperators()
}

func demoStrings() {
	utility.PrintTitle("STRINGS")
	utility.PrintNotes(`
	Simple Data Types (only 4)
		- Strings
		- Numbers
		- Booleans
		- Errors	
	`)

	var interpretedString = "This is an interpreted \nstring with a newline character"
	fmt.Println(interpretedString)
	utility.AddSeparator()
	var rawString = `This is a raw \n string with a newline character`
	fmt.Println(rawString)
	utility.AddSeparator()
	var multiLineRawString = `This is a
	Multi Line
		Raw String
	`
	fmt.Println(multiLineRawString)
	utility.AddSeparator()
}

func demoNumbers() {
	utility.PrintTitle("NUMBERS")
	utility.PrintNotes(`
	Note: There are more types than this
		int
		uint - unsigned integers
		float32 | float64 - floating point number
		complext64 | complex128 - complex numbers
	`)
}

func demoBooleans() {
	utility.PrintTitle("BOOLEANS")
	utility.PrintNotes(`
	 	Booleans
		true | false
	`)
}

func demoErrors() {
	utility.PrintTitle("ERRORS")
	utility.PrintNotes(`
		Error Types are Unique to Golang as they are premitive datatype	
	`)
}

func demoVariableTypesAndOnlyExplicityCoversions() {
	utility.PrintTitle("VARIABLES, TYPES, CONVE")
	var myNameDeclare string                 //preferred
	var myNameDeclareAndInit string = "Mike" //preferred

	var myNameInffered1 = "Mike"
	myNameInfferredPref := "Mikey" //preferred
	fmt.Println(myNameDeclare, myNameDeclareAndInit, myNameInffered1, myNameInfferredPref)

	//Type Conversions: Being clear over clever
	var i int = 35
	// var f float32

	// f = i // cannot implicitly conversion, go doesn't support that
	f := float32(i) //can use explicity type conversion because it is clear
	fmt.Println(reflect.TypeOf(f))

	d := 3.1414
	c := 234
	fmt.Printf("%f ", d)
	fmt.Printf("%d ", c)

	fmt.Println()
}

func demoArithmeticOperators() {
	utility.PrintTitle("ARITHMETIC OPERATORS")
	a, b := 10, 5
	utility.PrintNotes("a, b := 10, 5")

	var c int = a + b
	fmt.Print("a + b = ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a - b
	fmt.Print("a - b = ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a * b
	fmt.Print("a * b = ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a / b
	fmt.Print("a / b = ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a / 3
	fmt.Print("a / 3 = ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a % 3
	fmt.Print("a % 3 = ")
	utility.ShowOutput(fmt.Sprint(c))

	d := 7.0 / 2.0
	fmt.Print("7.0 / 2.0 = ")
	utility.ShowOutput(fmt.Sprint(d))
}

func demoComparisonOperators() {
	utility.PrintTitle("COMPARISON OPERATORS")
	a, b := 10, 5
	utility.PrintNotes("a, b := 10, 5")

	var c bool = a == b
	fmt.Print("a == b: ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a != b
	fmt.Print("a != b: ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a < b
	fmt.Print("a < b: ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a <= b
	fmt.Print("a <= b: ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a > b
	fmt.Print("a > b: ")
	utility.ShowOutput(fmt.Sprint(c))

	c = a >= b
	fmt.Print("a >= b: ")
	utility.ShowOutput(fmt.Sprint(c))
}
