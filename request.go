package calorigin

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) doRequest(req *http.Request) ([]byte, error) {
	httpClient := &http.Client{}
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
