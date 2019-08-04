package main


import(

)

type Graph struct {
	V  int
    AdjList map[string][]*AdjListNode
}

func CreateGraph(V int, distanceRecords []*DistanceRecord) Graph {
	graph := Graph{V, make(map[string][]*AdjListNode)}

	for _,distanceRecord := range distanceRecords {
		graph.AddEdge(distanceRecord.Start, distanceRecord.End, distanceRecord.Distance)
	}
	return graph
}

func (g Graph)AddEdge(src, dest string, weight int) {
	newNode := NewAdjListNode(src, weight)
	g.AdjList[dest] = append(g.AdjList[dest],newNode)
	newNode = NewAdjListNode(dest, weight)
	g.AdjList[src] = append(g.AdjList[src],newNode)
}

struct AdjListNode struct {
	Dest  string
	W     int	
}

func NewAdjListNode(dest string, weight int) * AdjListNode {
	adjListNode := new(AdjListNode)
	adjListNode.Dest = dest
	adjListNode.W = weight
	return adjListNode
}


