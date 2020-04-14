package metrick

import (
	"time"
)

type QDay struct {
	Period
	Hours [CONST_SECTIONS_QDAY]*Hour
}

// Hour returns the hour associated with the supplied time.
func (self *QDay) Hour(t time.Time) *Hour {

	x := t.Hour() % CONST_SECTIONS_QDAY

	self.Lock()
	defer self.Unlock()

	hour := self.Hours[x]
	if hour == nil {
		hour = &Hour{
			NewPeriod(t.Hour(), self),
			[CONST_SECTIONS_HOUR]*QHour{},
		}
		self.Hours[x] = hour
	}

	return hour
}
