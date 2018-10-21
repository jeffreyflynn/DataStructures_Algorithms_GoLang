package main

import "fmt"

// Node holds an interface value
type Node struct {
	val interface{}
}

// return node value as string
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}

// Graph is composed of Nodes and Edges (Connections)
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
	// for undirected graph, execute both statements - undirected edges are a two-way street
	// for a directed graph, only run the first statement - directed edges only go one direction
	G.edges[*n1] = append(G.edges[*n1], n2)
	// G.edges[*n2] = append(G.edges[*n2], n1)
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

// BFS traverses a graph starting from the root node
// first it traverses all the nodes directly linked to the current node
// then we process the nodes directly linked to those nodes
func (G *Graph) BFS(cb func(*Node)) {

	// initialize a queue and add the graph's root node to it
	q := []*Node{G.nodes[0]}

	// initialize a map to keep track of the visited nodes
	// a map is an unordered collection of key/value pairs
	visited := make(map[*Node]bool)

	for {
		// loop until the queue is empty
		if len(q) == 0 {
			break
		}

		// set ref variable to the next node in the queue
		node := q[0]
		// remove that node from the queue - aka dequeue
		q = q[1:]
		// set ref variable for the current node's edges (connections)
		near := G.edges[*node]

		// loop through each edge
		for i := 0; i < len(near); i++ {
			j := near[i]
			// if this edge node has not yet been visited...
			if !visited[j] {
				// ...add it to the queue
				q = append(q, j)
				visited[j] = true
			}
		}

		if cb != nil {
			cb(node)
		}
	}
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
