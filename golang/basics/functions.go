package basics

import "fmt"

/*
func functionName(parameters)(return values if any){}
func functionName(incoming values)(outgoing values if any){}
*/
func DemoFunctions() {
	demoParametersAndArguments()
	demoFunctionWithMultipleReturnTypes()
}

func greet(name1 string, name2 string) {
	fmt.Print(name1)
	fmt.Print(",")
	fmt.Println(name2)
}

func conciseParams(name1, name2 string) {
	fmt.Print(name1)
	fmt.Print(",")
	fmt.Println(name2)
}

func variadicParameters(names ...string) {
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
	conciseParams("name1", "name2")
	variadicParameters()
	variadicParameters("name1", "name2", "name3", "name4", "nameN")
}

func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}

func demoFunctionWithMultipleReturnTypes() {
	PrintTitle("Multiple Return Types/Values")
	result, ok := divide(1, 2)
	if ok {
		fmt.Println(result)
	}
}
