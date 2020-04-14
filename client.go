package metrick

import (
	"io"
	"os"
	"fmt"
	"time"
	"bytes"
	"testing"
	"encoding/gob"
	//
	"github.com/golangdaddy/go.uuid"
)

const (
	CONST_QUEUE_LENGTH = 500
)

type ClientSettings struct {
	Testing bool
	VolumePath string
	BatchDuration time.Duration
}

func NewClient(settings ClientSettings) (*Client, error) {

	var uid string
	for {
		u, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		} else {
			uid = u.String()
			break
		}
	}

	defaultSettings := ClientSettings{
		VolumePath: settings.VolumePath,
		BatchDuration: 5 * time.Second,
	}
	if len(defaultSettings.VolumePath) == 0 {
		defaultSettings.Testing = true
	}
	if settings.BatchDuration > 0 {
		defaultSettings.BatchDuration = settings.BatchDuration
	}

	fmt.Println("MetricK_client: USING SETTINGS", defaultSettings)

	client := &Client{
		settings: defaultSettings,
		clientID: uid,
		batchQueue: make(chan *Metric, CONST_QUEUE_LENGTH),
	}

	if !client.settings.Testing {
		go client.batchLogs()
	}

	return client, nil
}

func (client *Client) batchLogs() {

	var received int

	var buf *bytes.Buffer
	var enc *gob.Encoder

	for {
		t := time.After(client.settings.BatchDuration)
		filename := fmt.Sprintf("%s/%s_%v.gobs", client.settings.VolumePath, client.clientID, time.Now().UTC().Unix())

		for {
			select {

			case metric := <- client.batchQueue:

				if enc == nil {
					buf = bytes.NewBuffer(nil)
					enc = gob.NewEncoder(buf)
				}

				received++
				if err := metric.Serialise(enc); err != nil {
					panic(err)
				}
				continue

			case <- t:

				if buf != nil {
					file, err := os.Create(filename)
					if err != nil {
						panic(err)
					}
					_, err = io.Copy(file, buf)
					if err != nil {
						panic(err)
					}
					if err := file.Close(); err != nil {
						panic(err)
					}
					//fmt.Println("WROTE", n, "BYTES TO LOG")
				}

				buf = nil
				enc = nil

			}

			break
		}
	}

}

type Client struct {
	t *testing.T
	settings ClientSettings
	clientID string
	batchQueue chan *Metric
}

// NewMetric creates a metric at the current moment with the specified labels.
func (client *Client) NewMetric(config *MetricConfig, v interface{}, labels ...*Label) *Metric {

	return TestMetric(time.Now(), config, v, labels...)
}

func (client *Client) Publish(metric *Metric) {
	client.batchQueue <- metric
}
