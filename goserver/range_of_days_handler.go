package goserver

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	// range URL components are literal "range" and arguments for userId, startDay, numOfDays
	// as a string in that order
	userIdPosition = 1
	startDayPosition = 2
	numOfDaysPosition = 3
	invalidIntStub = -1
)

type RangeOfDaysHandler struct {
}

func (rangeOfDaysHandler *RangeOfDaysHandler) ServeHttp(res http.ResponseWriter, req *http.Request) {
	rangeUrlComponents := strings.Split(req.URL.Path, urlPathSeparator)
	userId := rangeUrlComponents[userIdPosition]
	startDay, numOfDays, err := parseStartAndNumOfDays(rangeUrlComponents[startDayPosition], rangeUrlComponents[numOfDaysPosition])
	if err != nil {
		http.Error(res, err.Error(), badRequestErrorCode)
	} else {
		queryRangeOfDays(res, req, userId, startDay, numOfDays)
	}
}

func parseStartAndNumOfDays (startDayString, numOfDaysString string) (int, int, error) {
    startDay, err := strconv.Atoi(startDayString)
    if err != nil {
    	return invalidIntStub, invalidIntStub, err
	}
    numOfDays, err := strconv.Atoi(numOfDaysString)
    if err != nil {
    	return invalidIntStub, invalidIntStub, err
	}
    return startDay, numOfDays, nil
}
