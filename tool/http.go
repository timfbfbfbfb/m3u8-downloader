package tool

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

// NOTE: Firefox on Windows
const DUMMY_USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0"

func Get(url string) (io.ReadCloser, error) {
	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(60) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", DUMMY_USER_AGENT)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
