package main

import (
	"errors"
	"log"
)

// SmrtService provides operations on strings.
type SmrtService interface {
	AddLine(*Line) error
	SearchPath(src, dest, criteria string) []string
}

type smrtService struct{
	store *Storage
	graph *Graph
}

func (svc smrtService)AddLine(line *Line) error {
	if line == nil {
		return ErrEmpty
	}
	distRecs, err := svc.store.StoreLine(line)
	if err != nil {
		return err
	}
	svc.graph.AddEdges(distRecs)
	svc.graph.V = svc.store.GetNumberOfStations()
	return nil
}

func (svc smrtService) SearchPath(src, dest, criteria string) []string {

	log.Printf("len %+v", len(svc.store.StationTable))
    for v,k := range svc.graph.AdjList {
		log.Printf("end one %+v", v)
		for _,val := range k {
			log.Printf("end two %+v", val.Dest)
		}
	}
	path,_ := svc.graph.ShortestPath(src,dest,criteria)
	return path
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")