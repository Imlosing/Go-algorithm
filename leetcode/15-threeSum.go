package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		target := -nums[i]
		ans := twoSum(nums, i+1, target)
		for _, tmp := range ans {
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	var res [][]int
	low := start
	high := len(nums) - 1
	for low < high {
		var path []int
		left := nums[low]
		right := nums[high]
		sum := nums[low] + nums[high]
		if sum == target {
			path = append(path, nums[low], nums[high])
			res = append(res, path)
			for low < high && nums[low] == left {
				low++
			}
			for low < high && nums[high] == right {
				high--
			}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return res
}
