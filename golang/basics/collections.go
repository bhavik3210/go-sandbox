package basics

import (
	"fmt"
)

func DemosForCollectionsInGoLang() {
	fmt.Print("\033[H\033[2J")
	PrintHeader("GO CLI APP: Collections")
	demoArray()
}

func demoArray() {
	var array [3]int
	fmt.Println(array)
	array = [3]int{1, 2, 3}
	fmt.Println(array)
	array[1] = 99
	fmt.Println(array)
	fmt.Println(len(array))

	originalArray := [3]string{"foo", "bar", "baz"}

	copiedArray := originalArray // copied by value
	fmt.Println(copiedArray)

	originalArray[0] = "quux"
	fmt.Println(originalArray)
	fmt.Println(copiedArray)

	fmt.Println(originalArray == copiedArray)
}
