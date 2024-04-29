package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendRequest(respUrl string, apiKey string, q url.Values) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, respUrl, nil)
	req.Header.Add("Authorization", apiKey)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return body, nil
}
