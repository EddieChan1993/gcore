package consul_test

import (
	"github.com/gcore/consul"
	"github.com/gcore/glog"
)

type config2 map[string]interface{}

func (c *config2) Reload() {
	glog.Infow("reload", "config", c)
}

var cfg2 config2 = make(config2)

func ExampleWatchConfigByService() {
	glog.ResetToDevelopment()
	serInfo := consul.ServiceInfo{
		Cluster: "pub",
		Service: "gcore",
		Index:   1,
	}
	err := consul.WatchConfigByService(&serInfo, &cfg2)
	if err != nil {
		glog.Fatalw("failed to wath config", "err", err)
	} else {
		glog.Infow("succeed to wath config", "data", cfg2)
	}
	select {}
}
