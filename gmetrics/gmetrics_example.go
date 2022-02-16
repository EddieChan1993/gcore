package gmetrics

import "math/rand"

const (
	QOnline = "metrics.online.count"
)

type MetricOnline struct {
	MetricsBase
	Online int64
}

func NewMetricOnline() *MetricOnline {
	ret := &MetricOnline{}
	ret.MetricsBase = NewMetricsBase(5)
	return ret
}

func (m *MetricOnline) Body() []*MetricsReq {
	m.Online = rand.Int63n(100)
	req := &MetricsReq{
		TAGS:   "service_name=gcore",
		Metric: QOnline,
		Value:  m.Online,
	}
	return []*MetricsReq{req.Fill()}
}
