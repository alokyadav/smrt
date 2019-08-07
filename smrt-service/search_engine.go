package main 


import(
	"errors"
)

type SearchEngineInterface interface {
	SearhPath(src,dest,criteria string) ([][]string,error)
}

type SearchEngine struct {
	graph *Graph
	lineGraph *LineGraph
}


func NewSearchEngine() *SearchEngine {
	graph := NewGraph()
	lineGraph := NewLineGraph()
	return &SearchEngine{graph,lineGraph}
}
//ShortestPath -  This method will find shortest between two given stations using modified dijkstra algorithm
//wth custom weight - least time, less number of hops , least switches . For least time weight 
func (s *SearchEngine) SearhPath(src,dest,criteria string) ([][]string,error) {

	paths := [][]string{}
	err := errors.New("")
	switch criteria {
	case "LEAST_TIME": // least time , just return the time taken 
		paths,err =  s.graph.ShortestPathLeastTime(src,dest)
		break;
	case "LEAST_STATION_NUMBER": // least number of of station always return 1
		paths,err = s.graph.ShortestPathLeastStation(src,dest)
		break;
	case "LEAST_SWITCH":
		paths,err = s.lineGraph.ShortestPathLeastSwitch(src,dest)
	}
	return paths, err
		
}