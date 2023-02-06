package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gaoxiaosong/safedns"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("使用格式: safedns [TXT_PATH]\n")
		os.Exit(1)
	}
	var domains []string
	if len(os.Args) == 2 {
		txtPath := os.Args[1]
		content, err := ioutil.ReadFile(txtPath)
		if err != nil {
			fmt.Printf("读取文件错误: %s", err)
			os.Exit(1)
		}
		domains = strings.Split(string(content), "\n")
	} else {
		var domain string
		for {
			fmt.Printf("请输入一个域名: ")
			_, err := fmt.Scanln(&domain)
			if err != nil {
				continue
			}
			if len(domain) == 0 {
				continue
			}
			break
		}
		domains = []string{domain}
	}

	result := []string{}
	for _, domain := range domains {
		ipAddr := safedns.ValidIp(domain)
		if len(ipAddr) > 0 {
			result = append(result, fmt.Sprintf("%s %s", ipAddr, domain))
		}
	}
	fmt.Println()
	if len(result) > 0 {
		fmt.Printf("以下为hosts内容:\n")
		fmt.Printf("%s\n", strings.Join(result, "\n"))
	} else {
		fmt.Println("未找到合适的IP地址")
	}
	fmt.Println()
	fmt.Println("按回车键退出")
	fmt.Scanln()
}
