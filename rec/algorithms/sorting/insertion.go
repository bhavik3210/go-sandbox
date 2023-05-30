package algorithms

import (
	"fmt"

	"rec.golang/utils"
)

func DemoInsertionSort() {
	utils.PrintTitle(" QUICK SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted: %v", arr)
	fmt.Println()

	insertionSort(&arr)

	utils.ShowOutput(fmt.Sprintf("Sorted: %v", arr))
	fmt.Println()
}

func insertionSort(array *[]int) {
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
