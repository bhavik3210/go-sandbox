package basics

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func DemosForCollectionsInGoLang() {
	fmt.Print("\033[H\033[2J")
	PrintHeader("Collections")
	// demoArray()
	// demoSlices()
	// demoMaps()
	demoStructs()
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
	var orgMap map[string][]string
	fmt.Println(orgMap)

	orgMap = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappppuccino"},
		"tea":    {"HotTea", "ChaiTea", "ChaiLatte"},
	}

	fmt.Println(orgMap)

	fmt.Println(orgMap["coffee"])

	AddSeparator()
	orgMap["other"] = []string{"HotCoco"}
	fmt.Println(orgMap)

	delete(orgMap, "tea")
	fmt.Println(orgMap)
	AddSeparator()

	fmt.Println(orgMap["tea"])
	v, ok := orgMap["tea"]
	fmt.Println(v, ok)
	AddSeparator()

	copiedMap := orgMap
	copiedMap["coffee"] = []string{"Coffee"}
	orgMap["tea"] = []string{"HotTea"}

	fmt.Println(orgMap)
	fmt.Println(copiedMap)
	PrintNotes(` Maps are copied by reference, so if you apply changes to one the copy will also have that change since they are essentially poiting to the same mem address`)
}

func demoStructs() {
	PrintTitle("STRUCTS")
	PrintNotes("Copied By Value")
	AddSeparator()
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.54, "medium": 2.12, "large": 3.30}},
		{name: "Tea", prices: map[string]float64{"small": 1.54, "medium": 2.12, "large": 3.30}},
	}
	fmt.Println(menu)
}
