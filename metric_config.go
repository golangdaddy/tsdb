package metrick

import (
	"github.com/segmentio/fasthash/fnv1a"
)

// Config creates a metric configuration with the given name.
func (self *Tree) MetricConfig(name string) *MetricConfig {
	self.Lock()
	defer self.Unlock()
	metricConfig, ok := self.configs[name]
	if !ok {
		metricConfig = NewConfig(name)
		self.configs[name] = metricConfig
	}
	return metricConfig
}

func NewConfig(name string) *MetricConfig {
	return &MetricConfig{
		hashableUint64: hashableUint64(fnv1a.HashString64(name)),
		Name: name,
	}
}

type MetricConfig struct {
	hashableUint64
	Name string
}
