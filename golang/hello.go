package main

import (
	"fmt"
	"reflect"
)

func Hello() string {
	return "Hello World!"
}

func main() {
	// fmt.Println(Hello());
	demoVariablesAndDataTypes()
}

func demoVariablesAndDataTypes() {
	/* 
		Simple Data Types (only 4)
		- Strings
		- Numbers
		- Booleans
		- Errors
	*/

	// Strings
	//  
	// `this is a raw string \n` newline character is not interpreted and printed out as \n
	// also valid `
	// multiple line strings`
	var interpretedString = "This is a interpreted \n string" 
	var rawString = `this is a raw \n string `
	var multiLineRawString = `
		This is a
		Multi Line
		Raw String
	`


	fmt.Println(interpretedString)
	fmt.Println(rawString)
	fmt.Println(multiLineRawString)


	// Numbers
	/*
			int
			uint - unsigned integers
			float32 | float64 - floating point number
			complext64 | complex128 - complex numbers
	*/

	// Booleans
	/*
		true | false 
	*/

	// Error 
	/*
		error type: 
	*/

	var myNameDeclare string; //preferred
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
	fmt.Printf("%f ",d)
	fmt.Printf("%d ",c)


	//Arithmetics Operators


}