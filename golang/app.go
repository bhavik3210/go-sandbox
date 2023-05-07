package main

import (
	"fmt"

	"golang.dojo/basics"
)

func main() {
	fmt.Print("\033[H\033[2J")
	basics.PrintHeader("GO CLI APP")
	// basics.DemosForBasicGoLang()
}
