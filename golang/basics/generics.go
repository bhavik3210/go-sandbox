package basics

import "fmt"

func DemoGenerics() {
	// demoGenerics()
	// demoGenerics2()
	demoGenericsWithCustomTypeConstraints()
}

func demoGenerics() {
	testScores := []float64{
		57.3,
		230,
		42.23,
		27,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)
}

func clone[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func demoGenerics2() {
	testScores2 := map[string]float64{
		"Harry":    87.4,
		"Hermione": 102,
		"Ronald":   23.5,
		"Neville":  37,
	}

	c := clone2(testScores2)

	fmt.Println(c)
}

func clone2[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func demoGenericsWithCustomTypeConstraints() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"Hello ", "World!"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
