package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var hashmap = make(map[int]int)

func build(inorder []int, postorder []int, inLeft int, inRight int, postRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	nodeVal := postorder[postRight]

	node := &TreeNode{Val: nodeVal}

	inMid := hashmap[nodeVal]

	length := inRight - inMid + 1

	node.Left = build(inorder, postorder, inLeft, inMid-1, postRight-length)

	node.Right = build(inorder, postorder, inMid+1, inRight, postRight-1)

	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	for idx, num := range inorder {
		hashmap[num] = idx
	}
	return build(inorder, postorder, 0, len(inorder)-1, len(inorder)-1)
}
