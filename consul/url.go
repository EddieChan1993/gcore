package consul

import "fmt"

var (
	url_base          = "http://127.0.0.1:2000/v1/"
	url_staticConfig  = url_base + "kv/app_static_cfg/"
	url_dynamicConfig = url_base + "kv/app_dynamic_cfg/"
)

func buildStaticConfigUrlByPath() (string, error) {
	ser, err := getServiceInfoByPath()
	if err != nil {
		return "", err
	}
	return buildStaticConfigUrl(ser), nil
}

func buildDynamicConfigUrlByPath() (string, error) {
	ser, err := getServiceInfoByPath()
	if err != nil {
		return "", err
	}
	return buildDynamicConfigUrl(ser), nil
}

func buildStaticConfigUrl(serInfo *ServiceInfo) string {
	return buildConfigUrl(url_staticConfig, serInfo)
}

func buildDynamicConfigUrl(serInfo *ServiceInfo) string {
	return buildConfigUrl(url_dynamicConfig, serInfo)
}

func buildConfigUrl(baseUrl string, ser *ServiceInfo) string {
	return fmt.Sprintf("%s%s/%s/%d?raw=true", baseUrl, ser.Cluster, ser.Service, ser.Index)
}

func buildDynamicDirUrl(dirURL string) string {
	return fmt.Sprintf("%s%s", url_dynamicConfig, dirURL)
}

func buildDynamicWatchDirUrl(dirURL string) string {
	if len(dirURL) != 0 && dirURL[len(dirURL)-1] == '/' {
		dirURL = dirURL[:len(dirURL)-1]
	}
	return fmt.Sprintf("%s%s/?recurse=true", url_dynamicConfig, dirURL)
}
