package tutum

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func TutumCall(url string, requestType string, requestBody []byte) ([]byte, error) {

	if os.Getenv("TUTUM_BASE_URL") != "" {
		BaseUrl = os.Getenv("TUTUM_BASE_URL")
	}

	if !IsAuthenticated() {
		return nil, fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
	}
	client := &http.Client{}
	req, err := http.NewRequest(requestType, BaseUrl+url, bytes.NewBuffer(requestBody))
	if ApiKey != "" {
		AuthHeader = fmt.Sprintf("ApiKey %s:%s", User, ApiKey)
	}
	req.Header.Add("Authorization", AuthHeader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-tutum "+version)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 300 {
		return nil, fmt.Errorf("Failed API call: %s ", response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(data))
	return data, nil
}
