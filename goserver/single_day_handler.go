package goserver

import (
	"net/http"
	"strconv"
)

type SingleDayHandler struct {
}

func (singleDayHandler *SingleDayHandler) ServeHttp(res http.ResponseWriter, req *http.Request) {
	userId, restOfUrl := SplitPath(req.URL.Path)
	dayToQueryAsString, restOfUrl := SplitPath(restOfUrl)
	dayToQuery, err := strconv.Atoi(dayToQueryAsString)
	if err != nil {
		http.Error(res, err.Error(), badRequestErrorCode)
	} else {
		querySpecificDaySteps(res, req, userId, dayToQuery)
	}
}
