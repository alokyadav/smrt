package main

import (
	"errors"
)
// Line 
type Line struct {
	ID 			string 			`json:"id"`
	Name		string   		`json:"name"` 
	Stations   []*Station  		`json:"stations"`
	Times  []int      		    `json:"distances"` 
}

func (l Line) Validate() error {
	if len(l.Stations) < 1 {
       return errors.New("More than two station should require to create a line")
	}
	if len(l.Stations) !=  len(l.Times) + 1 {
		return errors.New("Time interval does not have complete data")
	}

	return nil
}


// Mysql table definition

// Station - This model or table will store the station static information
type Station struct {
	ID		string 	`json:"id" gorm:"primary_key"`
	Name	string  `json:"name"`
}


//TimeRecord - This table or model will store the time take between two station, line on which they fall and hop number
type TimeRecord struct {
	Start       string   
	End         string  
	Time        int
	LineID      string
	HopNumber  int   
}

// This table will store the line Record
type LineRecord struct {
	ID 			string   `gorm:"primary_key"`			
	Name		string
}