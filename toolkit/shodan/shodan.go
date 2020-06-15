package shodan

import (
	"fmt"
	"time"
	"strconv"

	browser "github.com/headzoo/surf/browser"
	surf "github.com/headzoo/surf"
	query "github.com/PuerkitoBio/goquery"
)

const DefaultPageLimit = 2

type Server struct {
	IP string
	Ports []string
	CVEs []string
}

type Shodan struct {
	IPs []string
	PageLimit int
	Page *browser.Browser
	Servers []*Server
}

func New() *Shodan {
	s := &Shodan{
		IPs: []string{},
		PageLimit: DefaultPageLimit,
		Page: surf.NewBrowser(),
	}
	s.Page.SetUserAgent("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0")
	//s.Page.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//s.Page.AddRequestHeader("Accept-Language", "en-US,en;q=0.5")
	//s.Page.AddRequestHeader("Accept-Encoding", "gzip, deflate, br")
	//s.Page.AddRequestHeader("Connection", "keep-alive")
	//s.Page.AddRequestHeader("Upgrade-Insecure-Requests", "1")
	//s.Page.AddRequestHeader("TE", "Trailers")
	//s.Page.Open("https://shodan.io")
	return s
}

func Login(username, password string) *Shodan {
	fmt.Println("Attempting to login,..")
	shodan := New()
	err := shodan.Page.Open("https://account.shodan.io/login")
	fm, _ := shodan.Page.Form("form")

	fm.Input("username", username)
	fm.Input("password", password)
	if err != nil && fm.Submit() != nil {
		fmt.Println("[Error] Failed to load login page, will try again shortly...")
		time.Sleep(5)
		Login(username, password)
	}
	return shodan
}

func (self *Shodan) Search(value string) (ips []string, err error) {
	// If you are not loggedin, only 2 pages can be requested
	for i := 1; i <= self.PageLimit; i++ {
		fmt.Println("Loading page: [", i, "] searching for value [", value ,"]")
		err := self.Page.Open("https://www.shodan.io/search?query="+value+"&page="+strconv.Itoa(i))
		//fmt.Println("Page loaded?", self.Page.Body())

		fmt.Println("Title of page:", self.Page.Title())


		fmt.Println("[i]:", strconv.Itoa(i))

		if err == nil {
			self.Page.Find("div.ip").Each(func(i int, element *query.Selection) {
				fmt.Print(element.Text())
				ips = append(ips, element.Text())
			})
		}else{
			fmt.Println("[Error] Failed to make search request against shodan, waiting to try again:", err)
		}
	}
	return ips, nil
}

