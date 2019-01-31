package shodan

import (
	"gopkg.in/headzoo/surf.v1"
	"github.com/PuerkitoBio/goquery"

	"fmt"
)



func Search(value string) (ips []string) {
	bow := surf.NewBrowser()
	err := bow.Open("https://www.shodan.io/search?query="+value)
	if err != nil {
		panic(err)
	}
	bow.Find("div.ip").Each(func(i int, element *goquery.Selection) {
		fmt.Print(element.Text())
		ips = append(ips, element.Text())
	})
	return ips
}
