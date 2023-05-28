package basics

import "fmt"

func DemoLooping() {
	PrintHeader("LOOPING")
	demoLoopWithBreaks()
	demoLoopWithCondition()
	demoLoopWithCounter()
	demoLoopWithCollections()
}

func demoLoopWithBreaks() {
	PrintTitle("LOOPING BREAK")
	i := 99
	for {
		fmt.Println(i)
		i++
		break
	}
}

func demoLoopWithCondition() {
	PrintTitle("LOOP WITH CONDITIONS")
	i := 6
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func demoLoopWithCounter() {
	PrintTitle("LOOP WITH COUNTER")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func demoLoopWithCollections() {
	PrintTitle("LOOP WITH COLLECTIONS")

	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d", i)
		fmt.Print(" -> ")
		fmt.Printf("Value: %d", v)
		AddSeparator()
	}
	fmt.Println("Done!")
	PrintNotes(`Works same for slices and maps`)
}
