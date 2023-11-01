package leetcode

// preorder NLR
// inorder LNR
// postorder LRN

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion
var hashmap = make(map[int]int)

func build(preorder []int, inorder []int, preLeft int, inLeft int, inRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	node := &TreeNode{Val: preorder[preLeft]}

	inMid := hashmap[preorder[preLeft]]

	length := inMid - inLeft + 1

	node.Left = build(preorder, inorder, preLeft+1, inLeft, inMid-1)

	node.Right = build(preorder, inorder, preLeft+length, inMid+1, inRight)

	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for idx, num := range inorder {
		hashmap[num] = idx
	}
	return build(preorder, inorder, 0, 0, len(inorder)-1)
}

// iteration
// to complete
func buildTree(preorder []int, inorder []int) *TreeNode {

}
