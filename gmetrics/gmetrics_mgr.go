package gmetrics

import (
	"bytes"
	"encoding/json"
	"github.com/EddieChan1993/gcore/glog"
	"net/http"
)

var pushURL = "http://127.0.0.1:1988/v1/push"

var mgr *metricsMgr

type metricsMgr struct {
	array []IMCollector
}

func newMetricsMgr(mods ...IMCollector) {
	mgr = &metricsMgr{
		array: mods,
	}
}

func (m *metricsMgr) run() {
	for _, col := range m.array {
		if col == nil {
			continue
		}
		if col.Occur() {
			send(col.Body())
		}
	}
}

func send(array []*MetricsReq) {
	if len(array) == 0 {
		return
	}
	j, err := json.Marshal(array)
	if err != nil {
		glog.Warnf("metrics marshal json err %v", err)
		return
	}
	glog.Infof("metrics json %v", string(j))
	req, err := http.NewRequest(http.MethodPost, pushURL, bytes.NewReader(j))
	if err != nil {
		glog.Warnf("metrics http new request err %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Warnf("metrics send err %v", err)
	}
	glog.Infof("metrics send resp %+v", resp)
}
