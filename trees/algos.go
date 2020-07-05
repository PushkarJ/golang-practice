package trees

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func Algos() {
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(nil))
	root := Node{left: nil, right: nil, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	left := Node{left: nil, right: nil, value: 2}
	root = Node{left: &left, right: nil, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	right := Node{left: nil, right: nil, value: 3}
	root = Node{left: nil, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	leftOfleft := Node{left: nil, right: nil, value: 4}
	left = Node{left: &leftOfleft, right: nil, value: 2}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	rightOfright := Node{left: nil, right: nil, value: 5}
	right = Node{left: nil, right: &rightOfright, value: 3}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
	leftOfright := Node{left: nil, right: nil, value: 6}
	right = Node{left: &leftOfright, right: &rightOfright, value: 3}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
}

func findMinDepthOfBTree(root *Node) int {
	if root == nil {
		return 0
	}
	if isNodeLeaf(root) {
		return 1
	} else if isNodeLeaf(root.left) || isNodeLeaf(root.right) {
		return 2
	} else {
		if root.left != nil && root.right != nil {
			leftTree := 1 + findMinDepthOfBTree(root.left)
			rightTree := 1 + findMinDepthOfBTree(root.right)
			if leftTree > rightTree {
				return rightTree
			} else {
				return leftTree
			}
		} else if root.left != nil {
			return 1 + findMinDepthOfBTree(root.right)
		} else {
			return 1 + findMinDepthOfBTree(root.left)
		}
	}
}

func isNodeLeaf(node *Node) bool {
	if node == nil {
		return false
	}

	if node.left == node.right && node.left == nil {
		return true
	}
	return false
}

/*

Notes:

1. Pass arguments for struct vars as reference; its easier to check for nil references
2. Check node for nil-ness as well as left and/or right to prevent referencing like this (nil).left that causes panic
*/
