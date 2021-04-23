package logger

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

type NetOutput struct {
	URL string
}

func NewNetOutput(URL string) NetOutput {
	return NetOutput{URL: URL}
}

func (o NetOutput) Write(p []byte) (n int, err error) {
	if o.URL == "" {
		n = len(p)
		err = fmt.Errorf("URL not set")
		return
	}

	req, err := http.NewRequest("POST", o.URL, bytes.NewReader(p))
	if err != nil {
		n = len(p)
		err = fmt.Errorf("failed to create request with log message to %s: %w", o.URL, err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		n = len(p)
		err = fmt.Errorf("failed to send log message to %s: %w", o.URL, err)
		return
	}
	defer res.Body.Close()
	n = 0
	err = nil
	return
}
