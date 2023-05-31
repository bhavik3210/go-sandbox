package main

import (
	"fmt"

	"golang.dojo/basics"
)

func main() {
	fmt.Print("\033[H\033[2J")
	basics.PrintHeader("GO CLI APP")
	// basics.DemosForBasicGoLang()
	// basics.DemosForCollectionsInGoLang()
	// basics.DemoLooping()
	// basics.DemoBranching()
	// basics.DemoFunctions()
	// basics.DemoMethods()
	// basics.DemoInterfaces()
	// basics.DemoGenerics()
	basics.DemoErrorManagement()
}
