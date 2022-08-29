package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type OptimizelyClient struct {
	Address string
	Token   string
}

func (c OptimizelyClient) sendHttpRequest(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.Address, url), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	if body != nil {
		req.Header.Set("Content-type", "application/json")
	}

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[INFO] Request:\n%s\nend", string(reqDump))

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP status %d\n\n%s", resp.StatusCode, respBody)
	}

	return respBody, nil
}
