package main

import (
	"net/http"
	"os"
	"log"

	lg "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := lg.NewLogfmtLogger(os.Stderr)

	var svc SmrtService

	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Printf("Could not connect to DB: %+v", err)
	}

	InitDb(db)

	store := NewStorage(db) 
	newSearchEngine := NewSearchEngine()
	svc = smrtService{store,newSearchEngine}
	svc = loggingMiddleware{logger, svc}

	addLineHandler := httptransport.NewServer(
		makeAddLineEndpoint(svc),
		decodeAddLineRequest,
		encodeResponse,
	)

	searchPathHandler := httptransport.NewServer(
		makeSearchPathEndpoint(svc),
		decodeSearchPathRequest,
		encodeResponse,
	)

	http.Handle("/addline", addLineHandler)
	http.Handle("/searchpath", searchPathHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}