package algorithms

import (
	"fmt"

	"rec.golang/utils"
)

func DemoMergeSort() {
	utils.PrintTitle(" MERGE SORT ")
	arr := []int{23, 1, 28, 50, 2, 5, 10}

	fmt.Printf("Unsorted: %v", arr)
	fmt.Println()

	mergeSort(arr)

	utils.ShowOutput(fmt.Sprintf("Sorted: %v", arr))
	fmt.Println()
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

}
