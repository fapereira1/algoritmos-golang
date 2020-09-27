package binarytree

// Comparable define a new func
type Comparable func(c1 interface{}, c2 interface{}) bool

// BinaryTree represent the data struct
type BinaryTree struct {
	node    interface{}
	left    *BinaryTree
	right   *BinaryTree
	lessFun Comparable
}
