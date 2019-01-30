package goserver

import (
	"net/http"
	"strings"
)

const (
	urlPathSeparator = "/"
	// For POST requests, userId is used for routing and is stripped from the url
	dayPosition = 0
	hourPosition = 1
	stepCountPosition = 2
)

type PostHandler struct {
}

func (postHandler *PostHandler) ServeHttp(head string, res http.ResponseWriter, req *http.Request) {
	userId := head
	dayHourSteps := strings.Split(req.URL.Path, urlPathSeparator)
	day, hour, stepCount := dayHourSteps[dayPosition], dayHourSteps[hourPosition], dayHourSteps[stepCountPosition]
	if stepRecord, err := NewStepData(userId, day, hour, stepCount); err != nil {
		http.Error(res, err.Error(), badRequestErrorCode)
	} else {
		store(res, req, stepRecord)
	}
}
