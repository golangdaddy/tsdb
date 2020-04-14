package metrick

import (
	"fmt"
	"sync"
	"sort"
	"time"
)

func NewQuery() *Query {
	return &Query{
		Results: NewResults(),
	}
}

type Query struct {
	// toggles whether all labels have to match, or just some
	Or bool
	Limit int
	Offset time.Time
	Labels []*Label
	Metrics []*MetricConfig
	Results *Results
	resultsCount int
	sync.RWMutex
}

func (self *Query) NeedsResults() bool {
	self.RLock()
	defer self.RUnlock()
	return self.resultsCount < self.Limit
}

func NewResults() *Results {
	return &Results{
		Labels: map[string]int64{},
		results: []*Metric{},
	}
}

type Result struct {
	Txid string
	Exp int64
}

type Results struct{
	Labels map[string]int64
	count int64
	results []*Metric
	sync.RWMutex
}

func (self *Results) Reset() {
	self.results = []*Metric{}
}

func (self *Results) Sort() []string {
	sort.Slice(
		self.results,
		func (x, y int) bool {
			return self.results[x].Time < self.results[y].Time
		},
	)
	a := make([]string, len(self.results))
	for x, metric := range self.results {
		a[x] = metric.Value.(string)
	}
	return a
}

func (self *Query) Do(period interface{}) {

	switch p := period.(type) {

	case *Year:

		for _, x := range p.Months {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)

		}

	case *Month:

		for _, x := range p.Days {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)

		}

	case *Day:

		for _, x := range p.QDays {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)

		}

	case *QDay:

		for _, x := range p.Hours {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)
		}

	case *Hour:

		for _, x := range p.QHours {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)

		}

	case *QHour:

		for _, x := range p.Minutes {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)
		}

	case *Minute:

		for _, x := range p.QMinutes {

			if !self.NeedsResults() {
				return
			}

			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.Do(x)
		}

	case *QMinute:

		for _, x := range p.Seconds {

			if !self.NeedsResults() {
				return
			}
			if x == nil {
				continue
			}
			if !self.hasWarmLead_Metric(x) {
				continue
			}
			if len(self.Labels) > 0 {
				if !self.hasWarmLead_Label(x) {
					continue
				}
			}

			self.doSecond(x)
		}

	case *Second:

		self.doSecond(p)

	default:

		fmt.Println(period)

		panic("UNREACHABLE CODE")

	}

}

func (self *Query) doSecond(period *Second) {

	period.RLock()
	defer period.RUnlock()

	// range over each metric group
	for _, config := range self.Metrics {

		for metric, _ := range period.Metrics[config] {

			if !self.Or {
				for _, label := range self.Labels {
					if !metric.HasLabel(label) {
						continue
					}
				}
			}

			self.Results.Lock()

				for label, _ := range metric.Labels {
					self.Results.Labels[label.Label]++
				}
				self.Results.results = append(
					self.Results.results,
					metric,
				)
				self.resultsCount++

			self.Results.Unlock()

			if !self.NeedsResults() {
				return
			}

		}

	}

}

func (self *Query) hasWarmLead_Label(period Node) bool {

	// OR query
	if self.Or {
		for _, label := range self.Labels {
			if period.Contains(label) {
				return true
			}
		}
		return false
	}

	// AND query
	for _, label := range self.Labels {
		if !period.Contains(label) {
			return false
		}
	}
	return true
}

func (self *Query) hasWarmLead_Metric(period Node) bool {

	for _, metric := range self.Metrics {
		if period.Contains(metric) {
			return true
		}
	}

	return false
}
