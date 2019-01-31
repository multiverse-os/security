package shodan

import (
	"gopkg.in/headzoo/surf.v1"
	"github.com/PuerkitoBio/goquery"

	"fmt"
)



func Search(value string) (ips []string) {
	bow := surf.NewBrowser()

	// If you are not loggedin, only 2 pages can be requested
	for i := 1; i <= 2; i++ {
		err := bow.Open("https://www.shodan.io/search?query="+value+"&page="+strconv.Itoa(i))
		if err != nil {
			fmt.Println("[Error] Failed to connect to shodan, try again shortly...")

		bow.Find("div.ip").Each(func(i int, element *goquery.Selection) {
			fmt.Print(element.Text())
			ips = append(ips, element.Text())
		})
	}
	return ips
}
