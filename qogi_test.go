package qogi

import (
	"errors"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test_AtoUiWithMaxUint(t *testing.T) {
	g := NewWithT(t)

	uInt, err := AtoUi("18446744073709551615")

	g.Expect(uInt).Should(BeIdenticalTo(uint(18446744073709551615)))
	g.Expect(err).Should(BeNil())
}

func Test_AtoUi(t *testing.T) {
	g := NewWithT(t)

	uInt, err := AtoUi("7")

	g.Expect(uInt).Should(BeIdenticalTo(uint(7)))
	g.Expect(err).Should(BeNil())
}

func Test_AtoUiOutOfRange(t *testing.T) {
	g := NewWithT(t)

	_, err := AtoUi("18446744073709551616") //max uint +1

	g.Expect(errors.Unwrap(err)).Error().Should(MatchError(errors.New("value out of range")))
}

func Test_AtoUiWithNegativeValue(t *testing.T) {
	g := NewWithT(t)

	_, err := AtoUi("-1")

	g.Expect(errors.Unwrap(err)).Error().Should(MatchError(errors.New("invalid syntax")))
}

func Test_StringOrEmpty(t *testing.T) {
	g := NewWithT(t)

	stringPointer := ToPointer("Hello")
	returnedString := StringOrEmpty(stringPointer)
	g.Expect(returnedString).Should(BeIdenticalTo("Hello"))

	stringPointer = nil
	returnedString = StringOrEmpty(stringPointer)
	g.Expect(returnedString).Should(BeEmpty())
}

func Test_UintOrZero(t *testing.T) {
	g := NewWithT(t)

	uintPointer := ToPointer(uint(123))
	returnedUint := UintOrZero(uintPointer)
	g.Expect(returnedUint).Should(BeIdenticalTo(uint(123)))

	uintPointer = nil
	returnedUint = UintOrZero(uintPointer)
	g.Expect(returnedUint).Should(BeZero())
}

func Test_UtcOrZeroRFC3339ToZeroTime(t *testing.T) {
	g := NewWithT(t)

	zero := UtcOrZeroRFC3339(nil)
	g.Expect(zero.IsZero()).Should(BeTrue())

	invalidRFC3339String := "2016-09-28TTTT18:44:42.208Z"
	zero = UtcOrZeroRFC3339(ToPointer(invalidRFC3339String))
	g.Expect(zero.IsZero()).Should(BeTrue())
}

func Test_UtcOrZeroRFC3339ToUtcTime(t *testing.T) {
	g := NewWithT(t)

	timeString := "2016-09-28T18:44:42.208Z"
	timeInUtc := UtcOrZeroRFC3339(ToPointer(timeString))
	expectedTime, _ := time.Parse(time.RFC3339, "2016-09-28T18:44:42.208Z")
	g.Expect(timeInUtc).Should(BeIdenticalTo(expectedTime.UTC()))

	timeString = "2016-09-28T18:44:42+07:00"
	timeInUtc = UtcOrZeroRFC3339(ToPointer(timeString))
	expectedTime, _ = time.Parse(time.RFC3339, "2016-09-28T11:44:42Z")
	g.Expect(timeInUtc).Should(BeIdenticalTo(expectedTime.UTC()))
}

func Test_TimeUtcOrNilToNil(t *testing.T) {
	g := NewWithT(t)

	nilTime := UtcOrNilRFC3339(nil)
	g.Expect(nilTime).Should(BeNil())

	invalidRFC3339String := "2016-09-28TTTT18:44:42.208Z"
	nilTime = UtcOrNilRFC3339(ToPointer(invalidRFC3339String))
	g.Expect(nilTime).Should(BeNil())
}

func Test_TimeUtcOrNilToUtcTime(t *testing.T) {
	g := NewWithT(t)

	timeString := "2016-09-28T18:44:42.208Z"
	timeInUtc := UtcOrNilRFC3339(ToPointer(timeString))
	expectedTime, _ := time.Parse(time.RFC3339, "2016-09-28T18:44:42.208Z")
	g.Expect(*timeInUtc).Should(BeIdenticalTo(expectedTime.UTC()))

	timeString = "2016-09-28T18:44:42+07:00"
	timeInUtc = UtcOrNilRFC3339(ToPointer(timeString))
	expectedTime, _ = time.Parse(time.RFC3339, "2016-09-28T11:44:42Z")
	g.Expect(*timeInUtc).Should(BeIdenticalTo(expectedTime.UTC()))
}

func Test_ToTimeRFC3339(t *testing.T) {
	g := NewWithT(t)

	timeString := "2016-09-28T18:44:42.208Z"
	toTime := ToTimeRFC3339(timeString)
	expectedTime, _ := time.Parse(time.RFC3339, "2016-09-28T18:44:42.208Z")
	g.Expect(toTime).Should(BeIdenticalTo(expectedTime))

	timeString = "2016-09-28T18:44:42+07:00"
	toTime = ToTimeRFC3339(timeString)
	expectedTime, _ = time.Parse(time.RFC3339, "2016-09-28T18:44:42+07:00")
	g.Expect(toTime).Should(BeEquivalentTo(expectedTime))

	timeString = "Non RFC3339 string"
	toTime = ToTimeRFC3339(timeString)
	g.Expect(toTime.IsZero()).Should(BeTrue())
}

func Test_AtoF(t *testing.T) {
	g := NewWithT(t)

	float, err := AtoF("12345.54321")
	g.Expect(err).Should(BeNil())
	g.Expect(float).Should(BeIdenticalTo(12345.54321))

	float, err = AtoF("1.797693134862315708145274237317043567981e+309")
	g.Expect(err).Error()
}

func Test_UiToA(t *testing.T) {
	g := NewWithT(t)

	returnedString := UiToA(123)

	g.Expect(returnedString).Should(BeIdenticalTo("123"))
}

func Test_ToPointerWithUint(t *testing.T) {
	g := NewWithT(t)

	uintPointer := ToPointer(uint(123456))

	g.Expect(uintPointer).ShouldNot(BeNil())
	g.Expect(*uintPointer).Should(BeIdenticalTo(uint(123456)))
}

func Test_ToPointerWith(t *testing.T) {
	g := NewWithT(t)

	stringPointer := ToPointer("abcd123456")

	g.Expect(stringPointer).ShouldNot(BeNil())
	g.Expect(*stringPointer).Should(BeIdenticalTo("abcd123456"))
}
