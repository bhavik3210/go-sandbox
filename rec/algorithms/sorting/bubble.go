package algorithms

import (
	"fmt"

	"rec.golang/utils"
)

func DemoBubbleSort() {
	utils.PrintTitle(" BUBBLE SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted: %v", arr)
	fmt.Println()

	bubbleSort(&arr)

	utils.ShowOutput(fmt.Sprintf("Sorted: %v", arr))
	fmt.Println()
}

func bubbleSort(array *[]int) {
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
