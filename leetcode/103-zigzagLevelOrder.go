package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	flag := 0
	for len(queue) > 0 {
		l := len(queue)
		var p []*TreeNode
		var path []int
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			path = append(path, node.Val)
		}
		if flag%2 != 0 {
			for j, k := 0, len(path)-1; j < k; {
				path[j], path[k] = path[k], path[j]
				j++
				k--
			}
		}
		flag++
		res = append(res, path)
		path = []int{}
		queue = p
	}
	return res
}
