package main

import (
	"fmt"
	"net/url"

	"github.com/mmcdole/gofeed"
)

const BASE_URL = "https://crt.sh/atom"

type CrtSH struct {
	*gofeed.Parser
}

func Init() *CrtSH {
	return &CrtSH{
		Parser: gofeed.NewParser(),
	}
}

func Fetch(domain string) ([]string, error) {
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

	var test []string
	for _, item := range feed.Items {
		data := item.Description
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
