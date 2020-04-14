package metrick

import (
	"time"
)

type Day struct {
	Period
	QDays [CONST_SECTIONS_DAY]*QDay
}

// QDay returns the quarter-day associated with the supplied time.
func (self *Day) QDay(t time.Time) *QDay {

	x := computeSegment(t.Hour(), CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)

	self.Lock()
	defer self.Unlock()

	qday := self.QDays[x]
	if qday == nil {
		qday = &QDay{
			NewPeriod(x, self),
			[CONST_SECTIONS_QDAY]*Hour{},
		}
		self.QDays[x] = qday
	}

	return qday
}
