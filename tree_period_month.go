package metrick

import (
	"time"
)

type Month struct {
	Period
	Days [CONST_MEMBERS_MONTH]*Day
}

// Day returns the day associated with the supplied time.
func (self *Month) Day(t time.Time) *Day {

	x := t.Day() - 1

	self.Lock()
	defer self.Unlock()

	day := self.Days[x]
	if day == nil {
		day = &Day{
			NewPeriod(x, self),
			[CONST_SECTIONS_DAY]*QDay{},
		}
		self.Days[x] = day
	}

	return day
}
