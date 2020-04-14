package metrick

import (
	"io"
	"os"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/gob"
)

// ToMetric converts an exported metric to an internal one.
func (self *Tree) ToMetric(exp X) *Metric {
	metric := &Metric{
		Time: exp.T,
		Config: self.MetricConfig(exp.C),
		Labels: map[*Label]struct{}{},
		Value: exp.V,
	}
	for _, v := range exp.L {
		metric.Labels[self.Label(v)] = struct{}{}
	}
	return metric
}

// Ingest starts reading metrics from disk into time-series database.
func (self *Tree) Ingest() error {

	var processed int64

	for {
		// provide a way to clean up goroutine asap
		select {
		case <- self.quitChan:
			fmt.Println("STOPPING INGESTING AFTER", processed)
			return nil
		default:
		}

		time.Sleep(time.Second / 3)

		files, err := ioutil.ReadDir(self.settings.VolumePath)
		if err != nil {
			panic(err)
		}

		var p int64
		var i int

		for _, file := range files {

			if file.IsDir() {
				continue
			}
			path := self.settings.VolumePath + "/" + file.Name()

			if n, err := self.ingestFile(path); err != nil {
				fmt.Println(err)
				continue
			} else {
				p += int64(n)
				i++
			}

		}

		if p > 0 {
			fmt.Println("INGESTED", p, "METRICS FROM", i, "FILES:")
		}
		processed += p

	}
}

// ingestFile loads the specified file into the database.
func (self *Tree) ingestFile(path string) (int, error) {

	//fmt.Println("INGESTING", path)

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	decoder := gob.NewDecoder(file)

	var processed int

	for {
		exportMetric := X{}
		err := decoder.Decode(&exportMetric)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		self.AddMetric(
			self.ToMetric(exportMetric),
		)

		processed++
	}

	file.Close()

	//fmt.Println("REMOVING FILE", path)
	if err := os.Remove(path); err != nil {
		return processed, err
	}

	return processed, nil
}
