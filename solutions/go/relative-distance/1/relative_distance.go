package relativedistance

func DegreeOfSeparation(familyTree map[string][]string, personA, personB string) (int, bool) {
	g := buildGraph(familyTree)
	return bfs(g, personA, personB)
}

func buildGraph(familyTree map[string][]string) map[string][]string {
	graph := make(map[string][]string)
	addEdge := func(a, b string) {
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	for parent, children := range familyTree {
		for _, child := range children {
			addEdge(parent, child)
		}
		for i, childA := range children {
			for _, childB := range children[i+1:] {
				addEdge(childA, childB)
			}
		}
	}
	return graph
}

func bfs(graph map[string][]string, start, target string) (int, bool) {
	if start == target {
		return 0, true
	}
	visited := map[string]bool{start: true}
	queue := []string{start}
	for degree := 1; len(queue) > 0; degree++ {
		var next []string
		for _, person := range queue {
			for _, neighbor := range graph[person] {
				if visited[neighbor] {
					continue
				}
				if neighbor == target {
					return degree, true
				}
				visited[neighbor] = true
				next = append(next, neighbor)
			}
		}
		queue = next
	}
	return 0, false
}
