package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"strings"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const BASE_URL = "https://crt.sh/atom"

type CrtSH struct {
	*gofeed.Parser
}

func ParseHTML(first *html.Node) string {
	d := ""

	for c := first; c != nil; c = c.NextSibling {
		if c.DataAtom == atom.Br {
			d += "\n"
		}

		if c.DataAtom.String() == "" {
			d += c.Data
		}
	}

	return d
}

func Fetch(domain string) ([]*x509.Certificate, error) {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("q", domain)
	u.RawQuery = query.Encode()

	feed, err := gofeed.NewParser().ParseURL(u.String())
	if err != nil {
		return nil, err
	}

	var test []*x509.Certificate
	for _, item := range feed.Items {
		desc := strings.NewReader(item.Description)
		node, _ := html.Parse(desc)

		first := node.LastChild.LastChild.LastChild.FirstChild

		d := ParseHTML(first)
		block, _ := pem.Decode([]byte(d))
		if block == nil {
			return nil, fmt.Errorf("failed to decode PEM block")
		}

		data, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}

		test = append(test, data)
	}

	return test, err
}

func main() {
	data, err := Fetch("test.ochano.co")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", data)
}
