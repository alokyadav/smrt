package main




// Line 
type Line struct {
	ID 			string 			`json:"id"`
	Name		string   		`json:"name"` 
	Stations   []*Station  		`json:"stations"`
	Distances  []int      		`json:"distances"` 
}



// Station -
type Station struct {
	ID		string 	`json:"id"`
	Name	string  `json:"name"`
}



type DistanceRecord struct {
	Start       string   
	End         string  
	Distance    int
	LineID      string   
}

type LineRecord struct {
	ID 			string 			
	Name		string
}