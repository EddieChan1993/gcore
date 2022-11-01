var static = &StaticCfg{}

type StaticCfg struct {
	DebugMode             bool
	WaitReconnectWaitTime uint32 //离线后，等待重连时间，单位s
	Redis                 string //redis地址
	GPayCluster           string //gpay集群
	GPayDC                string //gpay数据中心
}

func NewStatic() *StaticCfg {
	return static
}

func (this_ *StaticCfg) Reload() {
	static = this_
	klog.Info("reload consul 静态配置完成")
}

func (this_ *StaticCfg) New() consul.IConfig {
	return &StaticCfg{}
}

//==================== 调用函数 ====================//

//StaticGPay gpay配置
func StaticGPay() (cluster, dc string) {
	return static.GPayCluster, static.GPayDC
}

//StaticRedisUrl redis地址
func StaticRedisUrl() string {
	return static.Redis
}

//StaticWaitReconnectWaitTime 等待重连时间
func StaticWaitReconnectWaitTime() uint32 {
	return static.WaitReconnectWaitTime
}
