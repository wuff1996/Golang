package main

func main() {
	var node [5]node
	for i := range node {
		node[i].data = 10 - i
		if i != len(node)-1 {

			node[i].next = &node[i+1]

		}

	}

}

type node struct {
	data int
	next *node
}
