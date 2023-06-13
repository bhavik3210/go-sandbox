package main

import "fmt"

func main() {
	fmt.Println("Test")
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	fmt.Println(checkStraightLine([][]int{{0, 0}, {0, 1}, {0, -1}}))
	// [[0,0],[0,1],[0,-1]]
	//[[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
}

func checkStraightLine(c [][]int) bool {

	for i := 0; i < len(c)-1; i++ {
		// fmt.Print(c[i][1])
		// fmt.Print(c[i+1][0])
		// fmt.Println()
		if c[i][1] != c[i+1][0] {
			return false
		}

		if c[i][0] == c[i+1][0]-1 || c[i][1] == 

	}
	return true
}
