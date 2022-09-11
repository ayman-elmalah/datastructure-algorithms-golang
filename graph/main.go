package main

import "fmt"

type Node struct {
	value string
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (graph *Graph) AddNode(n *Node) {
	graph.nodes = append(graph.nodes, n)
}

func (graph *Graph) AddEdge(n1, n2 *Node) {
	if graph.edges == nil {
		graph.edges = make(map[Node][]*Node)
	}
	graph.edges[*n1] = append(graph.edges[*n1], n2)
	graph.edges[*n2] = append(graph.edges[*n2], n1)
}

func (graph *Graph) Display() string {
	s := ""
	for i := 0; i < len(graph.nodes); i++ {
		s += graph.nodes[i].value + " -> "
		near := graph.edges[*graph.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].value + " "
		}
		s += "\n"
	}

	return s
}

func main() {
	graph := Graph{}

	nodeA := &Node{value: "A"}
	nodeB := &Node{value: "B"}
	nodeC := &Node{value: "C"}
	nodeD := &Node{value: "D"}

	graph.AddNode(nodeA)
	graph.AddNode(nodeB)
	graph.AddNode(nodeC)
	graph.AddNode(nodeD)

	graph.AddEdge(nodeA, nodeB)
	graph.AddEdge(nodeA, nodeC)
	graph.AddEdge(nodeA, nodeD)
	graph.AddEdge(nodeB, nodeC)

	fmt.Println(graph.Display())
	/*
		Output :
			A -> B C D
			B -> A C
			C -> A B
			D -> A
	*/
}
