package huffman

type HuffTree struct {
	root *HuffNode
}

type HuffNode struct {
	left *HuffNode
	right *HuffNode
	num uint8 // stores the occurrences of char or all children
	char uint8 // stores the ASCII-E character code for leaf node
}

func (tree *HuffTree) insert(node *HuffNode) {
	// TODO: implement insertion of nodes
}

func (node *HuffNode) next(dir bool) (ret *HuffNode) {
	// TODO: implement traversal of tree
	return
}