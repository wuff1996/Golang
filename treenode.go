package main

import "fmt"
type tree struct {
	root int
	left,right *tree

}

func (node *tree) print (){
	fmt.Println(node.root)
}

func (node *tree) treeverse () {
	if node ==nil {
		return
	}
	node.left.treeverse()
	node.print()
	node.right.treeverse()

}

func treenode() {
	var node tree
	node.root = 1
	node.left = new(tree)
	node.right = &tree{5,nil,nil}
	node.left.right = &tree{0,nil,nil}
	node.left.left = &tree{77,nil,nil}
	node.print()
	node.left.left.print()
	node.treeverse()
	
}
