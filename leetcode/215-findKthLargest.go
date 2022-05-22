package leetcode

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapAdjust(nums, i, len(nums)-1)
	}
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapAdjust(nums, 0, i-1)
	}
	return nums[len(nums)-k]
}

func heapAdjust(nums []int, a, b int) {
	tmp := nums[a]
	for i := 2*a + 1; i <= b; i = i*2 + 1 {
		if i < b && nums[i] < nums[i+1] {
			i++
		}
		if nums[i] < tmp {
			break
		}
		nums[a] = nums[i]
		a = i
	}
	nums[a] = tmp
}
