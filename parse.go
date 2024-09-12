package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func parseHTML(first *html.Node) string {
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

func ParseCertificate(src string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(src))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	data, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return data, nil
}
