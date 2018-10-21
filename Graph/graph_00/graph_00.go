package main

import "fmt"

type Node struct {
	val interface{}
}

// return node value as string
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

// add a node to the graph
func (G *Graph) AddNode(val interface{}) {
	G.nodes = append(G.nodes, Node{val})
}

// add edge to the graph
func (G *Graph) AddEdge(n1, n2 interface{}) {
	if G.edges == nil {
		G.edges = make(map[Node][]*Node)
	}
	G.edges[*n1] = append(G.edges[*n1], n2)
	G.edges[*n2] = append(G.edges[*n2], n1)
}

// print graph
func (G *Graph) Print() {
	s := ""

	for i := 0; i < len(G.nodes); i++ {
		s += G.nodes[i].String() + " -> "

		near := G.edges[*G.nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}

		s += "\n"
	}

	fmt.Println(s)
}

func main() {

}
