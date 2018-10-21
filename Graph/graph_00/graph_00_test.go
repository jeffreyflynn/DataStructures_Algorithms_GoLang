package main

import (
	"fmt"
	"testing"
)

var g Graph

func fillGraph() {
	nA := Node{"0"} // A
	nB := Node{"1"} // B
	nC := Node{"2"} // C
	nD := Node{"3"} // D
	nE := Node{"4"} // E
	nF := Node{"5"} // F

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nE)
	g.AddEdge(&nA, &nF)
	g.AddEdge(&nB, &nD)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nB)
	g.AddEdge(&nD, &nC)
	g.AddEdge(&nD, &nE)
}

func TestGraphInsert(t *testing.T) {
	fillGraph()
	g.Print()
}

func TestBFS(t *testing.T) {
	fmt.Println("Breadth First Traversal")
	g.BFS(func(n *Node) {
		fmt.Printf("\t%v\n", n)
	})
}

func TestDFS(t *testing.T) {
	fmt.Println("Depth First Traversal")
	g.DFS(func(n *Node) {
		fmt.Printf("\t%v\n", n)
	})
}
