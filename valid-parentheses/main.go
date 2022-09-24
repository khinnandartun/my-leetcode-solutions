package main

import "fmt"

func isValid(s string) bool {
	var temp []rune
	validator := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	if len(s)%2 != 0 {
		return false
	}
	for index, value := range s {
		if index == 0 {
			temp = append(temp, value)
		} else {
			if len(temp) > 0 {
				if value == validator[temp[len(temp)-1]] {
					temp = temp[:len(temp)-1]
				}
			} else {
				temp = append(temp, value)
			}
		}

	}
	return len(temp) == 0
}
func main() {
	fmt.Println(isValid("()[]{}")) // "({[)"
}
