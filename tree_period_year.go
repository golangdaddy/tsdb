package metrick

import (
	"time"
)

type Year struct {
	Period
	Months [CONST_MEMBERS_YEAR]*Month
}

// Month returns the month associated with the supplied time.
func (self *Year) Month(t time.Time) *Month {

	x := int(t.Month()) - 1

	self.Lock()
	defer self.Unlock()

	month := self.Months[x]
	if month == nil {
		month = &Month{
			NewPeriod(x, self),
			[CONST_MEMBERS_MONTH]*Day{},
		}
		self.Months[x] = month
	}

	return month
}
