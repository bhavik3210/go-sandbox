package basics

import "fmt"

func DemoBranching() {
	PrintHeader("BRANCHING")
	// demoIfStatement()
	// demoSwitchStatement()
	// demoDeferredFunctions()
	// demoPanicAndRecover()
	demoGoToFunctions()
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

func demoDeferredFunctions() {
	PrintTitle("DEFFERED FUNCTIONS")
	fmt.Println("Main 1")
	defer fmt.Println("Defer 1")
	fmt.Println("Main 2")
	defer fmt.Println("Defer 2")
	PrintNotes(` Defered calls as are in a stack, last one in first one out
	This is very useful for Datbase operations, opening and closeing DB connections`)

	// db, _ := sql.Open("driverName", "connection string")
	// defer db.close()
	// rows, _ := db.Query("some sql query here")
	// defer rows.clsoe()
	// here rows will close before db will because that would be the correct order
}

func demoPanicAndRecover() {
	PrintTitle("PANIC AND RECOVER")
	fmt.Println("Main 1")
	func1()
	fmt.Println("Main 2")

}

func func1() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("func1 A")
	panic("Oops!")
	fmt.Println("func1 B")
}

func demoGoToFunctions() {
	PrintTitle("GO TO STATEMENTS")
	i := 10
	if i < 15 {
		goto myLabel
		fmt.Println("After goto statement, unreachable")
	}

myLabel:
	j := 32
	for ; i < 15; i++ {
		j--
	}
	fmt.Println(j)
}
