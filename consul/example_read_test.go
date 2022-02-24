package consul_test

import (
	"github.com/EddieChan1993/gcore/consul"
	"github.com/EddieChan1993/gcore/glog"
)

type config map[string]interface{}

func (c *config) Reload() {
	glog.Infow("reload", "config", c)
}

var cfg config = make(config)

func ExampleWatchConfig() {
	glog.ResetToDevelopment()
	err := consul.WatchConfig(&cfg)
	if err != nil {
		glog.Fatalw("failed to wath config", "err", err)
	} else {
		glog.Infow("succeed to wath config", "data", cfg)
	}
	select {}
}
