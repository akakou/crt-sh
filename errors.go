package crtsh

import "github.com/pkg/errors"

var (
	ErrorParseURL         = errors.Errorf("failed to parse ParseURL")
	ErrorParseCertificate = errors.Errorf("failed to parse Certificate")
	ErrorParsePEM         = errors.Errorf("failed to parse PEM")
	ErrorParseHTML        = errors.Errorf("failed to parse HTML")
	ErrorFetchRSS         = errors.Errorf("failed to fetch RSS feed")
)
