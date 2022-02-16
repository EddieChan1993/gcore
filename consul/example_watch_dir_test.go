package consul_test

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	"git.dhgames.cn/svr_comm/gcore/glog"
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
