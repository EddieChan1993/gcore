package blog

import (
	"strconv"
	"time"
)

// 设置日志推送模式为，在内网（开发环境）打日志到测试环境
func SetUrlTestInternel() {
	logUrl = urlTestInternal
}

// 设置日志推送模式为，在任意生产环境打日志到测试环境
func SetUrlTestProduction() {
	logUrl = urlTestProduction
}

// 设置日志推送模式为，在大陆集群打日志到生产环境
func SetUrlProductionMainland() {
	logUrl = urlProductionMainland
}

// 设置日志推送模式为，在海外集群打日志到生产环境
func SetUrlProductionOversea() {
	logUrl = urlProductionOversea
}

// 推送日志，会阻塞直到成功或失败。
func Push(eventCode, eventType, eventName, gameCd, eventValue string) error {
	bilog := &biLog{
		EventCode:  eventCode,
		EventType:  eventType,
		EventName:  eventName,
		GameCd:     gameCd,
		EventValue: eventValue,
		CreateTs:   strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
	}
	return bilog.push()
}
