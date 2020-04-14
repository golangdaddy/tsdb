package metrick

// AddMetric applies the metric to all of the periods associated with it's time.
func (self *Tree) AddMetric(metric *Metric) *Second {

	// parse timestamp
	year, month, day, qday, hour, qhour, minute, qminute, second := self.When(
		metric.GetTime(),
	)

	second.Lock()
	if second.Metrics[metric.Config] == nil {
		second.Metrics[metric.Config] = map[*Metric]struct{}{}
	}
	second.Metrics[metric.Config][metric] = struct{}{}
	second.Unlock()

	metric.applyToBloomFilter(second)

	metric.applyToBloomFilter(qminute)

	metric.applyToBloomFilter(minute)

	metric.applyToBloomFilter(qhour)

	metric.applyToBloomFilter(hour)

	metric.applyToBloomFilter(qday)

	metric.applyToBloomFilter(day)

	metric.applyToBloomFilter(month)

	metric.applyToBloomFilter(year)


	self.stats.Lock()
	self.stats.Added++
	self.stats.Unlock()

	//fmt.Println("ADDED MTR", metric.GetTime())

	return second
}
