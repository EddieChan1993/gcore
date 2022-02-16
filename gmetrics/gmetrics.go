package gmetrics

import (
	"net"
	"time"
)

type IMCollector interface {
	Occur() bool // 是否触发发送
	Body() []*MetricsReq
}

type MetricsBase struct {
	next time.Time
	step int64
}

func NewMetricsBase(step int64) MetricsBase {
	interval := time.Duration(step) * time.Second
	return MetricsBase{
		step: step,
		next: time.Now().Add(interval),
	}
}

func (b *MetricsBase) Occur() bool {
	if b.next.After(time.Now()) {
		return false
	}
	interval := time.Duration(b.step) * time.Second
	b.next = time.Now().Add(interval)
	return true
}

type MetricsReq struct {
	// Metric 统计值命名
	Metric string
	// Value 统计值
	Value int64
	// TAGS 标签，非空，自定使用但不超过5个，格式 "${tagName}=${tagValue},${tagName}=${tagValue}..."
	TAGS string

	// ---- 以下为默认值，项目不用关心 ----

	// ContentType 固定值GAUGE
	ContentType string
	// Endpoint 机器hostname
	Endpoint string
	// Timestamp 发送时间戳，秒
	Timestamp int64
}

func (m *MetricsReq) Fill() *MetricsReq {
	m.Endpoint = localIPString()
	m.ContentType = "GAUGE" // 默认值
	m.Timestamp = time.Now().Unix()
	return m
}

// SetPushURL 设置推送地址，不使用默认127.0.0.1
func SetPushURL(url string) {
	pushURL = url
}

// StartMetrics 启动Metrics
func StartMetrics(mods ...IMCollector) {
	newMetricsMgr(mods...)
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				mgr.run()
			}
		}
	}()
}

// LocalIP 获取本地ip
func localIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return nil, nil
}

// localIPString 获取本地ip string
func localIPString() string {
	ip, err := localIP()
	if err != nil {
		return ""
	}
	if ip == nil {
		return ""
	}
	return ip.String()
}
