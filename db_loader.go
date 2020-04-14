package metrick

import (
	"io"
	"os"
	"fmt"
	"time"
	"encoding/gob"
)

// Loadbackupfromfile puts the contents of a file into memory, and returns -1 if the file doesn't exist
func (self *Tree) LoadBackupFile(filename string, testing bool) (n int) {

	file, err := os.Open(filename)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		fmt.Println(err)
		return -1
	}

	decoder := gob.NewDecoder(
		file,
	)

	for {
		exportMetric := X{}
		err := decoder.Decode(&exportMetric)
		if err != nil {
			if err == io.EOF {
				//fmt.Println("EOF AFTER", n)
				break
			}
			fmt.Println(err)
			return -1
		}

		metric := self.ToMetric(exportMetric)

		t := metric.GetTime()

		// clean up old backups which are ancient (40 days)
		if time.Since(t) > 40 * 24 * time.Hour {
			// close file before removing it
			file.Close()
			fmt.Println("IGNORING & REMOVING OLD BACKUP:", filename, t)
			err = os.Remove(filename)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			return
		}

		// load metric into memory
		self.AddMetric(
			metric,
		)

		n++

	}

	return
}
