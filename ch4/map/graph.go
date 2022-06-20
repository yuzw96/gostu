package main

import "fmt"

//map的value也可以是一个聚合类型，比如是一个map或者slice
var graph = make(map[string]map[string]bool)

func main() {
	addEdge("test", "bool")
	fmt.Println(graph)
	fmt.Println(hasEdge("test", "bool"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
