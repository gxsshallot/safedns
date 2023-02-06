package validip

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ipAddress(domain string) (ips []string, err error) {
	ipAddressAddr := fmt.Sprintf("https://www.ipaddress.com/site/%s", domain)
	res, err := http.Get(ipAddressAddr)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error: %d", res.StatusCode)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	doc.Find("#map-1+table tr:last-child td>ul>li").Each(func(i int, s *goquery.Selection) {
		ips = append(ips, s.Text())
	})
	return
}
