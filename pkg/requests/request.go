package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest(respUrl string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, respUrl, nil)
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
