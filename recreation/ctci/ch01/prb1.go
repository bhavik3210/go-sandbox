package ch01

/*
	1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters.
	What if you cannot use additional data structures?
*/

func IsUnique(str string) bool {
	result := true
	for i := 0; i < len(str); i++ {
		compareChar := string([]rune(str)[i])
		for j := i + 1; j < len(str); j++ {
			currentChar := string([]rune(str)[j])
			if compareChar == currentChar {
				result = false
				break
			}
		}
	}
	return result
}
