package metrick

import (
	"time"
)

type QHour struct {
	Period
	Minutes [CONST_SECTIONS_QHOUR]*Minute
}

// Minute returns the minute associated with the supplied time.
func (self *QHour) Minute(t time.Time) *Minute {

	x := t.Minute() % CONST_SECTIONS_QHOUR

	self.Lock()
	defer self.Unlock()

	minute := self.Minutes[x]
	if minute == nil {
		minute = &Minute{
			NewPeriod(t.Minute(), self),
			[CONST_SECTIONS_MINUTE]*QMinute{},
		}
		self.Minutes[x] = minute
	}

	return minute
}
