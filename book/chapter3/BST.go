package chapter3

type BST struct {
	parent *BST
	left   *BST
	right  *BST
	item   int
}

func searchBst(tree *BST, item int) *BST {
	if tree == nil {
		return &BST{}
	}

	if item > tree.item {
		return searchBst(tree.right, item)
	} else if item < tree.item {
		return searchBst(tree.left, item)
	} else {
		return tree
	}
}

// finds the minimum element in a BST
func minBst(tree *BST) *BST {
	if tree == nil {
		return &BST{}
	}

	if tree.left == nil {
		return tree
	}

	return minBst(tree.left)

}
