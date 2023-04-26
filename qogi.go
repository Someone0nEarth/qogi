package qogi

import (
	"strconv"
	"time"
)

// AtoUi short for strconv.ParseUint(string, 10, 0)
func AtoUi(str string) (uint, error) {
	ui64, err := strconv.ParseUint(str, 10, 0)

	if err != nil {
		return 0, err
	}

	return uint(ui64), err
}

// AtoF short for strconv.ParseFloat(string, 64)
func AtoF(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, err
	}

	return float, err
}

// UiToA short for strconv.FormatUint(uint64(uint), 10)
func UiToA(uInt uint) string {
	return strconv.FormatUint(uint64(uInt), 10)
}

// UintOrZero pointer unit value or 0 for nil
func UintOrZero(pointer *uint) uint {
	if pointer == nil {
		return 0
	}

	return *pointer
}

// StringOrEmpty pointer string value or "" for nil
func StringOrEmpty(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

// UtcOrZeroRFC3339 Utc time of RFC3339 string or zero time for nil or non-RFC3339 string
func UtcOrZeroRFC3339(timeRFC3339 *string) time.Time {
	parsedTimeUtc := UtcOrNilRFC3339(timeRFC3339)

	if parsedTimeUtc == nil {
		return time.Time{}
	}

	return *parsedTimeUtc
}

// UtcOrNilRFC3339 Utc time of RFC3339 string or nil for nil or non-RFC3339 string
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

// ToTimeRFC3339 short for time.Parse(time.RFC3339, timeRFC3339). Attention: parsing errors will be discarded
func ToTimeRFC3339(timeRFC3339 string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, timeRFC3339)
	return parsedTime
}

// ToPointer returns the pointer for the given argument (even if x is already a pointer)
func ToPointer[T any](x T) *T {
	return &x
}
