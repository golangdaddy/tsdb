package metrick

import (
	"fmt"
	"time"
	"sync"
)

func (self *Tree) CleanQHourFromMemory(t time.Time) (error, *CleaningQuery) {

	_, _, _, _, _, qhour, _, _, _ := self.When(t)
	q := &CleaningQuery{}
	return q.Do(qhour), q
}

type CleaningQuery struct {
	added int64
	filename string
	sync.RWMutex
}

func (self *CleaningQuery) Do(period interface{}) error {

	switch p := period.(type) {

	case *Year:
		for _, x := range p.Months {
			if x == nil { continue }
			self.Do(x)
		}

	case *Month:
		for _, x := range p.Days {
			if x == nil { continue }
			self.Do(x)
		}

	case *Day:
		for _, x := range p.QDays {
			if x == nil { continue }
			self.Do(x)
		}

	case *QDay:
		for _, x := range p.Hours {
			if x == nil { continue }
			self.Do(x)
		}

	case *Hour:

		for _, x := range p.QHours {
			if x == nil { continue }
			self.Do(x)
		}

	case *QHour:

		p.Lock()
		for x, _ := range p.Minutes {
			p.Minutes[x] = nil
		}
		p.Unlock()

		fmt.Println("CLEANED QHOUR", p.Index)

	default:

		fmt.Println(period)

		panic("UNREACHABLE CODE")

	}

	return nil
}
