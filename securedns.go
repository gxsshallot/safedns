package safedns

import (
	"context"
	"time"

	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
)

func secureDns(domain string) (ips []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c := doh.Use()
	defer c.Close()
	rsp, err := c.Query(ctx, dns.Domain(domain), dns.TypeA)
	if err != nil {
		return
	}
	for _, item := range rsp.Answer {
		ips = append(ips, item.Data)
	}
	return
}
