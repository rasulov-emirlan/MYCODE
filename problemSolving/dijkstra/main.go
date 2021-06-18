package main

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	Name        string
	Connections []string
	Cost        map[string]int
}

func (g *Graph) addNode(parent string, destenation string, cost int) {
	tempNode := Node{
		Name: destenation,
	}
	g.Nodes[destenation] = &tempNode
	g.Nodes[parent].Connections = append(g.Nodes[parent].Connections, destenation)
	g.Nodes[parent].Cost[parent] = cost
}

func main() {
	var graphic Graph

	graphic.addNode("NewYork", "Washington", 10)
	graphic.addNode("Texas", "Washington", 10)
	graphic.addNode("NewYork", "Washington", 10)
}
