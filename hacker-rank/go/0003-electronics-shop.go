package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */

	maxSpent := int32(-1)

	for _, kbPrice := range keyboards {
		for _, drPrice := range drives {
			total := kbPrice + drPrice
			if total <= b && total >= maxSpent {
				maxSpent = total
			}
		}
	}

	return maxSpent
}

func main() {
	// Sample Input 0
	b1 := int32(10)
	kb1 := []int32{3, 1}
	dr1 := []int32{5, 2, 8}
	result1 := getMoneySpent(kb1, dr1, b1)
	fmt.Printf("Input: budget = %v, keyboards = %v, drives = %d\nOutput: %v\n", b1, kb1, dr1, result1)

	// Sample Input 1
	b2 := int32(5)
	kb2 := []int32{4}
	dr2 := []int32{5}
	result2 := getMoneySpent(kb2, dr2, b2)
	fmt.Printf("Input: budget = %v, keyboards = %v, drives = %d\nOutput: %v\n", b2, kb2, dr2, result2)

}
