package graph

func IsCyclic(graph map[string][]string) bool {
	for origin := range graph {
		seen := make(map[string]bool)
		var visitAll func(string, []string) bool
		visitAll = func(node string, refs []string) bool {
			for _, n := range refs {
				if n == origin {
					return true
				}
				if !seen[n] {
					seen[n] = true
					return visitAll(n, graph[n])
				}
			}
			return false
		}
		if visitAll(origin, graph[origin]) {
			return true
		}
	}
	return false
}
