package remotewrite

import (
	"testing"
	"time"
)

func TestHttpClient_RemoteWrite(t *testing.T) {
	c, err := NewClient("http://<REPLACE_WITH_IP_ADDRESS>:9090/api/v1/write", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	metrics := []MetricPoint{
		{
			Metric: "test1",
			Tags:   map[string]string{"env": "test_env1", "job": "test_job1"},
			Time:   time.Now().Add(-1 * time.Minute).Unix(),
			Value:  2,
		}, {
			Metric: "test2",
			Tags:   map[string]string{"env": "test_env2", "job": "test_job2"},
			Time:   time.Now().Add(-2 * time.Minute).Unix(),
			Value:  3,
		}, {
			Metric: "test1",
			Tags:   map[string]string{"env": "test_env3", "job": "test_job3"},
			Time:   time.Now().Add(-3 * time.Minute).Unix(),
			Value:  4,
		},
	}
	err = c.RemoteWrite(metrics)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("end..")
}
