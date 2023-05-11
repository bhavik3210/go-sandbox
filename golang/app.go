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
	basics.DemoBranching()
	// in := bufio.NewReader(os.Stdin)
	// inputText, _ := in.ReadString('\n')
	// inputText = strings.TrimSpace(inputText)
	// basics.PrintTitle(fmt.Sprint(inputText))
}
