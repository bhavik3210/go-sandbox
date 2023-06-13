package algorithms

import (
	"fmt"

	"rec.golang/utils"
)

func DemoSelectionSort() {
	utils.PrintTitle(" SELECTION SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted: %v", arr)
	fmt.Println()

	selectionSort(&arr)

	utils.ShowOutput(fmt.Sprintf("Sorted: %v", arr))
	fmt.Println()
}

func selectionSort(array *[]int) {
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
