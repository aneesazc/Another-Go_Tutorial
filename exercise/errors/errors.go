//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package errors

import (
	"errors"
	"strconv"
	"strings"
)

type TimeStruct struct {
	Hour   int
	Minute int
	Second int
}

func ParseTime(time string) (TimeStruct, error) {
	timeComponents := strings.Split(time, ":")
	if len(timeComponents) != 3 {
		return TimeStruct{}, errors.New("time string must have 3 components")
	}

	hour, err := strconv.Atoi(timeComponents[0])
	if err != nil {
		return TimeStruct{}, errors.New("invalid hour component")
	}

	minute, err := strconv.Atoi(timeComponents[1])
	if err != nil {
		return TimeStruct{}, errors.New("invalid minute component")
	}

	second, err := strconv.Atoi(timeComponents[2])
	if err != nil {
		return TimeStruct{}, errors.New("invalid second component")
	}

	if hour < 0 || hour > 23 {
		return TimeStruct{}, errors.New("hour must be between 0 and 23")
	}
	if minute < 0 || minute > 59 {
		return TimeStruct{}, errors.New("minute must be between 0 and 59")
	}
	if second < 0 || second > 59 {
		return TimeStruct{}, errors.New("second must be between 0 and 59")
	}

	return TimeStruct{Hour: hour, Minute: minute, Second: second}, nil
}
