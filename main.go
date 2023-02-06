package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"safedns/validip"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("使用格式: safedns [TXT_PATH]\n")
		os.Exit(1)
	}
	txtPath := os.Args[1]
	content, err := ioutil.ReadFile(txtPath)
	if err != nil {
		fmt.Printf("读取文件错误: %s", err)
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	result := []string{}
	for _, domain := range lines {
		ipAddr := validip.ValidIp(domain)
		if len(ipAddr) > 0 {
			result = append(result, fmt.Sprintf("%s %s", ipAddr, domain))
		}
	}
	fmt.Println()
	fmt.Printf("以下为hosts内容:\n")
	fmt.Printf("%s\n", strings.Join(result, "\n"))
}
