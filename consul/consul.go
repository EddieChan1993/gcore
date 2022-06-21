// Package consul 提供通过consul获取、监控静态配置或者动态配置的方法，以及更新动态配置的方法。
//
// 静态配置：存放在 http://{consul_url}/kv/app_static_cfg/cluster/service/index 的json数据。
//
// 动态配置：存放在 http://{consul_url}/kv/app_dynamic/cluster/service/index 的json数据。
//
// 动态目录：存放在 http://{consul_url}/kv/app_dynamic/{dirURL}/的所有json数据
//
// 其中，静态配置程序只允许读取，不允许修改。只能让运维或者开发通过consul的ui界面修改。
// 动态配置既允许程序读取，也允许修改。一般来说，这种场景是发生在A服务更新动态配置，B服务收到通知，然后改变自己的状态。
//
// 另外，获取配置时，我们提供了两种方法获取。
// 一种是通过显式的传入服务地址获取，一般用于调试或者访问其他服务的配置。
// 另一种是通过当前服务的工作目录去获取，规则是{cluster}-{service}-{index}，比如pub-gcore-1。如果访问自己的服务配置，建议使用这种方法，方便运维部署。
package consul

// IConfig config接口，要求实现Reload方法。
//
// Reload方法会在远端consul配置更新时调用。
type IConfig interface {
	// Reload 注意，该方法是在IConfig值被成功更新之后调用
	Reload()
	// New 创建一个IConfig用于json解码
	New() IConfig
}

// IMutiConfig 多config配置接口，用在监听consul目录处。调用者需要实现realod方法，并在内部完成对数据的更新操作。
type IMutiConfig interface {
	// 每该目录下有配置新增、修改、删除时，此方法会调用。
	// key=配置的key值，比如：com.droidhang.aod.google
	Reload(key string, data []byte)
	Delete(key string)
}

// 服务信息
//
// 配置信息总是需要精确到具体每个服务节点，一个集群会有多个服务，一个服务也会有多个节点。
type ServiceInfo struct {
	// Cluster 集群名
	Cluster string
	// Service 服务名
	Service string
	// Index 服务节点
	Index int
}

// 根据本地路径读取静态配置
//
// ptrToConfig:指向配置结构体的指针，用来接受配置。注意：必须传指针类型，否则会返回序列化错误异常。
func ReadConfig(ptrToConfig interface{}) error {
	url, err := buildStaticConfigUrlByPath()
	if err != nil {
		return err
	}
	return readConfig(url, ptrToConfig)
}

// 根据服务信息读取静态配置
func ReadConfigByService(ser *ServiceInfo, ptrToConfig interface{}) error {
	url := buildStaticConfigUrl(ser)
	return readConfig(url, ptrToConfig)
}

// 根据本地路径监控配置
//
// config: 此处需要需要传入指针类型，否则会返回序列化异常。
//
// 注意：当此方法返回时，如果未报错，则config值保证已经刷新，可以直接使用，并且配置监控也启动。否则，config内容不会刷新，监控也不会开始。
func WatchConfig(config IConfig) error {
	configUrl, err := buildStaticConfigUrlByPath()
	if err != nil {
		return err
	}

	if err = readConfig(configUrl, &config); err != nil {
		return err
	}

	go blockQuery(configUrl, config)
	return nil
}

// 根据服务信息监控静态配置
func WatchConfigByService(serInfo *ServiceInfo, config IConfig) error {
	configUrl := buildStaticConfigUrl(serInfo)
	if err := readConfig(configUrl, config); err != nil {
		return err
	}

	go blockQuery(configUrl, config)
	return nil
}

// 根据本地路径读取动态配置
func ReadDynamicConfig(ptrToConfig interface{}) (err error) {
	url, err := buildDynamicConfigUrlByPath()
	if err != nil {
		return err
	}
	return readConfig(url, ptrToConfig)
}

// 根据服务信息读取动态配置
func ReadDynamicConfigByService(serInfo *ServiceInfo, ptrToConfig interface{}) (err error) {
	url := buildDynamicConfigUrl(serInfo)
	return readConfig(url, ptrToConfig)
}

// 根据本地路径监控动态配置
func WatchDynamicConfig(config IConfig) error {
	configUrl, err := buildDynamicConfigUrlByPath()
	if err != nil {
		return err
	}

	if err = readConfig(configUrl, config); err != nil {
		return err
	}

	go blockQuery(configUrl, config)
	return nil
}

// WatchDynamicConfigByService 根据服务信息监控动态配置
func WatchDynamicConfigByService(serInfo *ServiceInfo, config IConfig) error {
	configUrl := buildDynamicConfigUrl(serInfo)
	if err := readConfig(configUrl, config); err != nil {
		return err
	}

	go blockQuery(configUrl, config)
	return nil
}

// 根据本地路径更新动态配置
//
// config: 需要传入可序列化为jason格式的任意数据
func UpdateDynamicConfig(config interface{}) error {
	url, err := buildDynamicConfigUrlByPath()
	if err != nil {
		return err
	}
	return httpPut(url, config)
}

// 根据服务信息更新动态配置
func UpdateDynamicConfigByService(ser *ServiceInfo, config interface{}) error {
	url := buildDynamicConfigUrl(ser)
	return httpPut(url, config)
}

// 根据dir构成url更新动态配置
func UpdateDynamicConfigByDir(dir string, config interface{}) error {
	url := buildDynamicDirUrl(dir)
	return httpPut(url, config)
}

// GetServiceInfoByPath 解析当前路径，然后返回服务信息
func GetServiceInfoByPath() (service *ServiceInfo, err error) {
	return getServiceInfoByPath()
}

// DeleteDynamicConfigByDir 通过dir删除动态配置
func DeleteDynamicConfigByDir(dir string) error {
	url := buildDynamicDirUrl(dir)
	return httpDelete(url)
}

// WatchDir 监控动态配置目录，会监控配置的增删改,不支持子目录
// dirURL 需要传入需要监控的相对目录地址，比如:lwk_dev/login，在该目录下保存所有的kv
// config 在配置更新时不会自动修改，需要调用者自己实现更新方法
// 如果返回error，watch不会启动。否则会立即刷新所有配置，并开始监控。
func WatchDir(dirURL string, config IMutiConfig) error {
	url := buildDynamicWatchDirUrl(dirURL)
	if err := readDir(url, config); err != nil {
		return err
	}

	go blockQueryDir(url, config)
	return nil
}
