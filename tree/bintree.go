package tree

import "cmp"

type BSTNode[T cmp.Ordered] struct {
	Key    T
	Left   *BSTNode[T]
	Right  *BSTNode[T]
	Parent *BSTNode[T]
}

func (n *BSTNode[T]) IsLeftChild() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Left == n
}

func (n *BSTNode[T]) IsRightChild() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Right == n
}

type BST[T cmp.Ordered] struct {
	Root *BSTNode[T]
}

func NewBinaryTree[T cmp.Ordered](keys []T) *BST[T] {
	t := &BST[T]{}
	for _, k := range keys {
		t.Insert(k)
	}
	return t
}

func (t *BST[T]) Insert(key T) {
	x := &BSTNode[T]{Key: key}
	if t.Root == nil {
		t.Root = x
		return
	}
	insertBST(t.Root, x)
}

func (t *BST[T]) Delete(key T) {
	x := t.Search(t.Root, key)
	p := x.Parent

	if x.Left == nil && x.Right == nil {
		if p == nil {
			t.Root = nil
		} else if x.IsLeftChild() {
			p.Left = nil
		} else {
			p.Right = nil
		}
	} else if x.Left == nil {
		if p == nil {
			t.Root = x.Right
		} else if x.IsLeftChild() {
			p.Left = x.Right
		} else {
			p.Right = x.Right
		}
		x.Right.Parent = p
	} else if x.Right == nil {
		if p == nil {
			t.Root = x.Left
		} else if x.IsLeftChild() {
			p.Left = x.Left
		} else {
			p.Right = x.Left
		}
		x.Left.Parent = p
	} else {
		next := t.Next(x)
		x.Key = next.Key
		pn := next.Parent
		if pn.Right == next {
			pn.Right = next.Right
			if next.Right != nil {
				next.Right.Parent = pn
			}
		} else if pn.Left == next {
			pn.Left = next.Right
			if next.Right != nil {
				next.Right.Parent = pn
			}
		}
	}
}

func (t *BST[T]) Search(currRoot *BSTNode[T], key T) *BSTNode[T] {
	if currRoot == nil || key == currRoot.Key {
		return currRoot
	}
	if key < currRoot.Key {
		return t.Search(currRoot.Left, key)
	} else {
		return t.Search(currRoot.Right, key)
	}
}

func (t *BST[T]) Next(x *BSTNode[T]) *BSTNode[T] {
	if x.Right != nil {
		return t.Minimum(x.Right)
	}
	p := x.Parent
	for p != nil && x == p.Right {
		x = p
		p = x.Parent
	}
	return p
}

func (t *BST[T]) Prev(x *BSTNode[T]) *BSTNode[T] {
	if x.Left != nil {
		return t.Maximum(x.Left)
	}
	p := x.Parent
	for p != nil && x == p.Left {
		x = p
		p = x.Parent
	}
	return p
}

func (t *BST[T]) Minimum(x *BSTNode[T]) *BSTNode[T] {
	for curr := x; curr != nil; curr = curr.Left {
		x = curr
	}
	return x
}

func (t *BST[T]) Maximum(x *BSTNode[T]) *BSTNode[T] {
	for curr := x; curr != nil; curr = curr.Right {
		x = curr
	}
	return x
}

func (t *BST[T]) PreorderTraversal(x *BSTNode[T], fn func(n *BSTNode[T]) bool) {
	if x == nil {
		return
	}
	fn(x)
	t.PreorderTraversal(x.Left, fn)
	t.PreorderTraversal(x.Right, fn)
}

func (t *BST[T]) InorderTraversal(x *BSTNode[T], fn func(n *BSTNode[T]) bool) {
	if x == nil {
		return
	}
	t.InorderTraversal(x.Left, fn)
	fn(x)
	t.InorderTraversal(x.Right, fn)
}

func (t *BST[T]) PostorderTraversal(x *BSTNode[T], fn func(n *BSTNode[T]) bool) {
	if x == nil {
		return
	}
	t.PostorderTraversal(x.Left, fn)
	t.PostorderTraversal(x.Right, fn)
	fn(x)
}

func insertBST[T cmp.Ordered](currRoot, node *BSTNode[T]) {
	if node.Key > currRoot.Key {
		if currRoot.Right == nil {
			currRoot.Right = node
			node.Parent = currRoot
		} else {
			insertBST(currRoot.Right, node)
		}
	}

	if node.Key < currRoot.Key {
		if currRoot.Left == nil {
			currRoot.Left = node
			node.Parent = currRoot
		} else {
			insertBST(currRoot.Left, node)
		}
	}
}
