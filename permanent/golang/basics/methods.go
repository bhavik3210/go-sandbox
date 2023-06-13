package basics

import "fmt"

func DemoMethods() {
	var number myInt = 23
	fmt.Println(number.demoIsEvenMethod())

	var integerNum int = 23
	fmt.Println(demoIsEvenNotMethod(integerNum))
}

/*
Method functions are extensions in kotlin.
you also use the variable name instead of this keyword
*/
type myInt int // need a type to bind a method to, doesn't have to be a struct
func (i myInt) demoIsEvenMethod() bool { //method receiver
	return int(i)%2 == 0
}

func demoIsEvenNotMethod(i int) bool {
	return i%2 == 0
}

/*
Method functions with value receiver and pointer receiver
*/
type user struct {
	id       int
	username string
}

// Passed by value
func (u user) printFormattedString() string {
	return fmt.Sprintf("%v (%v) \n", u.username, u.id)
}

// Passed by reference via a pointer so value located at this memory address will be mutated
func (u *user) updateUserWithName(newUserName string) {
	u.username = newUserName
}
