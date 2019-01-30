package goserver

import (
	"errors"
	"strconv"
)

const (
	InvalidUserIdMessage    = "invalid user id - get real dude"
	InvalidDayMessage       = "invalid day given - get real dude"
	InvalidHourMessage      = "invalid hour given - get real dude"
	InvalidStepCountMessage = "invalid step count given - get real dude"
)

type StepData struct {
	// TODO: Convert UserId this to int64 instead?
	UserId    int `datastore:"userId"`
	Day       int `datastore:"day"`
	Hour      int `datastore:"hour"`
	StepCount int `datastore:"stepCount"`
}

// NewStepData constructs a StepData structure after input validation
// Rather than using a heavy validation framework, we just use custom validation
func NewStepData(userIdString, dayString, hourString, stepCountString string) (stepData *StepData, err error) {
	var userId, day, hour, stepCount int
	// TODO: Validate values as well?
	if userId, err = strconv.Atoi(userIdString); err != nil {
		return nil, errors.New(InvalidUserIdMessage)
	}
	if day, err = strconv.Atoi(dayString); err != nil {
		return nil, errors.New(InvalidDayMessage)
	}
	if hour, err = strconv.Atoi(hourString); err != nil {
		return nil, errors.New(InvalidHourMessage)
	}
	if stepCount, err = strconv.Atoi(stepCountString); err != nil {
		return nil, errors.New(InvalidStepCountMessage)
	}
	return &StepData{userId, day, hour, stepCount}, nil
}
