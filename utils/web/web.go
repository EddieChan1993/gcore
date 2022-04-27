package web

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
)

// LocalIP 获取本地ip
func LocalIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return nil, nil
}

// LocalIPString 获取本地ip string
func LocalIPString() string {
	ip, err := LocalIP()
	if err != nil {
		return ""
	}
	if ip == nil {
		return ""
	}
	return ip.String()
}

//CopyReader 查看body内容
func CopyReader(b io.ReadCloser) (bs []byte, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return buf.Bytes(), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
