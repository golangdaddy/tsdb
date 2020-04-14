package metrick

import (
	"time"
	"testing"
)

func TestPeriods(t *testing.T) {

	tree := NewTree(
		TreeSettings{},
	)
	x := time.Time{}
	x = x.Add(24 * time.Hour)

	year, month, day, qday, hour, qhour, minute, qminute, second := tree.When(x)

	t.Run(
		"TEST PERIOD YEAR",
		func (t *testing.T) {

			if year.Index != x.Year() {
				t.Errorf("INVALID INDEX: %d", year.Index)
				t.Fail()
			}

			if len(year.Months) != CONST_MEMBERS_YEAR {
				t.Errorf("INVALID SECTIONS: %v, EXPECTING: %v", len(year.Months), CONST_MEMBERS_YEAR)
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD MONTH",
		func (t *testing.T) {

			if month.Index != int(x.Month()) - 1 {
				t.Fail()
			}

			if len(month.Days) != CONST_MEMBERS_MONTH {
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD DAY",
		func (t *testing.T) {

			if day.Index != x.Day() - 1 {
				t.Fail()
			}

			if len(day.QDays) != CONST_SECTIONS_DAY {
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD QDAY",
		func (t *testing.T) {

			if qday.Index != computeSegment(x.Hour(), 24, CONST_SECTIONS_DAY)  {
				t.Fail()
			}

			if len(qday.Hours) != CONST_SECTIONS_QDAY {
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD HOUR",
		func (t *testing.T) {

			if hour.Index != x.Hour() {
				t.Errorf("INVALID INDEX: %v EXPECTING %v", hour.Index, x.Hour())
				t.Fail()
			}

			if len(hour.QHours) != CONST_SECTIONS_HOUR {
				t.Errorf("INVALID SECTIONS: %v, EXPECTING: %v", len(hour.QHours), CONST_SECTIONS_HOUR)
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD QHOUR",
		func (t *testing.T) {

			if qhour.Index != computeSegment(x.Minute(), 60, CONST_SECTIONS_HOUR) {
				t.Fail()
			}

			if len(qhour.Minutes) != CONST_SECTIONS_QHOUR {
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD MINUTE",
		func (t *testing.T) {

			if minute.Index != x.Minute() {
				t.Error("INVALID INDEX")
				t.Fail()
			}

			if len(minute.QMinutes) != CONST_SECTIONS_MINUTE {
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD QMINUTE",
		func (t *testing.T) {

			n := computeSegment(x.Second(), CONST_MEMBERS_MINUTE, CONST_SECTIONS_MINUTE)
			if qminute.Index != n {
				t.Errorf("INVALID INDEX: %d EXPECTING: %v FOR SECOND: %v", qminute.Index, n, x.Minute())
				t.Fail()
			}
			//t.Logf("VALID INDEX %v", qminute.Index)

			if len(qminute.Seconds) != CONST_SECTIONS_QMINUTE {
				t.Errorf("INVALID SECTIONS: %v, EXPECTING: %v", len(qminute.Seconds), CONST_SECTIONS_QMINUTE)
				t.Fail()
			}

		},
	)

	t.Run(
		"TEST PERIOD SECOND",
		func (t *testing.T) {

			if second.Index != x.Second() {
				t.Fail()
			}

			//n := int(float64(int(float64(second.Index) / 60.0)) + float64(second.Index % 15))


			//n := computeSegment(x.Second(), 60, CONST_SECTIONS_MINUTE)


		},
	)

}
