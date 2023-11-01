package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder
var treeNodeSlice = make([]*TreeNode, 0)

func preorder(node *TreeNode) {
	if node == nil {
		return
	}
	treeNodeSlice = append(treeNodeSlice, node)
	preorder(node.Left)
	preorder(node.Right)
}

func flatten(root *TreeNode) {
	preorder(root)
	p := treeNodeSlice[0]

	for i := 1; i < len(treeNodeSlice); i++ {
		p.Left = nil
		p.Right = treeNodeSlice[i]
		p = p.Right
	}
}

// direct
// Look for the precursor node
func flatten(root *TreeNode) {
	var p = root
	var next, pre *TreeNode
	for p != nil {
		if p.Left == nil {
			p = p.Right
			continue
		}
		next = p.Left
		pre = next
		for pre.Right != nil {
			pre = pre.Right
		}
		pre.Right = p.Right
		p.Right = next
		p.Left = nil
		p = p.Right
	}
}
