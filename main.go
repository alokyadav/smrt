package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc SmrtService

	store := NewStorage() 
	graph := NewGraph()
	svc = smrtService{store,graph}
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