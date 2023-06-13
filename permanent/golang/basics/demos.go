package basics

import (
	"fmt"
	"reflect"
)

func DemosForBasicGoLang() {
	fmt.Print("\033[H\033[2J")
	PrintHeader("GO CLI APP")
	printOutUsefulResources()
	demoStrings()
	demoNumbers()
	demoBooleans()
	demoErrors()
	demoVariableTypesAndOnlyExplicityCoversions()
	demoArithmeticOperators()
	demoComparisonOperators()
	demoContants()
	demoPointers()
}

func printOutUsefulResources() {
	PrintNotes("https://go.dev/ref/spec")
}

func demoStrings() {
	PrintTitle("STRINGS")
	PrintNotes(`
	Simple Data Types (only 4)
		- Strings
		- Numbers
		- Booleans
		- Errors	
	`)

	singleCharacter := 'a'
	fmt.Print(singleCharacter)
	AddSeparator()
	var interpretedString = "This is an interpreted \nstring with a newline character"
	fmt.Println(interpretedString)
	AddSeparator()
	var rawString = `This is a raw \n string with a newline character`
	fmt.Println(rawString)
	AddSeparator()
	var multiLineRawString = `This is a
	Multi Line
		Raw String
	`
	fmt.Println(multiLineRawString)
	AddSeparator()
}

func demoNumbers() {
	PrintTitle("NUMBERS")
	PrintNotes(`
	Note: There are more types than this
		int
		uint - unsigned integers
		float32 | float64 - floating point number
		complext64 | complex128 - complex numbers
	`)
}

func demoBooleans() {
	PrintTitle("BOOLEANS")
	PrintNotes(`
	 	Booleans
		true | false
	`)
}

func demoErrors() {
	PrintTitle("ERRORS")
	PrintNotes(`
		Error Types are Unique to Golang as they are premitive datatype	
	`)
}

func demoVariableTypesAndOnlyExplicityCoversions() {
	PrintTitle("VARIABLES, TYPES, CONVE")
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
	PrintTitle("ARITHMETIC OPERATORS")
	a, b := 10, 5
	PrintNotes("a, b := 10, 5")

	var c int = a + b
	fmt.Print("a + b = ")
	ShowOutput(fmt.Sprint(c))

	c = a - b
	fmt.Print("a - b = ")
	ShowOutput(fmt.Sprint(c))

	c = a * b
	fmt.Print("a * b = ")
	ShowOutput(fmt.Sprint(c))

	c = a / b
	fmt.Print("a / b = ")
	ShowOutput(fmt.Sprint(c))

	c = a / 3
	fmt.Print("a / 3 = ")
	ShowOutput(fmt.Sprint(c))

	c = a % 3
	fmt.Print("a % 3 = ")
	ShowOutput(fmt.Sprint(c))

	d := 7.0 / 2.0
	fmt.Print("7.0 / 2.0 = ")
	ShowOutput(fmt.Sprint(d))
}

func demoComparisonOperators() {
	PrintTitle("COMPARISON OPERATORS")
	a, b := 10, 5
	PrintNotes("a, b := 10, 5")

	var c bool = a == b
	fmt.Print("a == b: ")
	ShowOutput(fmt.Sprint(c))

	c = a != b
	fmt.Print("a != b: ")
	ShowOutput(fmt.Sprint(c))

	c = a < b
	fmt.Print("a < b: ")
	ShowOutput(fmt.Sprint(c))

	c = a <= b
	fmt.Print("a <= b: ")
	ShowOutput(fmt.Sprint(c))

	c = a > b
	fmt.Print("a > b: ")
	ShowOutput(fmt.Sprint(c))

	c = a >= b
	fmt.Print("a >= b: ")
	ShowOutput(fmt.Sprint(c))
}

func demoContants() {
	PrintTitle("CONSTANTS or CONSTANT EXPRESSION")
	const a = 42
	ShowOutput(fmt.Sprint(a))

	const b string = "Hello"
	ShowOutput(b)

	const c = a
	ShowOutput(fmt.Sprint(c))

	const (
		d  = true
		pi = 3.14
	)
	ShowOutput(fmt.Sprint(d))
	ShowOutput(fmt.Sprint(pi))

	// constant expression
	const (
		s = "foo"
		t // will get assigned foo from above
	)
	ShowOutput(fmt.Sprint(s))
	ShowOutput(fmt.Sprint(t))

	const iota = iota
	ShowOutput(fmt.Sprint(iota))
	PrintNotes(`
	const (
		_ iota is 0
		b = iota is 1
		c is 1
		d = 3 * iota is 6
		Note: iota resets to 0 or whatever position in a separate const block. 
		iota is just a an index of a value in const block if it was an array or a list with index that is 0 based		
	)`)
}

func demoPointers() {
	PrintTitle("POINTERS, VALUES, REFERNCES")
	PrintNotes(`	
		a := 42
		b := a
	`)
	a := 42
	b := a
	ShowOutput(fmt.Sprint(a))
	ShowOutput(fmt.Sprint(&a))
	ShowOutput(fmt.Sprint(b))
	ShowOutput(fmt.Sprint(&b))

	PrintNotes(`
		Memory address of the above are different because it is copied by value. 
		- Pointers are primarily used to share memory. USE copies whevern possible. 
		- Why? there are potential bugs that you can run into when sharing memory with pointers. (i.e. race conditions)
	`)

	as := "foo"
	bs := &as // returns an mem address
	ShowOutput(as)
	ShowOutput(fmt.Sprint(bs))

	*bs = "bar" //dereference a pointer with asterisk
	ShowOutput(as)
	// new keyword will always create a new memory address
}
