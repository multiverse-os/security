package shodan

import (
	"fmt"
	"time"
	"strconv"

	"gopkg.in/headzoo/surf.v1"
	browser "github.com/headzoo/surf/browser"
	"github.com/PuerkitoBio/goquery"

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
	return &Shodan{
		IPs: []string{},
		PageLimit: DefaultPageLimit,
		Page: surf.NewBrowser(),
	}
}

func Login(username, password string) *Shodan {
	shodan := New()
	err := shodan.Page.Open("https://account.shodan.io/login")
	fm, _ := shodan.Page.Form("form.login-form")
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
		err = self.Page.Open("https://www.shodan.io/search?query="+value+"&page="+strconv.Itoa(i))
		if err != nil {
			fmt.Println("[Error] Failed to connect to shodan, try again shortly...")
			self.Page.Find("div.ip").Each(func(i int, element *goquery.Selection) {
				fmt.Print(element.Text())
				ips = append(ips, element.Text())
			})
		}else{
			return ips, err
		}
	}
	return ips, nil
}

