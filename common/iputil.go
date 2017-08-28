package common

import (
	"net"
	"errors"
)

/*
 * 获取本机的ip地址(ipv4)
 */
func Localipv4() (string, error) {
	addrs, _ := net.InterfaceAddrs()
	
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && len(ipnet.IP.To4())==4 {
			
			return ipnet.IP.String(), nil
		}
	}
	
	return "", errors.New("invalid ip(ipv4) address")
}

func Remoteipv4(host string) (string, error) {
	ips, _ := net.LookupIP(host)
	for _, ip := range ips {
		if !ip.IsLoopback() && len(ip.To4())==4 {
			return ip.String(), nil
		}
	}
	
	return "", errors.New("invalid ip(ipv4) address")
}