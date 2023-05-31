package basics

import (
	"errors"
	"fmt"
)

func DemoErrorManagement() {
	PrintHeader("Error Management")
	DemoSimpleErrors()
}

func DemoSimpleErrors() {
	PrintTitle(" SIMPLE ERRORS ")

	err := errors.New("this is a first error")
	fmt.Println(err)
	AddSeparator()

	//wrapping above error into another
	err2 := fmt.Errorf("this second error wraps the first error: %w", err)
	fmt.Println(err2)
	AddSeparator()

}
