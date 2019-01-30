package goserver

import (
	"net/http"
	"path"
)

type CurrentDayHandler struct {
}

func (currentDayHandler *CurrentDayHandler) ServeHttp(res http.ResponseWriter, req *http.Request) {
    _, userId := path.Split(req.URL.Path)
    queryCurrentDaySteps(res, req, userId)
}
