package metrick

import (
	"time"
)

type Minute struct {
	Period
	QMinutes [CONST_SECTIONS_MINUTE]*QMinute
}

// QMinute returns the quarter-minute associated with the supplied time.
func (self *Minute) QMinute(t time.Time) *QMinute {

	x := computeSegment(t.Second(), CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)

	self.Lock()
	defer self.Unlock()

	qminute := self.QMinutes[x]
	if qminute == nil {
		qminute = &QMinute{
			NewPeriod(x, self),
			[CONST_SECTIONS_QMINUTE]*Second{},
		}
		self.QMinutes[x] = qminute
	}

	return qminute
}
