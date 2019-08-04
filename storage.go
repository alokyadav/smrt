package main

import (
	"errors"
)



// NewLineStorage initializes the storage
func NewStorage() *Storage {
	return &Storage{make(map[string]*Station),[]*DistanceRecord{},make(map[string]*LineRecord)}
}


// UserStorage stores all users
type Storage struct {
	StationTable  map[string]*Station
	DistanceTable []*DistanceRecord
	LineTable     map[string]*LineRecord 
}


func (s Storage) StoreLine(line *Line) error {
	if _, ok := s.LineTable[line.ID]; ok {
		return ErrLinePresent
	}
	lineRecord := new(LineRecord)
	lineRecord.ID = line.ID
	lineRecord.Name = line.Name
	for _,station := range line.Stations {
		lineRecord.Stations = append(lineRecord.Stations,station.ID)
	}
	s.StoreStation(line.Stations)
	s.StoreDistance(line.Stations,line.Distances)
	return nil
}

func (s Storage) StoreStation(stations []*Station) {
	for _,station := range stations {
		if _, ok := s.StationTable[station.ID]; !ok {
			s.StationTable[station.ID] = station
		}
	}
}

func (s Storage) StoreDistance(stations []*Station, distances []int) {
	for i := 0; i < len(stations)-1 ; i++ {
		distanceRecord := new(DistanceRecord)
		distanceRecord.Start  = stations[0].ID
		distanceRecord.End  = stations[1].ID
		distanceRecord.Distance = distances[0]
	}
}


var ErrLinePresent = errors.New("line already present")
