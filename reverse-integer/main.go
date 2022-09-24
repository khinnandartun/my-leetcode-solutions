package main

import (
	"fmt"
)

func reverse(x int) int {
	var result int
	result = 0
	if x < 0 {
		if x%10 == 0 {
			for x%10 == 0 {
				x = x / 10
			}
		}
		for x <= -10 {
			result = result*10 + (x % 10)
			x = x / 10
		}
	} else if x == 0 {
		return x
	} else {
		if x%10 == 0 {
			for x%10 == 0 {
				x = x / 10
			}
		}
		for x >= 10 {
			result = result*10 + (x % 10)
			x = x / 10
			fmt.Println(result)
			fmt.Println(x)
		}
	}

	result = result*10 + x
	max := 2147483647
	min := -2147483648
	if min >= result {
		return 0
	} else if result >= max {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(-102))

}
