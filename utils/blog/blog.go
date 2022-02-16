package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type biLog struct {
	EventCode  string `json:"event_code"`
	EventType  string `json:"event_type"`
	EventName  string `json:"event_name"`
	GameCd     string `json:"game_cd"`
	EventValue string `json:"event_value,omitempty"`
	CreateTs   string `json:"create_ts"`
}

const (
	warnCacheLogCount = 10
	maxCacheLogCount  = 50

	urlTestInternal       = "http://ulog-inner-test.dhgames.cn:8180/inner/push_log"
	urlTestProduction     = "https://ulog-test.dhgames.cn:8181/external/push_svr_log"
	urlProductionMainland = "http://ulog-inner-cn.dhgames.cn:8180/inner/push_log"
	urlProductionOversea  = "http://ulog-inner-us.dhgames.cn:8180/inner/push_log"
)

var logUrl = urlTestInternal

func (l *biLog) push() error {
	dataLog, err := json.Marshal(l)
	if err != nil {
		return err
	}

	bodyString := fmt.Sprintf("data=%s", url.QueryEscape(string(dataLog)))
	return postLog(bodyString)
	// err = postLog(bodyString)
	// if err == nil {
	// 	return nil
	// }

	// if !cache {
	// 	return err
	// }

	// return save(bodyString)
}

func postLog(content string) error {
	res, err := http.Post(logUrl, "application/x-www-form-urlencoded", strings.NewReader(content))
	defer func() {
		if err == nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.Errorf("log post error,status_code=%v", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	if result["error_msg"] != "ok" {
		return errors.Errorf("请求ulog失败，对方访问err_msg不为ok，error=%v", result)
	}

	return nil
}
