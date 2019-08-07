package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeAddLineEndpoint(svc SmrtService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addLineRequest)
		err := svc.AddLine(req.SMRTLine)
		if err != nil {
			return addLineResponse{"", err.Error()}, nil
		}
		return addLineResponse{"Added Successfully", ""}, nil
	}
}

func makeSearchPathEndpoint(svc SmrtService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(searchPathRequest)
		v,err := svc.SearchPath(req.Src,req.Dest,req.Criteria)
		if err != nil {
			return searchPathResponse{nil, err.Error()}, nil
		}
		return searchPathResponse{v,""}, nil
	}
}

func decodeAddLineRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addLineRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSearchPathRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request searchPathRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type addLineRequest struct {
	SMRTLine  *Line  `json:"line"`
}

type addLineResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type searchPathRequest struct {
	Src 		string `json:"source"`
	Dest        string `json:"destination"`
	Criteria    string `json:"criteria"`
}

type searchPathResponse struct {
	Path [][]string `json:"path"`
	Err  string    `json:"err,omitempty"`
}