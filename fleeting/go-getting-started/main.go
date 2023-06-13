package main

import (
	"errors"
	"fmt"
)

func main() {
	// valueTypesDemo()
	// pointerTypesDemo()
	// dataTypeArraysDemo()
	// dataTypeSliceDemo()
	// dataTypeMapDemo()
	// dataTypeStructDemo()

	// _, err := startWebServer(3000, 2)
	// fmt.Println(err)

	// loopThatTerminatesBasedOnCondition()
	// loopConditionPostClause()
	// loopInfinite()
	// loopWithCollections()

	// panic("Something bad just happened")

	// ifStatmentDemo()

	switchStatementDemo()
}

type HTTPRequest struct {
	method string
}

func switchStatementDemo() {
	GET := "GET"
	DELETE := "DELETE"
	POST := "POST"
	PUT := "PUT"

	request := HTTPRequest{
		method: GET,
	}

	switch request.method {
	case GET:
		fmt.Println("GET Request")
	case DELETE:
		fmt.Println("DELETE Request")
	case POST:
		fmt.Println("POST Request")
	case PUT:
		fmt.Println("PUT Request")
	default:
		fmt.Println("Unhandled method")
	}
}

type User struct {
	ID        int
	firstName string
	lastName  string
}

func ifStatmentDemo() {
	u1 := User{
		ID:        1,
		firstName: "Arthur",
		lastName:  "Dent",
	}

	u2 := User{
		ID:        2,
		firstName: "Ford",
		lastName:  "Prefect",
	}

	if u1.ID == u2.ID {
		fmt.Println("Same user!")
	} else {
		fmt.Println("Different Users!")
	}
}

func loopWithCollections() {
	slice := []int{1, 2, 3}
	// for i := 0; i < len(slice); i++ {
	for i, v := range slice { // same with map with key,value pairs
		fmt.Println(i, v)
	}
}

func loopInfinite() {
	var i int
	for {
		if i == 5 {
			break
		}
		fmt.Print(i)
		i++
	}
}

func loopConditionPostClause() {
	for i := 0; i < 5; i++ {
		fmt.Printf("[%d]", i)
	}
}

func loopThatTerminatesBasedOnCondition() {
	var i int
	for i < 5 {
		fmt.Printf("[%d]", i)
		i++
		if i == 3 {
			continue
		}
	}
}

func valueTypesDemo() {

	printSeparator("Value Types")

	var i int = 32
	fmt.Println(i)

	var f float32 = 23.32
	fmt.Println(f)

	firstName := "Arthur" // := means compiler will implicitly apply type for user
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 5)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	printBottomSeparator()
}

func pointerTypesDemo() {

	printSeparator("Pointer Types")

	// Pointer types
	firstName := "Arthur"
	var firstNamePointer *string = new(string) //new(string) allocates memory for string type
	*firstNamePointer = "Arthur"

	fmt.Println(firstNamePointer)  //prints mem address
	fmt.Println(*firstNamePointer) //prints value

	// do not use pointer arthimetics in go because it is not supported by golang
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)

	printBottomSeparator()
}

func dataTypeArraysDemo() {
	printSeparator("Data Type: Arrays")

	// Data Type: Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	arr2 := [3]int{4, 5, 6}

	fmt.Println(arr)
	fmt.Println(arr[2])

	fmt.Println(arr2)

	printBottomSeparator()
}

func dataTypeSliceDemo() {

	printSeparator("Data Type: Slice")

	// Data Type: Slice
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	slice := arr[:] // slice is pointing to underlying arr therefore change in one reflects on other
	fmt.Println(arr, slice)

	arr[1] = 32
	slice[1] = 47
	fmt.Println(arr, slice)

	sliceWOutArray := []int{1, 2, 3}
	fmt.Println(sliceWOutArray)
	sliceWOutArray = append(sliceWOutArray, 5) // sliceWOutArray is like arraylist and this will assing a new array and gc will collect
	fmt.Println(sliceWOutArray)

	sliceWOutArray = append(sliceWOutArray, 6, 7, 8)

	s2 := sliceWOutArray[1:]
	s3 := sliceWOutArray[:2]
	s4 := sliceWOutArray[1:2]

	fmt.Println(s2, s3, s4)

	printBottomSeparator()
}

func dataTypeMapDemo() {
	printSeparator("Data Type: Map")
	// Map
	m := map[string]int{"foo": 43}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
	printBottomSeparator()
}

func dataTypeStructDemo() {

	printSeparator("Data Type: Structs")
	// Structs
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Morgan"
	fmt.Println(u)

	u2 := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u2)
	printBottomSeparator()
}

func printSeparator(label string) {
	fmt.Println("")
	fmt.Println("--------------------------------")
	fmt.Println(label)
	fmt.Println("--------------------------------")
}

func printBottomSeparator() {
	fmt.Println("--------------------------------")
	fmt.Println("")
}

/*
error: is a pointer data type
*/
func startWebServer(port, numberOfRetries int) (int, error) { //go applies int to both port and numberOfRetries
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server Started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("Something went wrong")
}
