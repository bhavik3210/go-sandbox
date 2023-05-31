package basics

import (
	"errors"
	"fmt"
)

func DemoErrorManagement() {
	PrintHeader("Error Management")
	DemoSimpleErrors()
	DemoWithErrorReturn()
	DemoConvertPanicWithError()
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

func DemoWithErrorReturn() {
	PrintTitle(" ERROR AS RETURN VALUE ")
	err := checkIfValueIsNegative(-1)
	if err != nil {
		fmt.Println(err)
	}

	err2 := checkIfValueIsNegative(1)
	if err2 != nil {
		fmt.Println(err)
	}
}

func checkIfValueIsNegative(value int) error {
	if value < 0 {
		return errors.New("value is less than -1")
	} else {
		return nil
	}
}

func DemoConvertPanicWithError() {
	PrintTitle(" CONVERT PANIC WITH ERROR ")
	// a := divideWithPossiblePanic(1, 0)
	b, errB := divideWithoutPanicButError(1, 0)
	c, errC := divideWithoutPanicButError(10, 2)

	// fmt.Println(a)
	fmt.Println(b)
	fmt.Println(errB)
	fmt.Println(c)
	fmt.Println(errC)

	d, errD := divideWithDeferAndErrorValue(1, 0)
	if errD != nil {
		fmt.Println(errD)
		return
	}
	fmt.Println("result: ", d)
}

// func divideWithPossiblePanic(l, r int) int {
// 	return l / r
// }

func divideWithoutPanicButError(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("divisor cannot be 0")
	} else {
		return l / r, nil
	}
}

func divideWithDeferAndErrorValue(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
