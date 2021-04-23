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

func AddNetOutput(URL string, level Level, format ...Format) {
	output := Output{NetOutput{URL: URL}, level, JSON, ""}
	if len(format) > 0 {
		output.Format = format[0]
	}
	outputs = append(outputs, output)
}

func (o NetOutput) Write(p []byte) (n int, err error) {
	n = len(p)

	if o.URL == "" {
		err = fmt.Errorf("URL not set")
		return
	}

	req, err := http.NewRequest("POST", o.URL, bytes.NewReader(p))
	if err != nil {
		err = fmt.Errorf("failed to create request with log message to %s: %w", o.URL, err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send log message to %s: %w", o.URL, err)
		return
	}
	defer res.Body.Close()
	return
}
