package validip

import (
	"fmt"
	"time"
)

// 获取IP地址列表
func ValidIp(domain string) (r string) {
	result := []string{}
	if ips, e := ipAddress(domain); e == nil {
		result = append(result, ips...)
	}
	if ips, e := secureDns(domain); e == nil {
		result = append(result, ips...)
	}
	var minAvg = 1 * time.Hour
	fmt.Printf("尝试 %s\n", domain)
	for _, ipAddr := range result {
		avgTime, e := PingCmd(ipAddr)
		if e != nil {
			fmt.Printf(" %s: failure\n", ipAddr)
			continue
		}
		fmt.Printf(" %s: success\n", ipAddr)
		if avgTime < minAvg {
			r, minAvg = ipAddr, avgTime
		}
	}
	return
}
