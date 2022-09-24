package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var result string
	var tmp string
	if len(s) == 1 {
		result = s
		return len(s)
	}
	for index := 0; index < len(s); index++ {
		if index != len(s)-1 {
			tmp = string(s[index])
			for next := index + 1; next < len(s); next++ {
				if strings.Contains(tmp, string(s[next])) {
					if len(tmp) > len(result) {
						result = tmp
					}
					tmp = ""
					break
				} else {
					tmp = tmp + string(s[next])
					if next == len(s)-1 {
						if len(tmp) > len(result) {
							result = tmp
						}
					}
				}

			}
		}

	}
	fmt.Println(tmp)
	if len(tmp) > len(result) {
		result = tmp
	}
	return len(result)
}
func main() {
	fmt.Println(lengthOfLongestSubstring("bwfu"))
}
