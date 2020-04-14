package metrick

import (
	"fmt"
	"time"
	"sync"
	"testing"
	//
	"github.com/fortytw2/leaktest"
)

type States struct {
	Ready bool
	sync.RWMutex
}

type Stats struct {
	Added int64
	Finished bool
	sync.RWMutex
}

type TreeSettings struct {
	Testing bool
	VolumePath string
	BufferDuration time.Duration
}

// NewTree creates an instance of the time series database.
func NewTree(settings TreeSettings) *Tree {

	defaultSettings := TreeSettings{
		VolumePath: settings.VolumePath,
		BufferDuration: 24 * time.Hour,
	}
	if len(settings.VolumePath) == 0 {
		defaultSettings.Testing = true
	}
	if settings.BufferDuration > 0 {
		defaultSettings.BufferDuration = settings.BufferDuration
	}

	fmt.Println("MetricK_server: USING SETTINGS", defaultSettings)

	tree := &Tree{
		settings: defaultSettings,
		configs: map[string]*MetricConfig{},
		Labels: map[string]*Label{},
		years: map[int]*Year{},
		quitChan: make(chan struct{}),
		stats: &Stats{},
		states: &States{},
	}
	if !tree.settings.Testing {
		// number of routines to clean up after
		tree.goroutines = 2

		// ingest new files into memory
		go tree.Ingest()

		// keep track of time so we can trigger a backup
		go func (qc chan struct{}) {
			for t := range time.NewTicker(time.Second).C {

				t = t.UTC()

				// provide a way to clean up goroutine asap
				select {
				case <- qc:
					fmt.Println("STOPPING BACKUPS")
					return
				default:
				}
				// every 15th minute on the first second make a backup
				if OnInterval(t) {

					t = t.Add(CONST_BACKUP_OFFSET)

					// dump encs to disk
					if err, _ := tree.BackupToDisk(t); err != nil {
						panic(err)
					}

					// remove the old stuff from memory
					if err, _ := tree.CleanQHourFromMemory(
						t.Add(-tree.settings.BufferDuration),
					); err != nil {
						panic(err)
					}
				}
			}
		}(tree.quitChan)
	}
	return tree
}

type Tree struct {
	settings TreeSettings
	configs map[string]*MetricConfig
	Labels map[string]*Label
//	Nodes map[string]*Node
	years map[int]*Year
	stats *Stats
	states *States
	goroutines int
	quitChan chan struct{}
	sync.RWMutex
}

func (self *Tree) Stop(t *testing.T) {
	for x := 0; x < self.goroutines; x++ {
		self.quitChan <- struct{}{}
		fmt.Println("STOPPED GOROUTINE", x)
	}
	if t != nil {
		leaktest.Check(t)()
	}
}

func (self *Tree) Stats() *Stats {
	return self.stats
}

func (self *Tree) Inspect() map[int]map[int]map[int]map[int]bool {

	t := time.Now().UTC()

	year, _, _, _, _, _, _, _, _ := self.When(t)

	data := map[int]map[int]map[int]map[int]bool{}

	for _, month := range year.Months {

		if month == nil {
			continue
		}
		if data[month.Index] == nil {
			data[month.Index] = map[int]map[int]map[int]bool{}
		}

		for _, day := range month.Days {

			if day == nil {
				continue
			}
			if data[month.Index][day.Index] == nil {
				data[month.Index][day.Index] = map[int]map[int]bool{}
			}

			for _, qday := range day.QDays {

				if qday == nil {
					continue
				}
				if data[month.Index][day.Index][qday.Index] == nil {
					data[month.Index][day.Index][qday.Index] = map[int]bool{}
				}

				for _, hour := range qday.Hours {
					if hour == nil {
						continue
					}
					data[month.Index][day.Index][qday.Index][hour.Index] = true
				}
			}
		}
	}

	return data
}
