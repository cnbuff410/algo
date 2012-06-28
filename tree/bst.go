// Author: likunarmstrong@gmail.com

// Binary tree implementation

package main

import "fmt"

var (
	maxPathLength int
)

type node struct {
	left  *node
	right *node
	// Custom
	value int
}

type BSTree struct {
	root *node
	num  int
}

func insert(n *node, v int) {
	if v > n.value { // Right
		if n.right == nil {
			n.right = &node{nil, nil, v}
			return
		} else {
			insert(n.right, v)
		}
	} else { // Left
		if n.left == nil {
			n.left = &node{nil, nil, v}
			return
		} else {
			insert(n.left, v)
		}
	}
}

func search(n *node, v int) bool {
	if n == nil {
		return false
	}
	if n.value == v {
		return true
	}
	return search(n.left, v) || search(n.right, v)
}

/* Find the left most node under current tree */
func minValue(n *node) int {
	if n.left == nil {
		return n.value
	}
	return minValue(n.right)
}

func del(n *node, parent *node, v int) bool {
	switch {
	case n.value > v:
		if n.left == nil {
			return false
		}
		return del(n.left, n, v)
	case n.value < v:
		if n.right == nil {
			return false
		}
		return del(n.right, n, v)
	case n.value == v:
		if n.left != nil && n.right != nil {
			n.value = minValue(n.right)
			return del(n.right, n, n.value)
		} else if parent.left == n {
			if n.left != nil {
				parent.left = n.left
			} else {
				parent.left = n.right
			}
		} else if parent.right == n {
			if n.left != nil {
				parent.right = n.left
			} else {
				parent.right = n.right
			}
		}
		return true
	}
	return false
}

func inOrderTraverse(n *node) []int {
	l := make([]int, 0)
	if n == nil {
		return l
	}
	l1 := inOrderTraverse(n.left)
	l = append(l, l1...)
	l = append(l, n.value)
	l2 := inOrderTraverse(n.right)
	l = append(l, l2...)
	return l
}

func printPreOrder(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%d\n", n.value)
	printPreOrder(n.left)
	printPreOrder(n.right)
}

func printInOrder(n *node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Printf("%d\n", n.value)
	printInOrder(n.right)
}

/* Calculate the longest path in the tree */
func calLength(n *node) int {
	var length, localmaxPathLength int
	switch {
	case n.left == nil && n.right == nil:
		if 1 > maxPathLength {
			maxPathLength = 1
		}
		return 1
	case n.left == nil:
		length = calLength(n.right)
		localmaxPathLength = length + 1 // Single branch
		if localmaxPathLength > maxPathLength {
			maxPathLength = localmaxPathLength
		}
		return localmaxPathLength
	case n.right == nil:
		length = calLength(n.left)
		localmaxPathLength = length + 1 // Single branch
		if localmaxPathLength > maxPathLength {
			maxPathLength = localmaxPathLength
		}
		return localmaxPathLength
	default: // n with two children
		length1 := calLength(n.left)
		length2 := calLength(n.right)
		if length1+length2+2 > maxPathLength {
			maxPathLength = length1 + length2 + 2
		}
		// We only return the longer one
		if length1 > length2 {
			return length1 + 1
		}
		return length2 + 1
	}
	return 0
}

/******************Public method********************/

func NewBSTree(v int) *BSTree {
	return &BSTree{&node{nil, nil, v}, 0}
}

/* Inser a new node */
func (bst *BSTree) Insert(v int) {
	bst.num++
	insert(bst.root, v)
}

/* Check if a node with given value is in the tree */
func (bst *BSTree) Exists(v int) bool {
	bst.num--
	return search(bst.root, v)
}

/* Delete a node if there is any */
func (bst *BSTree) Delete(v int) bool {
	// Check if the value is in tree
	if !bst.Exists(v) {
		return false
	}
	// Deletion
	if bst.root == nil {
		return false
	}
	if bst.root.value == v {
		tempRoot := &node{nil, nil, 0}
		tempRoot.left = bst.root
		r := del(bst.root, tempRoot, v)
		bst.root = tempRoot.left
		return r
	}
	return del(bst.root.left, bst.root, v) || del(bst.root.right, bst.root, v)
}

/* Return a list with all the values in ascend order */
func (bst *BSTree) IterAscend() []int {
	return inOrderTraverse(bst.root)
}

func (bst *BSTree) Root() *node {
	return bst.root
}

func (bst *BSTree) Size() int {
	return bst.num
}

func (bst *BSTree) Print() {
	rootNode := bst.Root()
	//printInOrder(rootNode)
	printPreOrder(rootNode)
}

/* For fun. Return the longest path in the tree */
func (bst *BSTree) longestPath() int {
	calLength(bst.root)
	return maxPathLength
}

/* Test TODO: need to spend some time to actually write a test file*/
func main() {
	tree := NewBSTree(10)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(13)
	tree.Insert(11)
	tree.Insert(14)
	tree.Insert(12)
	tree.Insert(18)
	tree.Insert(20)
	tree.Print()
	fmt.Printf("The longest path in this tree is of length %d\n", tree.longestPath())
	fmt.Println(tree.Exists(10))
	fmt.Println(tree.Exists(19))
	fmt.Println(tree.IterAscend())
	fmt.Println(tree.Delete(4))
	fmt.Println(tree.Delete(13))
	fmt.Println(tree.Exists(10))
	fmt.Println(tree.Exists(19))
	fmt.Println(tree.IterAscend())
	fmt.Println(tree.Delete(20))
}
