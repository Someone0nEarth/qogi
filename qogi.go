package qogi

import (
	"strconv"
	"time"
)

func AtoUi(string string) (uint, error) {
	ui64, err := strconv.ParseUint(string, 10, 0)

	if err != nil {
		return 0, err
	}

	return uint(ui64), err
}

func AtoF(string string) (float64, error) {

	float, err := strconv.ParseFloat(string, 64)

	if err != nil {
		return 0, err
	}

	return float, err
}

func UiToA(uint uint) string {
	return strconv.FormatUint(uint64(uint), 10)
}

func UintOrZero(pointer *uint) uint {
	if pointer == nil {
		return 0
	}

	return *pointer
}

func StringOrEmpty(string *string) string {

	if string == nil {
		return ""
	}

	return *string
}

func UtcOrZeroRFC3339(timeRFC3339 *string) time.Time {

	parsedTimeUtc := UtcOrNilRFC3339(timeRFC3339)

	if parsedTimeUtc == nil {
		return time.Time{}
	}

	return *parsedTimeUtc
}

func UtcOrNilRFC3339(timeRFC3339 *string) *time.Time {

	if timeRFC3339 == nil {
		return nil
	}

	parsedTime, err := time.Parse(time.RFC3339, *timeRFC3339)
	if err != nil {
		return nil
	}

	timeUtc := parsedTime.UTC()

	return &timeUtc
}

func ToTimeRFC3339(timeRFC3339 string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, timeRFC3339)
	return parsedTime
}

func ToPointer[T interface{}](x T) *T {
	return &x
}
