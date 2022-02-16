package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	Infof("test info")
	Warnf("test info")
	Errorf("test info")
	panic("test panic")
}
