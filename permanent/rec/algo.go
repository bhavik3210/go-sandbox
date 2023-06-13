package main

import (
	sort "rec.golang/algorithms/sorting"
	"rec.golang/utils"
)

func DemoSortingAlgo() {
	utils.PrintHeader(" SORTING ALGORITHMS ")
	sort.DemoBubbleSort()
	sort.DemoInsertionSort()
	sort.DemoSelectionSort()
	sort.DemoMergeSort()
}
