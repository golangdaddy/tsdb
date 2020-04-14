package metrick

import (
	"time"
	"encoding/gob"
)

// TestMetric creates a metric at the specified time UTC, with the supplied metric config and labels.
func TestMetric(t time.Time, config *MetricConfig, v interface{}, labels ...*Label) *Metric {

	metric := &Metric{
		Time: t.UTC().Unix(),
		Labels: map[*Label]struct{}{},
		Config: config,
		Value: v,
	}
	for _, label := range labels {
		metric.Labels[label] = struct{}{}
	}
	return metric
}

// exported version of metric
type X struct {
	T int64
	C string
	L []string
	V interface{}
}

// Serialise encodes the exported metric to gob.
func (self X) Serialise(enc *gob.Encoder) error {
	return enc.Encode(self)
}

type Metric struct {
	Time int64
	Labels map[*Label]struct{}
	Config *MetricConfig
	Value interface{}
}

// GetTime converts the metric timestamp to a real UTC time.
func (self *Metric) GetTime() time.Time {
	return time.Unix(self.Time, 0).UTC()
}

// HasLabel returns whether a metric contains a certain label.
func (self *Metric) HasLabel(label *Label) bool {
	_, ok := self.Labels[label]
	return ok
}

//	applyToBloomFilter adds the labels and the metric name to the bloom filter for this period node.
func (metric *Metric) applyToBloomFilter(node Node) {
	// add labels
	for label, _ := range metric.Labels {
		//fmt.Println("ADDING LABEL: "+label.Label)
		node.Add(label)
	}
	// add metric name
	node.Add(metric.Config)
}

// Serialise encodes an internal metric into gob.
func (self *Metric) Export() X {

	var x int
	labels := make([]string, len(self.Labels))
	for label, _ := range self.Labels {
		labels[x] = label.Label
		x++
	}

	return X{
		T: self.Time,
		L: labels,
		V: self.Value,
		C: self.Config.Name,
	}
}

// Serialise encodes an internal metric into gob.
func (self *Metric) Serialise(enc *gob.Encoder) error {

	var x int
	labels := make([]string, len(self.Labels))
	for label, _ := range self.Labels {
		labels[x] = label.Label
		x++
	}

	exportedMetric := X{
		T: self.Time,
		L: labels,
		V: self.Value,
		C: self.Config.Name,
	}

	return exportedMetric.Serialise(enc)
}
