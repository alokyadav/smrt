package main


import(
	"container/heap"
	"strings"
	//"log"
)

type Graph struct {
	V  int
    AdjList map[string][]*AdjListNode
}

type AdjListNode struct {
	Dest   string
	W      int
	LineID string	
}

const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 

func CreateGraph(V int, distanceRecords []*DistanceRecord) *Graph {
	graph := Graph{V, make(map[string][]*AdjListNode)}

	for _,distanceRecord := range distanceRecords {
		graph.AddEdge(distanceRecord.Start, distanceRecord.End, distanceRecord.Distance,"")
	}
	return &graph
}

func NewGraph() *Graph {
	return &Graph{0,make(map[string][]*AdjListNode)}
}

func (g *Graph) AddEdges(distanceRecords []*DistanceRecord) {
	for _,distanceRecord := range distanceRecords {
		g.AddEdge(distanceRecord.Start, distanceRecord.End, distanceRecord.Distance, distanceRecord.LineID)
	}
}


func (g *Graph)AddEdge(src, dest string, weight int, lineID string) {
	newNode := NewAdjListNode(src, weight, lineID)
	g.AdjList[dest] = append(g.AdjList[dest],newNode)
	newNode = NewAdjListNode(dest, weight, lineID)
	g.AdjList[src] = append(g.AdjList[src],newNode)
}


func NewAdjListNode(dest string, weight int, lineID string) * AdjListNode {
	adjListNode := new(AdjListNode)
	adjListNode.Dest = dest
	adjListNode.W = weight
	adjListNode.LineID = lineID
	return adjListNode
}



func (g *Graph) ShortestPath(src, dest, criteria string) ([]string,error) {
	distMap := make(map[string]int)
	parent  := make(map[string]string)
	nodeLine  := make(map[string]string)
	minHeap := &MinHeap{{v:src,dist:0}}
	heap.Init(minHeap)
	distMap[src] = 0
	for k,_ := range g.AdjList {
         if strings.Compare(src,k) != 0 {
			distMap[k] = MaxInt
		 }
	}

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Node)
		for _,node := range g.AdjList[top.v] {		
			weight := g.getWeight(nodeLine,top,node,criteria)
			if distMap[node.Dest] > (distMap[top.v] + weight) {
				distMap[node.Dest] = distMap[top.v] + weight
				parent[node.Dest] = top.v
				nodeLine[node.Dest] = node.LineID
				node := Node {v:node.Dest,dist:distMap[node.Dest]}
				heap.Push(minHeap,node)
			}
		}
	}
	path := []string{}
	i := 0
    path = append(path,dest) 
	for strings.Compare(parent[dest],src) != 0 && (i < g.V) {
		path = append(path,parent[dest])
		dest = parent[dest]
	}

	if i >= g.V {
		return nil,nil
	} 
    path = append(path,src)
	return path,nil

}

func (g *Graph)getWeight(nodeLine map[string]string, parent Node, node *AdjListNode, criteria string) int {
	switch criteria {
	case "LEAST_TIME":
		return node.W
	case "LEAST_SWITCH":
		if _,ok := nodeLine[parent.v]; !ok {
			return 0
		} else {
			if strings.Compare(nodeLine[parent.v],nodeLine[node.Dest]) != 0 {
				return 1
			} else {
				return 0
			}
		}
	case "LEAST_STATION_NUMBER":
		return 1	

	}
    return MaxInt
}

