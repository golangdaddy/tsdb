package metrick

import (
	"os"
	"fmt"
	"time"
	"sync"
	"encoding/gob"
)

// OnInterval returns true every 15th minute on the first second.
func OnInterval(t time.Time) bool {
	t = t.UTC()
	return t.Minute() % 15 == 0 && t.Second() % 60 == 0
}

// DumpFilesToDisk finds all the metrics, and saves them in quarter hour chunks.
func (self *Tree) BackupToDisk(t time.Time) (error, *BackupQuery) {

	fmt.Println("STARTING FILE DUMP", t)
	_, _, _, _, _, qhour, _, _, _ := self.When(t)
	q := &BackupQuery{
		tree: self,
		encname: fmt.Sprintf(
			"%d_%02d_%02d_%02d_",
			t.Year(),
			int(t.Month()),
			t.Day(),
			t.Hour(),
		),
	}
	return q.Do(qhour, nil), q
}

type BackupQuery struct {
	tree *Tree
	added int64
	encname string
	sync.RWMutex
}

// Do executes the backupsquery
func (self *BackupQuery) Do(period interface{}, enc *gob.Encoder) error {

	switch p := period.(type) {

	case *Year:
		for _, x := range p.Months {
			if x == nil { continue }
			self.Do(x, nil)
		}

	case *Month:
		for _, x := range p.Days {
			if x == nil { continue }
			self.Do(x, nil)
		}

	case *Day:
		for _, x := range p.QDays {
			if x == nil { continue }
			self.Do(x, nil)
		}

	case *QDay:
		for _, x := range p.Hours {
			if x == nil { continue }
			self.Do(x, nil)
		}

	case *Hour:

		os.Mkdir(
			fmt.Sprintf("%s/backups", self.tree.settings.VolumePath),
			0666,
		)

		for _, x := range p.QHours {
			if x == nil { continue }
			//fmt.Println("LOOKING IN HOUR", n)
			self.Do(x, nil)
		}

	case *QHour:

		encname := fmt.Sprintf("%s/backups/%s%02d.bak", self.tree.settings.VolumePath, self.encname, p.Index)

		file, err := os.Create(encname)
		if err != nil {
			return err
		}
		enc := gob.NewEncoder(file)
		for _, x := range p.Minutes {
			if x == nil {
				//fmt.Println("MINUTE NIL", n)
				continue
			} else {
				//fmt.Println("LOOKING IN MINUTE", n)
			}
			self.Do(x, enc)
		}
		file.Sync()
		file.Close()

		fmt.Println("SAVED METRICS:", self.added)
		fmt.Println("WRITTEN FILE:", encname)

	case *Minute:
		for n, x := range p.QMinutes {
			if x == nil {
				fmt.Println("QMINUTE NIL", n)
				continue
			} else {
				//fmt.Println("LOOKING IN QMINUTE")
			}
			self.Do(x, enc)
		}

	case *QMinute:

		for _, x := range p.Seconds {
			if x == nil {
				//fmt.Println("SECOND NIL", n)
				continue
			}

			self.doSecond(x, enc)
		}

	default:

		fmt.Println(period)

		panic("UNREACHABLE CODE")

	}

	return nil
}

func (self *BackupQuery) doSecond(period *Second, enc *gob.Encoder) {

	// range over each group of metrics
	period.RLock()
	defer period.RUnlock()

	//s := fmt.Sprintf("METRICS IN SECOND (%d)", len(period.Metrics))

	for _, metricMap := range period.Metrics {

		//s +=  ", " + config.Name

		for metric, _ := range metricMap {

			err := metric.Serialise(enc)
			if err != nil {
				panic(err)
			}

			self.Lock()
			self.added++
			self.Unlock()

		}

	}

	//fmt.Println(s)
}
