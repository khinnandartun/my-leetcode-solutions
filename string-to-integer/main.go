package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {
	var result int
	var nFound bool
	var sfound bool
	var ifound bool
	for _, value := range s {
		x, err := strconv.Atoi(string(value))
		if err == nil {
			result = result*10 + x
			ifound = true
			if result > int(math.Pow(2, 31)-1) {
				result = int(math.Pow(2, 31) - 1)
				if nFound {
					result = result + 1
				}
			}
			if result < int(math.Pow(-2, 31)) {
				fmt.Println("less than")
				result = int(math.Pow(-2, 31)) - 1
			}
		} else if ifound {
			if unicode.IsLetter(value) || value == '.' || value == '-' || value == '+' || value == ' ' {
				break
			}
		} else {
			if sfound || nFound {
				result = 0
				break
			}
			switch value {
			case '-':
				nFound = true
				sfound = true

			case '+':
				sfound = true
			}
			if unicode.IsLetter(value) || value == '.' {
				break
			}
		}
	}
	if nFound {
		result = result * (-1)
	}
	if result > int(math.Pow(2, 31)-1) {
		result = int(math.Pow(2, 31) - 1)
	}
	if result < int(math.Pow(-2, 31)) {
		fmt.Println("less than")
		result = int(math.Pow(-2, 31)) - 1
	}
	return result
}
func main() {
	// fmt.Println(myAtoi("3.235"))
	// fmt.Println(myAtoi("0000-32gerg")) //0
	// fmt.Println(myAtoi("+42"))
	// fmt.Println(myAtoi("+-42"))
	// fmt.Println(myAtoi("words and 987"))
	// fmt.Println(myAtoi(".1"))
	// fmt.Println(myAtoi("     -1"))
	// fmt.Println(myAtoi("  -0012a42"))          // -12
	// fmt.Println(myAtoi("  +0 123"))            // 0
	//fmt.Println(myAtoi("9223372036854775808")) // 2147483647
	fmt.Println(myAtoi("-91283472332")) // 	-2147483648
}
