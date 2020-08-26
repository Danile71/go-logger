package logger

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

func send(data []byte) {
	if api == "" {
		return
	}

	req, err := http.NewRequest("POST", api, bytes.NewReader(data))
	if err != nil {
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
}
