package gmetrics

import (
	"git.dhgames.cn/svr_comm/gcore/glog"
	"testing"
)

func TestStartMetrics(t *testing.T) {
	glog.ResetToDevelopment()
	StartMetrics(NewMetricOnline())
}
