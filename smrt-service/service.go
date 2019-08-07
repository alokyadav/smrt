package main

import (
	"errors"
	"fmt"
	"log"
)

//SmrtService provides operations on strings.
type SmrtService interface {
	AddLine(*Line) error
	SearchPath(src, dest, criteria string) ([][]string,error)
	Init()
}

type smrtService struct{
	store *Storage
	searchEngine *SearchEngine
}

//AddLine - This method will add a new line to database 
// Also it will add stations to search data structure
func (svc smrtService)AddLine(line *Line) error {
	if line == nil {
		return ErrEmpty
	}
	//basic validation of input
	err := line.Validate()
	if err != nil {
		return err
	}
	// Store line will newly added edges
	timeRecs, err := svc.store.StoreLine(line)
	if err != nil {
		return err
	}
	// Add the new edges to the Graph 
	svc.searchEngine.graph.AddEdges(timeRecs)
	svc.searchEngine.graph.SetNumberOfVertex(svc.store.GetNumberOfStations())
	svc.searchEngine.lineGraph.AddVertex(line)

	lines := svc.store.getAllTimeRecords()
    log.Printf("len %d", len(lines))
	return nil
}

func (svc smrtService) SearchPath(src, dest, criteria string) ([][]string,error) {

	// Basic validation to check if both src and dest present in system
	if !svc.store.IsStationPresent(src) {
		errorMsg := fmt.Sprintf("Station with id %s not present in system", src)
		return nil,errors.New(errorMsg)
	}
	if !svc.store.IsStationPresent(dest) {
		errorMsg := fmt.Sprintf("Station with id %s not present in system", dest)
		return nil,errors.New(errorMsg)
	}
	paths,err := svc.searchEngine.SearhPath(src,dest,criteria)
	return paths,err
}

func (svc smrtService) Init() {
	lines := svc.store.getAllLines()
	for _,line := range lines {
		svc.searchEngine.lineGraph.AddVertex(line)
	} 
	timeRecords := svc.store.getAllTimeRecords()
	svc.searchEngine.graph.AddEdges(timeRecords)

}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("Empty string")