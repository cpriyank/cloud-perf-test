package goserver

import (
	"net/http"
	"strconv"
)

type App struct {
	CurrentDayHandler *CurrentDayHandler
	RangeOfDaysHandler *RangeOfDaysHandler
	SingleDayHandler *SingleDayHandler
	PostHandler *PostHandler
}

// Trying not to make a router...
func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = SplitPath(req.URL.Path)
	if head == currentDayRequestPrefix {
		app.CurrentDayHandler.ServeHttp(res, req)
		return
	} else if head == singleDayRequestPrefix {
		app.SingleDayHandler.ServeHttp(res, req)
		return
	} else if head == rangeOfDaysRequestPrefix {
		app.RangeOfDaysHandler.ServeHttp(res, req)
		return
	} else if _, err := strconv.Atoi(head); err != nil {
		// head is then a userId
		app.PostHandler.ServeHttp(head, res, req)
		return
	}
	http.Error(res, NotFoundString, http.StatusNotFound)
}

