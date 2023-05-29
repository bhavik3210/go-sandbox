package main

import "fmt"

func DemoSortingAlgo() {
	PrintHeader("DEMO SORTING ALGORITHMS")
	demoBubbleSort()
}

func demoBubbleSort() {
	PrintTitle("BUBBLE SORT")
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
