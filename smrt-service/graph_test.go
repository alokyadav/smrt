package main

import (
	"testing"
	"log"
	//"reflect"
	
)

func TestShortestPathLeastTime(t *testing.T) {
	line1 := []*TimeRecord{{Start : "1" , End : "2", Time: 2, LineID : "1"}, {Start : "2" , End : "3", Time: 2, LineID : "1"}, {Start : "3" , End : "4", Time: 2, LineID : "1"}}

	line2 := []*TimeRecord{{Start : "5" , End : "3", Time: 2, LineID : "2"}, {Start : "3" , End : "6", Time: 4, LineID : "2"}, {Start : "6" , End : "7", Time: 5, LineID : "2"}}

	line3 := []*TimeRecord{{Start : "8" , End : "2", Time: 5, LineID : "3"}, {Start : "2" , End : "6", Time: 6, LineID : "3"}, {Start : "6" , End : "9", Time: 7, LineID : "3"}}

	timeRecords := append(line1,line2...)
	timeRecords = append(timeRecords,line3...)

	graph := CreateGraph(9,timeRecords)

	paths,_ := graph.ShortestPathLeastTime("1","6")

	log.Printf("%+v",paths)
	if len(paths) != 2 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 2)
	}

	paths,_ = graph.ShortestPathLeastTime("7","1")
	log.Printf("%+v",paths)
	if len(paths) != 2 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 2)
	}

	line4 := []*TimeRecord{{Start : "1" , End : "7", Time: 20, LineID : "4"}, {Start : "7" , End : "10", Time: 1, LineID : "4"}}
	graph.AddEdges(line4)

	paths,_ = graph.ShortestPathLeastTime("1","10")
	log.Printf("%+v",paths)
	if len(paths) != 2 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 2)
	}
}

func TestShortestPathLeastStation(t *testing.T) {
	line1 := []*TimeRecord{{Start : "1" , End : "2", Time: 2, LineID : "1"}, {Start : "2" , End : "3", Time: 2, LineID : "1"}, {Start : "3" , End : "4", Time: 2, LineID : "1"}}

	line2 := []*TimeRecord{{Start : "5" , End : "3", Time: 2, LineID : "2"}, {Start : "3" , End : "6", Time: 4, LineID : "2"}, {Start : "6" , End : "7", Time: 5, LineID : "2"}}

	line3 := []*TimeRecord{{Start : "8" , End : "2", Time: 5, LineID : "3"}, {Start : "2" , End : "6", Time: 1, LineID : "3"}, {Start : "6" , End : "9", Time: 7, LineID : "3"}}

	timeRecords := append(line1,line2...)
	timeRecords = append(timeRecords,line3...)

	graph := CreateGraph(9,timeRecords)

	paths,_ := graph.ShortestPathLeastStation("1","7")
	log.Printf("%+v",paths)
	if len(paths[0]) != 4 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths[0])   , 4)
	}

	line4 := []*TimeRecord{{Start : "1" , End : "7", Time: 20, LineID : "4"}, {Start : "7" , End : "10", Time: 1, LineID : "4"}}

	graph.AddEdges(line4)
	paths,_ = graph.ShortestPathLeastStation("1","10")
	log.Printf("%+v",paths)
	if len(paths[0]) != 3 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths[0])   , 3)
	}

}


func TestShortestPathLeastSwitch1(t *testing.T) {
	graph := NewLineGraph()

	line := &Line{ID: "1", Name :"Green" , Stations: []*Station{&Station{ID: "1", Name : "A"}, &Station{ID: "2", Name : "A"}, &Station{ID: "3", Name : "A"}, &Station{ID: "4", Name : "A"}  }, Times : []int{2,4,2}}
	graph.AddVertex(line)
	line2 := &Line{ID: "2", Name :"Yellow" , Stations: []*Station{&Station{ID: "5", Name : "A"}, &Station{ID: "3", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "7", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line2)


	line3 := &Line{ID: "3", Name :"Blue" , Stations: []*Station{&Station{ID: "8", Name : "A"}, &Station{ID: "2", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "10", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line3)
	paths, _ := graph.ShortestPathLeastSwitch("1","7")
	log.Printf("paths %+v", paths)
	if len(paths) != 1 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 1)
	}

}

func TestShortestPathLeastSwitch2(t *testing.T) {
	graph := NewLineGraph()

	line := &Line{ID: "1", Name :"Green" , Stations: []*Station{&Station{ID: "1", Name : "A"}, &Station{ID: "2", Name : "A"}, &Station{ID: "3", Name : "A"}, &Station{ID: "4", Name : "A"}  }, Times : []int{2,4,2}}
	graph.AddVertex(line)
	line2 := &Line{ID: "2", Name :"Yellow" , Stations: []*Station{&Station{ID: "5", Name : "A"}, &Station{ID: "3", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "7", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line2)


	line3 := &Line{ID: "3", Name :"Blue" , Stations: []*Station{&Station{ID: "8", Name : "A"}, &Station{ID: "2", Name : "A"}, &Station{ID: "9", Name : "A"}, &Station{ID: "10", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line3)
	line4 := &Line{ID: "4", Name :"Red" , Stations: []*Station{&Station{ID: "11", Name : "A"}, &Station{ID: "9", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "12", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line4)
	paths,_ := graph.ShortestPathLeastSwitch("2","6")
	log.Printf("paths %+v", paths)
	if len(paths) != 2 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 2)
	}


}

func TestShortestPathLeastSwitch3(t *testing.T) {
	graph := NewLineGraph()

	line := &Line{ID: "1", Name :"Green" , Stations: []*Station{&Station{ID: "1", Name : "A"}, &Station{ID: "4", Name : "A"} }, Times : []int{2}}
	graph.AddVertex(line)
	line2 := &Line{ID: "2", Name :"Yellow" , Stations: []*Station{&Station{ID: "1", Name : "A"}, &Station{ID: "2", Name : "A"}}, Times : []int{2}}
	graph.AddVertex(line2)


	line3 := &Line{ID: "3", Name :"Blue" , Stations: []*Station{&Station{ID: "8", Name : "A"}, &Station{ID: "3", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "7", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line3)
	line4 := &Line{ID: "4", Name :"Red" , Stations: []*Station{&Station{ID: "2", Name : "A"}, &Station{ID: "6", Name : "A"}, &Station{ID: "9", Name : "A"}, &Station{ID: "12", Name : "A"} }, Times : []int{2,4,2}}
	graph.AddVertex(line4)

	line5 := &Line{ID: "5", Name :"Ora" , Stations: []*Station{&Station{ID: "4", Name : "A"}, &Station{ID: "3", Name : "A"} }, Times : []int{2}}
	graph.AddVertex(line5)

	paths,_ := graph.ShortestPathLeastSwitch("1","7")
	log.Printf("paths %+v", paths)

	if len(paths) != 2 {
		t.Errorf("Output number of data is not correct  got: %d, want: %d.", len(paths)   , 2)
	}


}

