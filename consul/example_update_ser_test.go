package consul_test

import (
	"github.com/gcore/consul"
	"github.com/gcore/glog"
)

type dyConfig map[string]interface{}

func (d *dyConfig) Reload() {
	glog.Infow("reald", "new config", d)
}

var dConfig dyConfig

func ExampleUpdateDynamicConfigByService() {
	glog.ResetToDevelopment()

	ser := &consul.ServiceInfo{Cluster: "yht", Service: "gcore", Index: 1}

	// 假设此时配置中心的json数据为 {"a":1}
	if err := consul.WatchDynamicConfigByService(ser, &dConfig); err != nil {
		glog.Warnw("failed to watch dynamic config", "err", err)
		return
	}

	// 执行成功后，dConfig.Reload会被调用，打印 {"new config": {"a":2}}
	if err := consul.UpdateDynamicConfigByService(ser, map[string]int{"a": 2}); err != nil {
		glog.Warnw("failed to update dynamic config", "err", err)
	} else {
		glog.Infow("succeed to update dynamic config")
	}
	select {}
}
