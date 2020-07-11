/*

Notes:

1. Pass arguments for struct vars as reference; its easier to check for nil references
2. ALWAYS: Check node for nil-ness as well as left and/or right to prevent referencing like this (nil).left that causes panic

*/
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
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(nil))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(nil))
	root := Node{left: nil, right: nil, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	left := Node{left: nil, right: nil, value: 2}
	root = Node{left: &left, right: nil, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	right := Node{left: nil, right: nil, value: 3}
	root = Node{left: nil, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	root = Node{left: &left, right: &right, value: 1}
	root.right.value = 2
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	leftOfleft := Node{left: nil, right: nil, value: 4}
	left = Node{left: &leftOfleft, right: nil, value: 2}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	rightOfright := Node{left: nil, right: nil, value: 5}
	right = Node{left: nil, right: &rightOfright, value: 3}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	leftOfright := Node{left: nil, right: nil, value: 6}
	right = Node{left: &leftOfright, right: &rightOfright, value: 3}
	root = Node{left: &left, right: &right, value: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	rightOfLeft := Node{left:nil, right: nil, value: 6}
	left = Node{left:nil, right: &rightOfLeft, value: 2}
	right = Node{left: &leftOfright, right: nil, value: 2}
        fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
        fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        left = Node{left:nil, right: nil, value: 2}
        right = Node{left: &leftOfright, right: nil, value: 2}
        fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
        fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
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

func findMaxDepthOfBTree(node *Node)(int){

	if node == nil {
                return 0
        }
        if isNodeLeaf(node) {
                return 1
        } else if isNodeLeaf(node.left) && isNodeLeaf(node.right) {
                return 2
        }else{
		left := 1 + findMaxDepthOfBTree(node.left)
		right := 1 + findMaxDepthOfBTree(node.right)
		if left > right{
			return left
		}else{
			return right
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

func isTreeSymmetrical(root *Node)(bool) {

	if root == nil || isNodeLeaf(root) {
		return true
	} else {
		return areTwoTreesMirrorTrees(root.left, root.right)
	}

}

func areTwoTreesMirrorTrees(left *Node, right *Node)(bool) {
	if left == right && left == nil{
		return true
	}
	if (left == nil && right != nil) || (right == nil && left != nil){
		return false
	}
	if isNodeLeaf(left) && isNodeLeaf(right) {
		return left.value == right.value
	} else if isNodeLeaf(left) != isNodeLeaf(right) {
		return false
	}
	//This means both are not leaf nodes

	return left.value == right.value && areTwoTreesMirrorTrees(left.right, right.left) && areTwoTreesMirrorTrees(left.left, right.right)
}

