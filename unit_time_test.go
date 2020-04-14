package metrick

import (
	"fmt"
	"time"
	"testing"
)

func TestTheTime(t *testing.T) {

	tree := NewTree(
		TreeSettings{},
	)

	config := tree.MetricConfig("http_status_200")

	startTime := time.Now().UTC()
	metric := TestMetric(startTime, config, 1)
	endTime := metric.GetTime()

	if startTime.Unix() != endTime.Unix() {
		fmt.Println(startTime, endTime)
		t.Fail()
	}

}
