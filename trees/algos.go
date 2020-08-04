/*

Notes:

1. Pass arguments for struct vars as reference; its easier to check for nil references
2. ALWAYS: Check node for nil-ness as well as Left and/or Right to prevent referencing like this (nil).left that causes panic

*/
package trees

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val int
}

func Algos() {
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(nil))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(nil))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(nil))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(nil, 5))
	root := TreeNode{Left: nil, Right: nil, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 1 exists?", hasPathSum(&root, 1))
	left := TreeNode{Left: nil, Right: nil, Val: 2}
	root = TreeNode{Left: &left, Right: nil, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	right := TreeNode{Left: nil, Right: nil, Val: 3}
	root = TreeNode{Left: nil, Right: &right, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	root = TreeNode{Left: &left, Right: &right, Val: 1}
	root.Right.Val = 2
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	leftOfleft := TreeNode{Left: nil, Right: nil, Val: 4}
	left = TreeNode{Left: &leftOfleft, Right: nil, Val: 2}
	root = TreeNode{Left: &left, Right: &right, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	rightOfright := TreeNode{Left: nil, Right: nil, Val: 5}
	right = TreeNode{Left: nil, Right: &rightOfright, Val: 3}
	root = TreeNode{Left: &left, Right: &right, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	leftOfright := TreeNode{Left: nil, Right: nil, Val: 6}
	right = TreeNode{Left: &leftOfright, Right: &rightOfright, Val: 3}
	root = TreeNode{Left: &left, Right: &right, Val: 1}
	fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
	fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
	rightOfleft := TreeNode{Left:nil, Right: nil, Val: 6}
	left = TreeNode{Left:nil, Right: &rightOfleft, Val: 2}
	right = TreeNode{Left: &leftOfright, Right: nil, Val: 2}
        fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
        fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
        fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
        left = TreeNode{Left:nil, Right: nil, Val: 2}
        right = TreeNode{Left: &leftOfright, Right: nil, Val: 2}
        fmt.Println("Minimum depth of BTree: ", findMinDepthOfBTree(&root))
        fmt.Println("Maximum depth of BTree: ", findMaxDepthOfBTree(&root))
        fmt.Println("Is Tree symmetrical? ", isTreeSymmetrical(&root))
	fmt.Println("Does Path to sum 5 exists?", hasPathSum(&root, 5))
}

func findMinDepthOfBTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isTreeNodeLeaf(root) {
		return 1
	} else if isTreeNodeLeaf(root.Left) || isTreeNodeLeaf(root.Right) {
		return 2
	} else {
		if root.Left != nil && root.Right != nil {
			leftTree := 1 + findMinDepthOfBTree(root.Left)
			rightTree := 1 + findMinDepthOfBTree(root.Right)
			if leftTree > rightTree {
				return rightTree
			} else {
				return leftTree
			}
		} else if root.Left != nil {
			return 1 + findMinDepthOfBTree(root.Right)
		} else {
			return 1 + findMinDepthOfBTree(root.Left)
		}
	}
}

func findMaxDepthOfBTree(node *TreeNode)(int){

	if node == nil {
                return 0
        }
        if isTreeNodeLeaf(node) {
                return 1
        } else if isTreeNodeLeaf(node.Left) && isTreeNodeLeaf(node.Right) {
                return 2
        }else{
		left := 1 + findMaxDepthOfBTree(node.Left)
		right := 1 + findMaxDepthOfBTree(node.Right)
		if left > right{
			return left
		}else{
			return right
		}
	}
}

func isTreeNodeLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}

	if node.Left == node.Right && node.Left == nil {
		return true
	}
	return false
}

func isTreeSymmetrical(root *TreeNode)(bool) {

	if root == nil || isTreeNodeLeaf(root) {
		return true
	} else {
		return areTwoTreesMirrorTrees(root.Left, root.Right)
	}

}

func areTwoTreesMirrorTrees(left *TreeNode, right *TreeNode)(bool) {
	if left == right && left == nil{
		return true
	}
	if (left == nil && right != nil) || (right == nil && left != nil){
		return false
	}
	if isTreeNodeLeaf(left) && isTreeNodeLeaf(right) {
		return left.Val == right.Val
	} else if isTreeNodeLeaf(left) != isTreeNodeLeaf(right) {
		return false
	}
	//This means both are not leaf nodes

	return left.Val == right.Val && areTwoTreesMirrorTrees(left.Right, right.Left) && areTwoTreesMirrorTrees(left.Left, right.Right)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    
    aSum := 0
    return doesMatchingSumExists(sum, aSum, root)
    
}

func doesMatchingSumExists(eSum int, aSum int, node *TreeNode)(bool){
    if node == nil {
        return false
    }else if isTreeNodeLeaf(node){
        return eSum == aSum + node.Val
    }else{
        aSum = aSum + node.Val
        if node.Left != nil && node.Right == nil && isTreeNodeLeaf(node.Left){
            return eSum == aSum + node.Left.Val
        }else if node.Right != nil && node.Left == nil && isTreeNodeLeaf(node.Right){
                return eSum == aSum + node.Right.Val
        }else if isTreeNodeLeaf(node.Left) && isTreeNodeLeaf(node.Right){
            return eSum == aSum + node.Right.Val || eSum == aSum + node.Left.Val
        }else{
            if node.Right == nil{
                return doesMatchingSumExists(eSum, aSum, node.Left)
            }else if node.Left == nil{
                return doesMatchingSumExists(eSum, aSum, node.Right)
            }else{
            return doesMatchingSumExists(eSum, aSum, node.Left) || doesMatchingSumExists(eSum, aSum, node.Right)
            }
        }
    }
}

