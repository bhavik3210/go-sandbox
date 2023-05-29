package main

import "fmt"

func DemoSortingAlgo() {
	PrintHeader("DEMO SORTING ALGORITHMS")
	// demoBubbleSort()
	// demoInsertionSort()
	demoSelectionSort()
}

func demoSelectionSort() {
	PrintTitle(" SELECTION SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted Array: %v", arr)
	fmt.Println()

	SelectionSort(&arr)

	fmt.Printf("Sorted Array: %v", arr)
	fmt.Println()
}

func SelectionSort(array *[]int) {
	arr := *array

	for i := 0; i < len(arr); i++ {
		l := i
		for j := i; j < len(arr); j++ {
			if arr[l] > arr[j] {
				l = j
			}
		}

		if arr[l] != arr[i] {
			tempValue := arr[i]
			arr[i] = arr[l]
			arr[l] = tempValue
		}
	}
}

func demoInsertionSort() {
	PrintTitle(" QUICK SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted Array: %v", arr)
	fmt.Println()

	InsertionSort(&arr)

	fmt.Printf("Sorted Array: %v", arr)
	fmt.Println()
}

func InsertionSort(array *[]int) {
	arr := *array
	for i := 0; i < len(arr)-1; i++ {

		for j := i + 1; j > 0; j-- {
			lValue := arr[j-1]
			rValue := arr[j]

			if lValue > rValue {
				arr[j-1] = rValue
				arr[j] = lValue
			}
		}
	}
}

func demoBubbleSort() {
	PrintTitle(" BUBBLE SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted Array: %v", arr)
	fmt.Println()

	BubbleSort(&arr)

	fmt.Printf("Sorted Array: %v", arr)
	fmt.Println()
}

func BubbleSort(array *[]int) {
	arr := *array
	for i := 0; i < len(arr)-1; i++ {
		isSwapped := false

		for j := 0; j < len(arr)-1; j++ {
			lValue := arr[j]
			rValue := arr[j+1]

			if lValue > rValue {
				arr[j] = rValue
				arr[j+1] = lValue
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}
}
