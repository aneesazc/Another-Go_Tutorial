package errors

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		input    string
		expected TimeStruct
		errMsg   string
	}{
		{"14:07:33", TimeStruct{Hour: 14, Minute: 7, Second: 33}, ""},
		{"23:59:59", TimeStruct{Hour: 23, Minute: 59, Second: 59}, ""},
		{"00:00:00", TimeStruct{Hour: 0, Minute: 0, Second: 0}, ""},
		{"12:34:56", TimeStruct{Hour: 12, Minute: 34, Second: 56}, ""},
		{"24:00:00", TimeStruct{}, "hour must be between 0 and 23"},
		{"12:60:00", TimeStruct{}, "minute must be between 0 and 59"},
		{"12:00:60", TimeStruct{}, "second must be between 0 and 59"},
		{"12:00", TimeStruct{}, "time string must have 3 components"},
		{"abc:def:ghi", TimeStruct{}, "invalid hour component"},
	}

	for _, tt := range tests {
		result, err := ParseTime(tt.input)
		if err != nil {
			if err.Error() != tt.errMsg {
				t.Errorf("ParseTime(%q) returned error %q, expected %q", tt.input, err.Error(), tt.errMsg)
			}
		} else {
			if result != tt.expected {
				t.Errorf("ParseTime(%q) = %+v, expected %+v", tt.input, result, tt.expected)
			}
		}
	}
}
