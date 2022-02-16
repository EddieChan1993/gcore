package gprometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var histograms = map[string]prometheus.Histogram{}
var counters = map[string]prometheus.Counter{}
var gauges = map[string]prometheus.Gauge{}

func InitWithGin(e *gin.Engine) {
	e.GET("/metrics",
		func(c *gin.Context) {
			promhttp.Handler().ServeHTTP(c.Writer, c.Request)
		})
	e.GET("/health_check", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"err_code": 0}) })
}

func HistogramTime(name string, timeBefore time.Time) {
	if histograms[name] == nil {
		histograms[name] = promauto.NewHistogram(prometheus.HistogramOpts{
			Name: name,
		})
	}
	value := float64(time.Now().UnixNano()-timeBefore.UnixNano()) / 1000000000
	histograms[name].Observe(value)
}

func Histogram(name string, value float64) {
	if histograms[name] == nil {
		histograms[name] = promauto.NewHistogram(prometheus.HistogramOpts{
			Name: name,
		})
	}
	histograms[name].Observe(value)
}

func CounterOne(name string) {
	if counters[name] == nil {
		counters[name] = promauto.NewCounter(prometheus.CounterOpts{
			Name: name,
		})
	}
	counters[name].Inc()
}

func CounterValue(name string, value float64) {
	if counters[name] == nil {
		counters[name] = promauto.NewCounter(prometheus.CounterOpts{
			Name: name,
		})
	}
	counters[name].Add(value)
}

func Gauge(name string, value float64) {
	if gauges[name] == nil {
		gauges[name] = promauto.NewGauge(prometheus.GaugeOpts{
			Name: name,
		})
	}
	gauges[name].Set(value)
}
