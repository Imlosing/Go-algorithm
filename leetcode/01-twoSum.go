package leetcode

func twoSum(nums []int, target int) []int {
	hmap:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		tmp:=target-nums[i]
		if v,ok:=hmap[tmp];ok{
			return []int{i,v}
		}
		hmap[nums[i]]=i
	}
	return nil
}