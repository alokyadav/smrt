package main

//This file implement graph data structure using modified adjacency list
import(
	"container/heap"
	"strings"
	"errors"
	"sync"
)

// Const definition
const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 

type Graph struct {
	V  int  // Number of vertex
	AdjList map[string][]*AdjListNode // adjacency list
	lock  sync.RWMutex
}

//AdjListNode - represent adjacent vertex of given vertex
type AdjListNode struct {
	Dest   string  // Opposite vertex
	W      int     // Weight 
	LineID string  // the line on which both vertex lies	
}


//CreateGraph - with given edges it will create the adjacency graph
func CreateGraph(V int, timeRecords []*TimeRecord) *Graph {
	graph := Graph{V, make(map[string][]*AdjListNode),sync.RWMutex{}}

	for _,timeRecord := range timeRecords {
		graph.AddEdge(timeRecord.Start, timeRecord.End, timeRecord.Time,timeRecord.LineID)
	}
	return &graph
}

//NewGraph - will create an empty graph
func NewGraph() *Graph {
	return &Graph{0,make(map[string][]*AdjListNode),sync.RWMutex{}}
}

//AddEdges - This method will add multiple edges to graph
func (g *Graph) AddEdges(timeRecords []*TimeRecord) {
	for _,timeRecord := range timeRecords {
		g.AddEdge(timeRecord.Start, timeRecord.End, timeRecord.Time, timeRecord.LineID)
	}
}

//AddEdge - Add two edges for one pair
func (g *Graph)AddEdge(src, dest string, weight int, lineID string) {
	g.lock.Lock()
	newNode := NewAdjListNode(src, weight, lineID)
	g.AdjList[dest] = append(g.AdjList[dest],newNode)
	newNode = NewAdjListNode(dest, weight, lineID)
	g.AdjList[src] = append(g.AdjList[src],newNode)
	g.lock.Unlock()
}

func (g *Graph) SetNumberOfVertex(v int) {
	g.V = v
}


//NewAdjListNode - create a new AdjListNode
func NewAdjListNode(dest string, weight int, lineID string) * AdjListNode {
	adjListNode := new(AdjListNode)
	adjListNode.Dest = dest
	adjListNode.W = weight
	adjListNode.LineID = lineID
	return adjListNode
}



//ShortestPathLeastTime - will return shortest path between two stations having least travel time
func (g *Graph) ShortestPathLeastTime(src,dest string) ([][]string,error) {
		// Store the distance of all vertices from source vertex
	distMap := make(map[string]int)
	//Min heap to get minimum weight path vertex for given vertex
	minHeap := &MinHeap{{v:src,dist:0,parents: make(map[string]string)}}
	heap.Init(minHeap)
	distMap[src] = 0 // init src min distance to 0

	isVisited := make(map[string]bool)
	//Initialize distances of all the vertices from src to infinity except src node
	for k,_ := range g.AdjList {
		if strings.Compare(src,k) != 0 {
			distMap[k] = MaxInt
		}
		isVisited[k] = false
	}
	// iterate till minHeap is empty or all connected vertices from src get iterated
	parents := make(map[string]map[string]bool) 
	for minHeap.Len() > 0 {
		//get minimum weight vertex from min Heap
		top := heap.Pop(minHeap).(Node)
		isVisited[top.v] = true
		// iterate over the vertex adjacent connected vertices
		for _,node := range g.AdjList[top.v] {
			if !isVisited[node.Dest]  {
				// compare the distance of adjacent vertex to distance of parent vertex + weight of edge joining two
				if (distMap[node.Dest] >= (distMap[top.v] + node.W))  {
					if (distMap[node.Dest] > (distMap[top.v] + node.W)) {
						delete(parents,node.Dest)
					}
					// update minimum distance for adjacent vertex
					distMap[node.Dest] = distMap[top.v] + node.W
					// Add all the path which have same number switches to the given node
					if _,ok := parents[node.Dest]; !ok {
						parents[node.Dest] = make(map[string]bool)
					}
					parents[node.Dest][top.v] = true
					//create a new node with min distance and put it in min heap
					node := Node {v:node.Dest,dist:distMap[node.Dest]}
					heap.Push(minHeap,node)
				}
			}
		}
		
	}
	// if there is no path return nil and error message
	if _,ok  := parents[dest]; !ok {
		return nil,NoPathPresent
	} 
	paths := CreatePaths(src,dest,parents)
	return paths,nil
}

//ShortestPathLeastStation - will return shortest path between  two stations having less station hops
func (g *Graph) ShortestPathLeastStation(src,dest string) ([][]string,error)  {
	// Store the distance of all vertices from source vertex
	distMap := make(map[string]int)
	//Min heap to get minimum weight path vertex for given vertex
	minHeap := &MinHeap{{v:src,dist:0,parents: make(map[string]string)}}
	heap.Init(minHeap)
	distMap[src] = 0 // init src min distance to 0

	isVisited := make(map[string]bool)
	//Initialize distances of all the vertices from src to infinity except src node
	for k,_ := range g.AdjList {
		if strings.Compare(src,k) != 0 {
			distMap[k] = MaxInt
		}
		isVisited[k] = false
	}
	parents := make(map[string]map[string]bool) 
	// iterate till minHeap is empty or all connected vertices from src get iterated
	for minHeap.Len() > 0 {
		//get minimum weight vertex from min Heap
		top := heap.Pop(minHeap).(Node)
		isVisited[top.v] = true
		
		// iterate over the vertex adjacent connected vertices
		for _,node := range g.AdjList[top.v] {
			if !isVisited[node.Dest]  {
				// compare the distance of adjacent vertex to distance of parent vertex + weight of edge joining two
				if (distMap[node.Dest] >= (distMap[top.v] + 1))  {
					if (distMap[node.Dest] > (distMap[top.v] + 1)) {
						delete(parents,node.Dest)
					}
					// update minimum distance for adjacent vertex
					distMap[node.Dest] = distMap[top.v] + 1
					//update parent
					if _,ok := parents[node.Dest]; !ok {
						parents[node.Dest] = make(map[string]bool)
					}
					parents[node.Dest][top.v] = true

					//parent[node.Dest] = top.v
					//create a new node with min distance and put it in min heap
					node := Node {v:node.Dest,dist:distMap[node.Dest]}
					heap.Push(minHeap,node)
				}
			}
		}
	}
	//if there is no path return nil and error message
	if _,ok  := parents[dest]; !ok {
		return nil,NoPathPresent
	} 
	paths := CreatePaths(src,dest,parents)
	return paths,nil
}


//LineGraph graph data structure will be used to implement minimum switch search 
// Vertex is each line item and there is edge between two vertex if they are crossing each other with
// weight = 1

type LineGraph struct {
	V  int  	// Number of vertex
	Vertices        map[string]*LineVertex // List of vertices in the graph
	LineAdjList 	map[string][]*LineAdjListNode // adjacency list
	lock  sync.RWMutex
}

//LineAdjListNode - represent adjacent vertex of given vertex
type LineAdjListNode struct {
	Dest   string  // Opposite vertex
	W       int     // Weight 
	ConnectingStation string // this will store the connecting station between two lines
}


//LineVertex represent one line as vertex
type LineVertex struct {
	LineID string    //
	StationSet []string  // Array of all the station id in given line
}

func NewLineGraph() *LineGraph {
	return &LineGraph{0,make(map[string]*LineVertex),make(map[string][]*LineAdjListNode ),sync.RWMutex{}}
}


func (lg * LineGraph) AddVertex(l *Line) {
	vertex := new(LineVertex)
	vertex.StationSet = []string{}
	for _,station := range l.Stations {
		vertex.StationSet = append(vertex.StationSet,station.ID)
	}
	lg.Vertices[l.ID] = vertex
	lg.V++
	lg.ConnectLine(l.ID)
}

func (lg *LineGraph) ConnectLine(lineID string) {
	stationSet :=  lg.Vertices[lineID].StationSet
	for k,vertex := range lg.Vertices {
        if strings.Compare(k,lineID) != 0 {
			 inter,isOvr := checkOverlap(stationSet,vertex.StationSet)
             if isOvr {
				lg.AddEdge(lineID,k,inter)
			 }
		}
	}
}


//NewAdjListNode - create a new AdjListNode
func NewLineAdjListNode(lineID, interSection string) *LineAdjListNode {
	adjListNode := new(LineAdjListNode)
	adjListNode.Dest = lineID
	adjListNode.ConnectingStation = interSection
	adjListNode.W = 1
	return adjListNode
}

func (lg *LineGraph) AddEdge(srcLine, destLine, interSection string) {
	srcAdjListNode := NewLineAdjListNode(destLine,interSection)
	destAdjListNode := NewLineAdjListNode(srcLine,interSection)
	lg.LineAdjList[srcLine] = append(lg.LineAdjList[srcLine] ,srcAdjListNode)
	lg.LineAdjList[destLine] = append(lg.LineAdjList[destLine],destAdjListNode)
}


//ShortestPathLeastSwitch - will return shortest path between two stations having less switch, 
func (g *LineGraph) ShortestPathLeastSwitch(src,dest string) ([][]string,error) {
	srcs := []string{}
	dests := []string{}

	//Get all starting line vertex and destination line vertex
	// all line vertex which contain src station are starting point and similarly all line vertex containing dest are destination
	// vertex in search
	for line,vertex := range g.Vertices {
		if IsPresentInLine(src,vertex.StationSet) {
			srcs = append(srcs,line)
		}
		if IsPresentInLine(dest,vertex.StationSet) {
			dests = append(dests,line)
		}
	}
	paths := [][]string{}

	//Check if src and dest lie on same line
	lineID,ok := checkOverlap(srcs,dests)
	
	if ok {
		onP := g.getPathFromLines(src,dest,[]string{lineID})
		paths := append(paths,onP)
		return paths,nil
	}

	//get minimum switch path for all start vertex to end vertex
	
	for _,lineID := range srcs {
		parents := g.ShortestPathLeastSwitchHelper(lineID)
		
		for _,destLine := range dests {
			if _,ok := parents[destLine] ; ok {
				paths = append(paths, CreatePaths(lineID,destLine, parents)...)
			}
			
		}
	}
	if len(paths) == 0 {
		return nil,NoPathPresent
	}
	//get shortest path
	sortestPaths := [][]string{}
	lenMax := MaxInt
	for _, path := range paths {
		if len(path) <= lenMax {
			if len(path) < lenMax {
				sortestPaths = [][]string{}
			} 
			sortestPaths = append(sortestPaths,path)
			lenMax = len(path)
		}
		
	}
	paths = [][]string{}

	// get station path from line paths
    for _,shortestPath := range sortestPaths {
		onP := g.getPathFromLines(src,dest,shortestPath)
		paths = append(paths,onP)
	}
	return paths,nil

}


//This method will create path from line vertex
func (g *LineGraph)getPathFromLines(src, dest string, linePath []string) []string {

	paths := []string{}
	i := 0
	start := src
	end := ""

	for i < len(linePath)-1 {
		adj := g.LineAdjList[linePath[i]]
		for _,node := range adj {
           if strings.Compare(node.Dest,linePath[i+1]) == 0{
			   end = node.ConnectingStation
			   break;
		   }
		}
		vertex := g.Vertices[linePath[i]]
		subPath := getSliceFromArrayUsingValue(start,end,vertex.StationSet)
		if len(paths) != 0 {
			paths = paths[:len(paths)-1]
		}

		paths = append(paths,subPath...)
		start = end
		i++
	}
	vertex := g.Vertices[linePath[i]]
	end = dest
	subPath := getSliceFromArrayUsingValue(start,end,vertex.StationSet)
	if len(paths) != 0 {
		paths = paths[:len(paths)-1]
	}
	paths = append(paths,subPath...)

	return paths
}

func (g *LineGraph) ShortestPathLeastSwitchHelper(src string) map[string]map[string]bool {
	// Store the distance of all vertices from source vertex
	distMap := make(map[string]int)
    parents := make(map[string]map[string]bool) 
	//Min heap to get minimum weight path vertex for given vertex
	minHeap := &MinHeap{{v:src,dist:0,parents: make(map[string]string)}}
	heap.Init(minHeap)
	distMap[src] = 0 // init src min distance to 0

	isVisited := make(map[string]bool)
	//Initialize distances of all the vertices from src to infinity except src node
	for k,_ := range g.LineAdjList {
		if strings.Compare(src,k) != 0 {
			distMap[k] = MaxInt
		}
		isVisited[k] = false
	}

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Node)
		isVisited[top.v] = true
		
		// iterate over the vertex adjacent connected vertices
		for _,node := range g.LineAdjList[top.v] {
			if !isVisited[node.Dest]  {
				// compare the distance of adjacent vertex to distance of parent vertex + weight of edge joining two
				if (distMap[node.Dest] >= (distMap[top.v] + 1))  {
					if (distMap[node.Dest] > (distMap[top.v] + 1)) {
						delete(parents,node.Dest)
					}
					// update minimum distance for adjacent vertex
					distMap[node.Dest] = distMap[top.v] + 1
					//update parent
					if _,ok := parents[node.Dest]; !ok {
						parents[node.Dest] = make(map[string]bool)
					}
					parents[node.Dest][top.v] = true
					//create a new node with min distance and put it in min heap
					node := Node {v:node.Dest,dist:distMap[node.Dest]}
					heap.Push(minHeap,node)
				}
			}
		}
	}

	return parents
}


func getSliceFromArrayUsingValue(start,end string, stationSet []string) []string {
	j := 0
	k := 0
	for j < len(stationSet) {
		if strings.Compare(start,stationSet[j]) == 0 {
			k = j
		}
		if strings.Compare(end,stationSet[j]) == 0 {
			break
		}
		j++
	}
	return stationSet[k:j+1]
}


func IsPresentInLine(src string, stationSet []string) bool {
    for _,k := range stationSet {
		if strings.Compare(k,src) == 0 {
			return true
		}
	}
	return false
}

func CreatePaths(src,dest string, parentMap map[string]map[string]bool) [][]string {
	paths := [][]string{}
	paths = append(paths, []string{dest})
	for {
		temp := [][]string{}
		incFlag := false
		for _,path := range paths {
			last := path[len(path)-1]
			if strings.Compare(src,last) != 0 {
				for k,_ := range parentMap[last] {
					tempPath := make([]string, len(path))
					copy(tempPath,path)
					temp = append(temp, append(tempPath,k))
					incFlag = true
					
				}
			} else {
				temp = append(temp, path)
			}
		}	
		if (incFlag == true) {
			paths = temp
		} else {
			break
		}
	}

	reversePaths := [][]string{}
	for _,path := range paths {
		reversePaths = append(reversePaths,reverseSlice(path))
	}
	return reversePaths
}

//reverseSlice - helpher method to reverse a slice
func reverseSlice(path []string) []string {
	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}
	return path
}

func checkOverlap(setA , setB []string) (string,bool) {
	for _,k := range setA {
       for _,l := range setB {
		   if strings.Compare(k,l) == 0 {
			   return k,true
		   }
	   }
	}
	return "",false
}


var NoPathPresent  = errors.New("No path between src and destination")


