package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	reverseNumber := 0
	originalNumber := x
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	for x >= 10 {
		reverseNumber = reverseNumber*10 + x%10
		x = x / 10
		// fmt.Println("reverse ", reverseNumber, "x ", x)
	}
	reverseNumber = reverseNumber*10 + x%10
	// fmt.Println(reverseNumber)
	return reverseNumber == originalNumber

}

func main() {
	result := isPalindrome(1001)
	fmt.Println(result)
}
