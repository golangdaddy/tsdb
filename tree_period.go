package metrick

import (
//	"fmt"
	"time"
	"sync"
	"hash"
	//
	"github.com/steakknife/bloomfilter"
)

func computeSegment(x, xmax, sections int) int {
	return int((float64(x) / float64(xmax)) * (float64(xmax) / float64(sections)))
}

// NewPeriod initialises a new period instance.
func NewPeriod(index int, period Node) Period {
	bf, err := bloomfilter.NewOptimal(10000, 0.0001)
	if err != nil {
		panic(err)
	}
	return Period {
		parent: period,
		Index: index,
		Filter: bf,
	}
}

type Period struct {
	parent Node
	Index int
	n int64
	*bloomfilter.Filter
	sync.RWMutex
}

func (self *Period) GetIndex() int {
	return self.Index
}

func (self *Period) Add(h hash.Hash64) {
	self.Lock()
	defer self.Unlock()
	self.Filter.Add(h)
	self.n++
}

func (self *Period) Parent() Node {
	return self.parent
}

// Year returns the top level period.
func (self *Tree) Year(t time.Time) *Year {
	x := t.Year()

	self.Lock()
	defer self.Unlock()

	year := self.years[x]
	if year == nil {
		year = &Year{
			NewPeriod(x, nil),
			[CONST_MEMBERS_YEAR]*Month{},
		}
		self.years[x] = year
	}

	return year
}

type Second struct {
	Period
	Metrics map[*MetricConfig]map[*Metric]struct{}
}
