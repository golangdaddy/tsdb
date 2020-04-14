package metrick

import (
	"time"
)

type Hour struct {
	Period
	QHours [CONST_SECTIONS_HOUR]*QHour
}

// QHour returns the quarter-hour associated with the supplied time.
func (self *Hour) QHour(t time.Time) *QHour {

	x := computeSegment(t.Minute(), CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)

	self.Lock()
	defer self.Unlock()

	qhour := self.QHours[x]
	if qhour == nil {
		qhour = &QHour{
			NewPeriod(x, self),
			[CONST_SECTIONS_QHOUR]*Minute{},
		}
		self.QHours[x] = qhour
	}

	return qhour
}
