package basics

import "fmt"

func DemoBranching() {
	PrintHeader("BRANCHING")
	// demoIfStatement()
	demoSwitchStatement()
}

func demoIfStatement() {
	PrintTitle("IF STATEMENT")
	i := 5
	if i < 5 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("i is not less than 5")
	}

	if i < 6 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("i is not less than 5")
	}

	// Concise syntax
	if i := 15; i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("is is atleast 10")
	}
}

func demoSwitchStatement() {
	PrintTitle("SWITCH STATEMENTS")

	i := 5
	switch i {
	case 1:
		fmt.Println(("First Case Hit"))
	case 2:
		fmt.Println("Second Case Hits")
	default:
		fmt.Println("Default Case")
	}

	i = 1
	switch i {
	case 1:
		fmt.Println(("First Case Hit"))
	case 2:
		fmt.Println("Second Case Hits")
	default:
		fmt.Println("Default Case")
	}

	i = 2
	switch i {
	case 1:
		fmt.Println(("First Case Hit"))
	case 2:
		fmt.Println("Second Case Hits")
	default:
		fmt.Println("Default Case")
	}
	AddSeparator()

	// Logical Switch
	switch i := 8; true {
	case i < 5:
		fmt.Println("i is less than 5")
	case i < 10:
		fmt.Println("i is less than 10")
	default:
		fmt.Println("i is greater than or equal to 10")
	}
}

func deferredFunctions() {

}

func demoGoToFunctions() {

}
