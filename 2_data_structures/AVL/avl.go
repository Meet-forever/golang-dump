package main

import "fmt"

type node struct {
	data   int
	height int
	left   *node
	right  *node
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func height(temp *node) int {
	if temp == nil {
		return 0
	}
	return temp.height
}

func rotateLeft(y *node) (x *node) {
	x = y.left
	temp := x.right

	x.right = y
	y.left = temp

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return
}

func rotateRight(y *node) (x *node) {
	x = y.right
	temp := x.left

	x.left = y
	y.right = temp

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return
}

func balanceFactor(temp *node) int {
	return height(temp.left) - height(temp.right)
}

func insert(ele *node, val int) *node {
	if ele == nil {
		return &node{data: val, height: 0}
	}
	if ele.data > val {
		ele.left = insert(ele.left, val)
	} else if ele.data < val {
		ele.right = insert(ele.right, val)
	} else {
		return ele
	}
	ele.height = max(height(ele.left), height(ele.right)) + 1

	bf := balanceFactor(ele)

	if bf > 1 && val < ele.left.data {

		return rotateRight(ele)

	} else if bf < -1 && val > ele.right.data {

		return rotateLeft(ele)

	} else if bf > 1 && val > ele.left.data {

		ele.left = rotateLeft(ele.left)
		return rotateRight(ele.right)

	} else if bf < -1 && val < ele.right.data {

		ele.right = rotateRight(ele.right)
		return rotateLeft(ele.left)

	}

	return ele
}

func printTree(temp *node) {
	if temp == nil {
		return
	}
	printTree(temp.left)
	fmt.Println(temp.data)
	printTree(temp.right)
}

func main() {
	var root *node
	root = insert(root, 3)
	root = insert(root, 2)
	root = insert(root, 1)
	printTree(root)
}
