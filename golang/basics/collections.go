package basics

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func DemosForCollectionsInGoLang() {
	fmt.Print("\033[H\033[2J")
	PrintHeader("Collections")
	demoArray()
	demoSlices()
	demoMaps()
}

func demoArray() {
	PrintTitle("ARRAYS")
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

func demoSlices() {
	PrintTitle("SLICES")
	var orgSlice []string //slices don't have set size
	fmt.Println(orgSlice) //default value is empty string, or 0

	orgSlice = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Print("Org Slice: ")
	fmt.Println(orgSlice)

	fmt.Print("Org Slice at 1: ")
	fmt.Println(orgSlice[1])
	orgSlice[1] = "ChaiTea"
	fmt.Print("Org Slice: ")
	fmt.Println(orgSlice)
	cpSlice := orgSlice
	fmt.Print("CP Slice: ")
	fmt.Println(cpSlice)
	cpSlice[1] = "ChaiLatte"
	fmt.Println("Updated CP Slice at 1: ChaiLatte")
	fmt.Print("Org Slice: ")
	fmt.Println(orgSlice)
	fmt.Print("CP Slice: ")
	fmt.Println(cpSlice)

	AddSeparator()

	fmt.Printf("Org Slice Ref: %p", &orgSlice)
	AddSeparator()
	fmt.Printf("CP Slice Ref: %p", &cpSlice)
	AddSeparator()
	AddSeparator()
	for i := 0; i < len(orgSlice); i++ {
		fmt.Printf("[%s]: %p", orgSlice[i], &orgSlice[i])
		AddSeparator()
	}
	AddSeparator()
	for i := 0; i < len(cpSlice); i++ {
		fmt.Printf("[%s]: %p", cpSlice[i], &cpSlice[i])
		AddSeparator()
	}
	AddSeparator()
	orgSlice = append(orgSlice, "HotChocolate", "HotTea")
	fmt.Print("Org Slice: ")
	fmt.Println(orgSlice)
	fmt.Print("CP Slice:")
	fmt.Println(cpSlice)
	PrintNotes(`
		When orgSlice was mutated, cpSlice started a new life and decoupled itself when append(...) was used 
	`)

	for i := 0; i < len(orgSlice); i++ {
		fmt.Printf("[%s]: %p", orgSlice[i], &orgSlice[i])
		AddSeparator()
	}
	AddSeparator()
	for i := 0; i < len(cpSlice); i++ {
		fmt.Printf("[%s]: %p", cpSlice[i], &cpSlice[i])
		AddSeparator()
	}
	AddSeparator()

	slices.Delete(orgSlice, 1, 3)
	fmt.Println(orgSlice)
}

func demoMaps() {
	PrintTitle("MAPS")
}
