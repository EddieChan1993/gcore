package consul

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/EddieChan1993/gcore/glog"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	query_wait           = "1m"
	query_errorSleepWait = 30 * time.Second
)

// cluster-service-index
func getServiceInfoByPath() (result *ServiceInfo, err error) {
	var dir string
	if dir, err = os.Getwd(); err != nil {
		return nil, fmt.Errorf("faild to get work dir,err=%s", err.Error())
	}

	workDir := path.Base(dir)
	serviceData := strings.Split(workDir, "-")
	if len(serviceData) != 3 {
		return nil, fmt.Errorf("work dir (%s) is not a valid service format", workDir)
	}

	var serviceIndex int
	if serviceIndex, err = strconv.Atoi(serviceData[2]); err != nil {
		return nil, fmt.Errorf("work dir (%s) is not a valid service format", workDir)
	}
	return &ServiceInfo{Cluster: serviceData[0], Service: serviceData[1], Index: serviceIndex}, nil
}

func blockQuery(url string, config IConfig) {
	oldIndex := 0
	oldBytes := []byte{}
	for {
		watchPath := fmt.Sprintf("%s&index=%d&wait=%s", url, oldIndex, query_wait)
		body, header, err := httpGet(watchPath)
		if err != nil {
			glog.Warn("failed to block query", "watch path", watchPath, "err", err)
			time.Sleep(query_errorSleepWait)
			continue
		}

		newIndex := praseConsulIndex(header)
		if newIndex <= oldIndex {
			continue
		}

		oldIndex = newIndex
		if bytes.Equal(oldBytes, body) {
			continue
		}

		newConfig := config.New()
		if err = json.Unmarshal(body, newConfig); err != nil {
			glog.Warnw("fail to block query config", "err", err)
			continue
		}

		newConfig.Reload()

		oldBytes = body
	}
}

func blockQueryDir(url string, config IMutiConfig) {
	oldIndex := 0
	oldBytes := []byte{}
	keyIndexs := map[string]int64{}
	for {
		watchPath := fmt.Sprintf("%s&index=%d&wait=%s", url, oldIndex, query_wait)
		body, header, err := httpGet(watchPath)
		if err != nil {
			glog.Warn("failed to block query", "watch path", watchPath, "err", err)
			time.Sleep(query_errorSleepWait)
			continue
		}

		newIndex := praseConsulIndex(header)
		if newIndex <= oldIndex {
			continue
		}

		oldIndex = newIndex
		if bytes.Equal(oldBytes, body) {
			continue
		}

		type kv struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
			Index int64  `json:"ModifyIndex"`
			data  []byte
		}
		kvs := []kv{}
		if bytes.Equal(body, []byte{}) {
			glog.Warnw("block query dir empty body")
		} else {
			if err = json.Unmarshal(body, &kvs); err != nil {
				glog.Warnw("fail to block query dir", "err", err)
				continue
			}
		}

		// 遍历新的kv，如果版本变更，表示有改动。
		newIndexes := map[string]int64{}
		for _, v := range kvs {
			i := strings.LastIndex(v.Key, "/")
			key := v.Key[i+1:]
			newIndexes[key] = v.Index
			if keyIndexs[key] == v.Index {
				continue
			}
			v.data, err = base64.StdEncoding.DecodeString(v.Value)
			if err != nil {
				glog.Warnw("fail to block query dir", "err", err)
				continue
			}
			keyIndexs[key] = v.Index
			config.Reload(key, v.data)
		}
		for k := range keyIndexs {
			if newIndexes[k] == 0 {
				config.Delete(k)
			}
		}

		oldBytes = body
		keyIndexs = newIndexes
	}
}

func httpGet(url string) ([]byte, http.Header, error) {
	resp, err := http.Get(url)
	defer func() {
		if err == nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return body, resp.Header, err
}

func praseConsulIndex(header http.Header) int {
	index := header.Get("X-Consul-Index")
	if index == "" {
		return 0
	}

	indexInt, err := strconv.Atoi(index)
	if err != nil {
		return 0
	}

	return indexInt
}

func readConfig(url string, ptrToConfig interface{}) error {
	body, _, err := httpGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, ptrToConfig)
	return err
}

func readDir(url string, config IMutiConfig) error {
	body, _, err := httpGet(url)
	if err != nil {
		return err
	}

	type kv struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
		Index int64  `json:"ModifyIndex"`
		data  []byte
	}
	kvs := []kv{}
	if bytes.Equal(body, []byte{}) {
		glog.Warnw("block query dir empty body")
	} else {
		if err = json.Unmarshal(body, &kvs); err != nil {
			glog.Warnw("fail to block query config", "err", err)
			return err
		}
	}

	for _, v := range kvs {
		v.data, err = base64.StdEncoding.DecodeString(v.Value)
		if err != nil {
			return err
		}
		i := strings.LastIndex(v.Key, "/")
		config.Reload(v.Key[i+1:], v.data)
	}
	return nil
}

var httpClient = http.Client{Timeout: time.Second * 1}

func httpPut(url string, config interface{}) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(configBytes))
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("failed to put consul,StatusCode=%d", response.StatusCode)
}

func httpDelete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("failed to delete consul,StatusCode=%d", resp.StatusCode)
}
