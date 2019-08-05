package main

import (
	"errors"
)



// NewLineStorage initializes the storage
func NewStorage() *Storage {
	return &Storage{make(map[string]*Station),make([]*DistanceRecord,0),make(map[string]*LineRecord)}
}


// UserStorage stores all users
type Storage struct {
	StationTable  map[string]*Station
	DistanceTable []*DistanceRecord
	LineTable     map[string]*LineRecord
}


func (s *Storage) StoreLine(line *Line) ([]*DistanceRecord,error) {
	if _, ok := s.LineTable[line.ID]; ok {
		return nil,ErrLinePresent
	}
	lineRecord := new(LineRecord)
	lineRecord.ID = line.ID
	lineRecord.Name = line.Name
	s.LineTable[lineRecord.ID] = lineRecord
	s.StoreStation(line.Stations)
	distanceRecords := s.StoreDistance(line.Stations,line.Distances,line.ID)
    s.DistanceTable = append(s.DistanceTable, distanceRecords...)
	return distanceRecords,nil
}


func (s *Storage) StoreStation(stations []*Station) {
	for _,station := range stations {
		if _, ok := s.StationTable[station.ID]; !ok {
			s.StationTable[station.ID] = station
		} else {
		}
	}
}


func (s *Storage) StoreDistance(stations []*Station, distances []int, lineID string) []*DistanceRecord {
	distanceTable := []*DistanceRecord {}
	for i := 0; i < (len(stations)-1) ; i++ {
		distanceRecord := new(DistanceRecord)
		distanceRecord.Start  = stations[i].ID
		distanceRecord.End  = stations[i+1].ID
		distanceRecord.Distance = distances[i]
		distanceRecord.LineID = lineID
		distanceTable = append(distanceTable,distanceRecord)
	}
	return distanceTable

}

func (s *Storage) GetNumberOfStations() int {
	return len(s.StationTable)
}


var ErrLinePresent = errors.New("line already present")
