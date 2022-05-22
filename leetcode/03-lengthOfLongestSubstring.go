package leetcode

func lengthOfLongestSubstring(s string) int {
	window:=make(map[byte]int)
	left,right:=0,0
	var res int
	for i:=0;i<len(s);i++{
		c:=s[right]
		right++
		window[c]++
		for window[c]>1{
			d:=s[left]
			left++
			window[d]--
		}
		res=max(res,right-left)
	}
	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
