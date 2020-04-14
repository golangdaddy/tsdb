package metrick

import (
	"time"
)

type QMinute struct {
	Period
	Seconds [CONST_SECTIONS_QMINUTE]*Second
}

// Second returns the second associated with the supplied time.
func (self *QMinute) Second(t time.Time) *Second {

	x := t.Second() % CONST_SECTIONS_QMINUTE

	self.Lock()
	defer self.Unlock()

	second := self.Seconds[x]
	if second == nil {
		second = &Second{
			NewPeriod(t.Second(), self),
			map[*MetricConfig]map[*Metric]struct{}{},
		}
		self.Seconds[x] = second
	}

	return second
}
