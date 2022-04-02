package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	ResetToDevelopment()
	Infof("test Infof")
	Warnf("test Warnf")
	Errorf("test Errorf")
	//panic("test panic")
}
