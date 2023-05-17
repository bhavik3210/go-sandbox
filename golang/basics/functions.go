package basics

import "fmt"

/*
func functionName(parameters)(return values if any){}
func functionName(incoming values)(outgoing values if any){}
*/
func DemoFunctions() {
	demoParametersAndArguments()
}

func greet(name1 string, name2 string) {
	fmt.Print(name1)
	fmt.Print(",")
	fmt.Println(name2)
}

func conciesParamas(name1, name2 string) {
	fmt.Print(name1)
	fmt.Print(",")
	fmt.Println(name2)
}

func variadicParamenters(names ...string) {
	// names: are gathered as a slice, have to be the last of the parameters list
	for i, n := range names {
		fmt.Print(n)
		if i <= len(names)-2 {
			fmt.Print(",")
		}
	}
}

func demoParametersAndArguments() {
	greet("name1", "name2")
	conciesParamas("name1", "name2")
	variadicParamenters()
	variadicParamenters("name1", "name2", "name3", "name4", "nameN")
}
