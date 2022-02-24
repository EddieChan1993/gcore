package consul_test

import (
	"github.com/EddieChan1993/gcore/consul"
	"github.com/EddieChan1993/gcore/glog"
)

type dirConfigs map[string]interface{}

func (d dirConfigs) Reload(key string, data []byte) {
	d[key] = string(data)
	glog.Infow("realod dir", "key", key, "data", string(data))
}

var dc dirConfigs

func ExampleWatchDir() {
	glog.ResetToDevelopment()
	if err := consul.WatchDir("lwk_dev/login", dc); err != nil {
		glog.Fatalw("failed to watch dir", "err", err)
	}
	select {}
}
