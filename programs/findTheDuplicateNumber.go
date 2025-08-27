package programs

// 287. Find the Duplicate Number

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and using only constant extra space.

// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2

// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3

// Example 3:
// Input: nums = [3,3,3,3,3]
// Output: 3

// Constraints:

// 1 <= n <= 105
// nums.length == n + 1
// 1 <= nums[i] <= n
// All the integers in nums appear only once except for precisely one integer which appears two or more times.

// Approach 1 :
// Use a map of int to track that the current number is present in this map or not
// and if we get found into the map then return the number
func FindDuplicateSol1(nums []int) int {
	numsMap := make(map[int]bool)
	for _, val := range nums {
		if _, found := numsMap[val]; found {
			return val
		}
		numsMap[val] = true
	}
	return 0
}
