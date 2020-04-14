package metrick

import (
	"time"
)

// When returns/creates all of the periods associated with the supplied time.
func (self *Tree) When(t time.Time) (year *Year, month *Month, day *Day, qday *QDay, hour *Hour, qhour *QHour, minute *Minute, qminute *QMinute, second *Second) {

	t = t.UTC()

	year = self.Year(t)
	month = year.Month(t)
	day = month.Day(t)
	qday = day.QDay(t)
	hour = qday.Hour(t)
	qhour = hour.QHour(t)
	minute = qhour.Minute(t)
	qminute = minute.QMinute(t)
	second = qminute.Second(t)

	return
}
