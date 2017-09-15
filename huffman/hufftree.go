package huffman

type hufftree struct {
	root *huffnode
}

type huffnode struct {
	left *huffnode
	right *huffnode
	data uint8
}


func (tree hufftree) insert(node *huffnode) {
	// TODO: implement insert
}