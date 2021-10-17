package ch04

type Node struct {
	value int
	children []*Node
}

func NewNode(n int) *Node {
	l := new(Node)
	l.value = n
	return l
}

func (n *Node) AddLeftLeaf(v int) {
	l := NewNode(v)
	if cap(n.children) == 0 {
		n.children = make([]*Node, 1)
	}
	if cap(n.children) != 2 {
		n.children = append([]*Node{l}, n.children...)
	} else {
		n.children[0] = l
	}
}

func (n *Node) AddRightLeaf(v int) {
	l := NewNode(v)
	if cap(n.children) == 0 {
		n.children = make([]*Node, 2)
	}
	if cap(n.children) != 2 {
		n.children = append(n.children, l)
	} else {
		n.children[1] = l
	}
}

func (n *Node) IsBinaryNode() bool {
	return len(n.children) == 2
}

func (n *Node) IsBinaryTree() bool {
	if len(n.children) == 0 {
		return true
	}
	if len(n.children) == 1 {
		return false
	}
	if len(n.children) == 2 {
		return n.children[0].IsBinaryTree() && n.children[1].IsBinaryTree()
	}
	return false
}

func (n *Node) IsBinarySearchTree(alsoCompare int) bool {
	if len(n.children) == 0 {
		return true
	}
	if len(n.children) == 1 {
		return false
	}
	if len(n.children) == 2 {
		if n.children[0].value > n.value || n.value >= n.children[1].value {
			return false
		}
		if alsoCompare != 0 {
			if n.children[0].value > alsoCompare || alsoCompare >= n.children[1].value {
				return false
			}
		}
		return n.children[0].IsBinarySearchTree(alsoCompare) && n.children[1].IsBinarySearchTree(alsoCompare) && n.children[0].IsBinarySearchTree(n.value) && n.children[1].IsBinarySearchTree(n.value) 
	}
	return false
}

func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) NumberOfBranches() int {
	return len(n.children)
}

func (n *Node) LeftBranch() *Node {
	return n.children[0]
}

func (n *Node) RightBranch() *Node {
	return n.children[1]
}
