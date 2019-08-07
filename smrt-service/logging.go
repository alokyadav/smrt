package main


//This file will implement the middleware logging for the service
import (
	"time"
	"encoding/json"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   SmrtService
}

//Addline - wrap and implement logging for AddLine service
func (mw loggingMiddleware) AddLine(line *Line) (err error) {
	b, _:= json.Marshal(line)
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "addLine",
			"input", b,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.AddLine(line)
	return
}

//SearchPath - wrap and implement logging for SearchPath
func (mw loggingMiddleware) SearchPath(src, dest, criteria string) (paths [][]string, err error) {
	defer func(begin time.Time) {
		b, _:= json.Marshal(paths)
		_ = mw.logger.Log(
			"method", "searchPath",
			"src", src,
			"dest", dest,
			"criteria",criteria,
			"path", b,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	paths,err = mw.next.SearchPath(src,dest,criteria)
	return
}