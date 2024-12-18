package main

import (
	"fmt"
)

// func isPalindrome(x int) bool {

// 	str := strconv.Itoa(x)
// 	var result []string

// 	for _, char := range str {
// 		digit := string(char)
// 		result = append(result, digit)
// 	}

// 	for idx := range result {
// 		// fmt.Println("result[idx]", result[idx])
// 		// fmt.Println("result[len(result)-1]", result[len(result)-idx-1])
// 		if result[idx] != result[len(result)-idx-1] {
// 			return false
// 		}
// 	}

// 	return true

// }

// Optimized Solution

func isPalindrome(x int) bool {

	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

func main() {
	x1 := 121
	result1 := isPalindrome(x1)
	fmt.Printf("Input: x = %d\nOutput: %v\n", x1, result1)
	x2 := -121
	result2 := isPalindrome(x2)
	fmt.Printf("Input: x = %d\nOutput: %v\n", x2, result2)
	x3 := 1000021
	result3 := isPalindrome(x3)
	fmt.Printf("Input: x = %d\nOutput: %v\n", x3, result3)

}
