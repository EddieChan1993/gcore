package gmetrics

import (
	"github.com/gcore/glog"
	"testing"
)

func TestStartMetrics(t *testing.T) {
	glog.ResetToDevelopment()
	StartMetrics(NewMetricOnline())
}
