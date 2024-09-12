package crtsh

import (
	"crypto/x509"
	"encoding/pem"

	"github.com/cockroachdb/errors"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func parseHTMLElement(first *html.Node) string {
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
		return nil, ErrorParsePEM
	}

	data, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Join(ErrorParseCertificate, err)
	}

	return data, nil
}
