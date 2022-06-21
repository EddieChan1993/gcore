package gmetrics

import (
	"github.com/EddieChan1993/gcore/glog"
	"testing"
)

func TestStartMetrics(t *testing.T) {
	glog.ResetToDevelopment()
	StartMetrics(NewMetricOnline())
}
