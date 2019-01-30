// this file comprises of set of functions to interact with datastore
// one feature could be adding a DAO to avoid calling datastore API directly
package goserver

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
)

func store(res http.ResponseWriter, req *http.Request, stepRecord *StepData) {
	ctx := appengine.NewContext(req)
	keyName := strconv.Itoa(stepRecord.UserId) + "#" + strconv.Itoa(stepRecord.Day) + strconv.Itoa(stepRecord.Hour)
	key := datastore.NewKey(ctx, keyKind, keyName, 0, nil)
	if _, err := datastore.Put(ctx, key, stepRecord); err != nil {
		http.Error(res, err.Error(), serverErrorCode)
		return
	}
}

func maxOfTwoInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// query stepCount for most recent day for given userId
// Assumes stepCount is not greater than size of maximum integer
func queryCurrentDaySteps(res http.ResponseWriter, req *http.Request, userId string) int {
	ctx := appengine.NewContext(req)
	query := datastore.NewQuery(keyKind).Filter("Uid =", userId)
	var mostRecentDay int
	// Todo: int or int64? profile.
	stepsOnDay := make(map[int]int)
	for iterator := query.Run(ctx); ; {
        var stepData StepData
        // returned key is not useful
        _, err := iterator.Next(&stepData)
        if err == datastore.Done {
        	break
		}
        if err != nil {
        	http.Error(res, err.Error(), serverErrorCode)
		}
        day := stepData.Day
        mostRecentDay = maxOfTwoInt(mostRecentDay, day)
        // modifies stepsOnDay
        updateValueAggregate(stepsOnDay, day, stepData.StepCount)
	}
	return stepsOnDay[mostRecentDay]
}

// query stepCount for specific day for given userId
func querySpecificDaySteps(res http.ResponseWriter, req *http.Request, userId string, day int) int {
	ctx := appengine.NewContext(req)
	query := datastore.NewQuery(keyKind).Filter("Uid =", userId).Filter("Day =", day)
	totalStepCount := 0
	for iterator := query.Run(ctx); ; {
		var stepData StepData
		// returned key is not useful
        _, err := iterator.Next(&stepData)
        if err == datastore.Done {
        	break
		}
        if err != nil {
        	http.Error(res, err.Error(), serverErrorCode)
		}
        totalStepCount += stepData.StepCount
	}
	return totalStepCount
}

// for a given userId, compute stepCount over half open [startDay, startDay+numOfDays) interval
func queryRangeOfDays(res http.ResponseWriter, req *http.Request, userId string, startDay, numOfDays int) ([]int, int) {
	ctx := appengine.NewContext(req)
	query := datastore.NewQuery(keyKind).Filter("Uid =", userId)
	// Todo: int or int64? profile.
	stepsOnDay := make(map[int]int)
	for iterator := query.Run(ctx); ; {
        var stepData StepData
        // returned key is not useful
        _, err := iterator.Next(&stepData)
        if err == datastore.Done {
        	break
		}
        if err != nil {
        	http.Error(res, err.Error(), serverErrorCode)
		}
        if stepData.Day >= startDay && stepData.Day < startDay + numOfDays {
        	// modifies stepsOnDay
			updateValueAggregate(stepsOnDay, stepData.Day, stepData.StepCount)
		}
	}
	return countStepsInRangeAndTotalSteps(stepsOnDay, startDay, numOfDays)
}

// Faster to compute steps in range of days and total steps over this interval in a
// single function
func countStepsInRangeAndTotalSteps(stepsOnDay map[int]int, startDay, numOfDays int) ([]int, int) {
	stepsInRange := make([]int, numOfDays)
    aggregateStepCount := 0
    for i := range stepsInRange {
    	stepsToAdd := stepsOnDay[i+startDay]
    	stepsInRange[i] = stepsToAdd
    	aggregateStepCount += stepsToAdd
	}
    return stepsInRange, aggregateStepCount
}

// Go idiom for map update, modifies map given as a parameter
func updateValueAggregate(stepsOnDay map[int]int, day, stepCount int) {
		if steps, ok := stepsOnDay[day]; !ok {
			stepsOnDay[day] = stepCount
		} else {
			stepsOnDay[day] += steps
		}
}
