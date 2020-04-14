package metrick

import (
	"fmt"
	"time"
	"testing"
)

func TestTree1(t *testing.T) {

	tree := NewTree(
		TreeSettings{},
	)
/*
	client, err := NewClient(
		ClientSettings{},
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	*/

	configs := struct{
		transaction *MetricConfig
	}{
		tree.MetricConfig("transaction"),
	}

	accounts := []*Label{
		tree.Label("account1"),
		tree.Label("account2"),
		tree.Label("account3"),
		tree.Label("account4"),
		tree.Label("account5"),
	}

	start := time.Time{}

	for x := 0; x < 1000; x++ {

		for _, account := range accounts {

			for _, metric := range []*Metric{
				TestMetric(
					start,
					configs.transaction,
					fmt.Sprintf("%s_%d.tx", account.Label, x),
					account,
				),
			} {
				tree.AddMetric(metric)
				start = start.Add(time.Second)

				fmt.Println(metric.Value)
			}
			start = start.Add(time.Second)

		}
	}

	fmt.Println(tree.Inspect())

	_, _, day, _, _, _, _, _, _:= tree.When(start)

	query := NewQuery()
	query.Or = true
	query.Limit = 10
	query.Metrics = []*MetricConfig{
		configs.transaction,
	}
	query.Labels = accounts

	query.Do(day)

	z := query.Results.Sort()

	fmt.Println("RESULTS", len(z), z)

}
