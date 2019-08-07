package main

//This file implement in memory storage and act as storage layer that can be removed 
import (
  "errors"
  "strings"
  "sync"
  "fmt"
  "os"
  "log"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

)



// NewLineStorage initializes the storage
func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db, make(map[string]*Station),make([]*TimeRecord,0),make(map[string]*LineRecord),sync.RWMutex{}}
}


// InMemoryStorage stores all users
type Storage struct {
	db            *gorm.DB   
	StationTable  map[string]*Station //In memory table to store Station
	TimeTable     []*TimeRecord       //In memory table to store time records
	LineTable     map[string]*LineRecord //In memory table to store line
	lock  sync.RWMutex
}


func (s *Storage) StoreLine(line *Line) ([]*TimeRecord,error) {

	var lr LineRecord
	//Check if line is already present
	s.db.Where(&LineRecord{ID : line.ID}).First(&lr)
	if strings.Compare(lr.ID,line.ID) == 0 {
		return nil,ErrLinePresent
	}
	lineRecord := new(LineRecord)
	lineRecord.ID = line.ID
	lineRecord.Name = line.Name
	//Store the line Record
	s.db.Create(lineRecord)

	//Store Stations
	s.StoreStation(line.Stations)

	//Store the records
	timeRecords := s.StoreTimeIntervals(line.Stations,line.Times,line.ID)
	//Return added edges
	return timeRecords,nil
}


func (s *Storage) StoreStation(stations []*Station) {
	for _,station := range stations {	
		//Check if station is present
		var st Station
		s.db.Where(&Station{ID : station.ID}).First(&st)
		if strings.Compare(st.ID,station.ID) != 0 {
			s.db.Create(station)
		}
	}
}


//StoreTimeIntervals Store the time interval
func (s *Storage) StoreTimeIntervals(stations []*Station, times []int, lineID string) []*TimeRecord {
	TimeTable := []*TimeRecord {}
	for i := 0; i < (len(stations)-1) ; i++ {
		timeRecord := new(TimeRecord)
		timeRecord.Start  = stations[i].ID
		timeRecord.End  = stations[i+1].ID
		timeRecord.Time = times[i]
		timeRecord.LineID = lineID
		timeRecord.HopNumber = i
		s.db.Create(timeRecord)
		TimeTable = append(TimeTable,timeRecord)
	}
	return TimeTable

}

func (s *Storage) GetNumberOfStations() int {
	var stations []Station
	s.db.Find(&stations)
	return len(stations)
}

//This method check if Station is present
func (s *Storage)IsStationPresent(stationID string) bool {
	var station Station
	s.db.Where(&Station{ID : stationID}).First(&station)
	if strings.Compare(station.ID,stationID) == 0 {
		return true
	}
	return false
}

func (s *Storage) getAllLines() []*Line {
	var lineRecords []LineRecord
	lines := []*Line{}
	s.db.Find(&lineRecords)

	log.Printf("line %+v", lineRecords)
	for _,lineRecord := range lineRecords {
		rows, err := s.db.Raw(`select time_records.start, 
							  time_records.end,
							  time_records.time,
							  time_records.hop_number from time_records 
                              where line_id = ? 
							  order by time_records.hop_number asc`,lineRecord.ID).
							Rows()
		if err != nil {
			log.Printf("line %+v", err)
			return nil
		}
	
		defer rows.Close()
		stations := []*Station{}
		Times := []int{}
		var st string
		var end string
		var t int
		var hop int
		for rows.Next() {
			rows.Scan(&st,&end,&t,&hop)
			stations = append(stations,&Station{ID:st})
			Times = append(Times,t)
		}
		stations = append(stations,&Station{ID:end})
		line := new(Line)
		line.ID = lineRecord.ID
		line.Name = lineRecord.Name
		line.Stations = stations
		line.Times = Times
		lines = append(lines, line)
	}
	return lines
}

func (s *Storage) getAllTimeRecords() []*TimeRecord {
	var timeRecords []*TimeRecord
	s.db.Find(&timeRecords)
    log.Printf("len %+v",len(timeRecords))
	return timeRecords
}

//CreateConnection -  Create connection to mysql Server
func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
    s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,host,port,DBName)
	//s := fmt.Sprintf( "host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", host, user, DBName, port, password)
	return gorm.Open("mysql", s)
}

//Initdb - this methow will drop the table and create new ones
func InitDb(db *gorm.DB) {
	//db.Debug().DropTableIfExists(&Station{}) 
	//Drops table if already exists
	if !db.HasTable(&Station{}) {
		db.Debug().AutoMigrate(&Station{}) 
	}

	//Drops table if already exists
	if !db.HasTable(&TimeRecord{}) {
		db.Debug().AutoMigrate(&TimeRecord{})
	} 

	
	//Drops table if already exists
	if !db.HasTable(&LineRecord{}) {
		db.Debug().AutoMigrate(&LineRecord{}) 
	}
}

var ErrLinePresent = errors.New("line already present")


