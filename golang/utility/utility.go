package utility

import (
	"strings"

	"github.com/fatih/color"
)

func PrintTitle(title string) {
 	boldType := color.New(color.BgYellow, color.FgBlack, color.Bold)

	color.Yellow(strings.Repeat("-", 20))
	boldType.Println(title)
	color.Yellow(strings.Repeat("-", 20))
}

func PrintSubtitle(subtitle string) {
	color.Magenta(subtitle)
}

func PrintNotes(notes string) {
	color.HiCyan(notes)
}