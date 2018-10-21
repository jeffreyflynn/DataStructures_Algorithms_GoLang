package main

import "fmt"

type Node struct {
	val interface{}
}

// return node value as string
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

// AddNode adds a node to the graph
func (G *Graph) AddNode(n *Node) {
	G.nodes = append(G.nodes, n)
}

// AddEdge adds an edge to the graph given two input nodes
func (G *Graph) AddEdge(n1, n2 *Node) {
	if G.edges == nil {
		G.edges = make(map[Node][]*Node)
	}
	G.edges[*n1] = append(G.edges[*n1], n2)
	G.edges[*n2] = append(G.edges[*n2], n1)
}

// Print prints the graph structure to the console
func (G *Graph) Print() {
	s := "\n"

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
	g := Graph{}

	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

	g.Print()
}
