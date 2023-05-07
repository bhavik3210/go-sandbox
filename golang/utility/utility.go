package utility

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintHeader(header string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println(header)
	fmt.Println(strings.Repeat("=", 20))
}

func PrintTitle(title string) {
	boldType := color.New(color.BgYellow, color.FgBlack, color.Bold)
	fmt.Println()
	// color.Yellow(strings.Repeat("-", 20))
	boldType.Println(title)
	// color.Yellow(strings.Repeat("-", 20))
}

func PrintSubtitle(subtitle string) {
	color.Magenta(subtitle)
}

func PrintNotes(notes string) {
	color.HiCyan(notes)
}

func AddSeparator() {
	fmt.Println()
}

func ShowOutput(output string) {
	color.White(output)
}
