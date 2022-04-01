package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	ResetToDevelopment()
	Infof("test info")
	Warnf("test info")
	Errorf("test info")
	//panic("test panic")
}
