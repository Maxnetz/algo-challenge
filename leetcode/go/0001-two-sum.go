package main

// https://leetcode.com/problems/two-sum/
// easy

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for currentIdx, num := range nums {
		previousStoredIdx, exists := idxMap[target-num]
		if exists {
			return []int{previousStoredIdx, currentIdx}
		}
		idxMap[num] = currentIdx
	}
	return nil
}

// func main() {
// 	// Test cases for the twoSum function
// 	nums1 := []int{2, 7, 11, 15}
// 	target1 := 9
// 	result1 := twoSum(nums1, target1)
// 	fmt.Printf("Input: nums = %v, target = %d\nOutput: %v\n", nums1, target1, result1)

// 	nums2 := []int{3, 2, 4}
// 	target2 := 6
// 	result2 := twoSum(nums2, target2)
// 	fmt.Printf("Input: nums = %v, target = %d\nOutput: %v\n", nums2, target2, result2)

// 	nums3 := []int{3, 3}
// 	target3 := 6
// 	result3 := twoSum(nums3, target3)
// 	fmt.Printf("Input: nums = %v, target = %d\nOutput: %v\n", nums3, target3, result3)
// }
