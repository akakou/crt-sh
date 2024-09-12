package main

import (
	"crypto/x509"
	"fmt"
	"net/url"
	"strings"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

const BASE_URL = "https://crt.sh/atom"

type CrtSH struct {
	*gofeed.Parser
}

func Fetch(domain, exclude string) ([]*x509.Certificate, error) {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("q", domain)
	query.Add("exclude", domain)
	u.RawQuery = query.Encode()

	feed, err := gofeed.NewParser().ParseURL(u.String())
	if err != nil {
		return nil, err
	}

	var certs []*x509.Certificate
	for _, item := range feed.Items {
		desc := strings.NewReader(item.Description)
		node, _ := html.Parse(desc)

		first := node.LastChild.LastChild.LastChild.FirstChild

		s := parseHTML(first)
		c, err := ParseCertificate(s)
		if err != nil {
			return nil, err
		}

		certs = append(certs, c)
	}

	return certs, err
}

func main() {
	data, err := Fetch("test.ochano.co", "expired")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", data)
}
