package tutum

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

func TutumCall(url string, requestType string, requestBody []byte) ([]byte, error) {
	/*

		Idea : To run the test, set a variable TUTUM_ENV = "test"
		at the beginning of each test, tell the TutumCall function
		to look for this variable
		if exist => use TestUrl
		if not => use BaseUrl

		Remove the ENV function after the tests are done
	*/

	if os.Getenv("TUTUM_BASE_URL") != "" {
		BaseUrl = os.Getenv("TUTUM_BASE_URL")
	}

	/*if os.Getenv("TUTUM_ENV") != "" {
		fmt.Println("Ok")
		BaseUrl = "http://127.0.0.1/api/v1/"
	} else {

	}*/

	if !IsAuthenticated() {
		return nil, fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
	}
	client := &http.Client{}
	req, err := http.NewRequest(requestType, BaseUrl+url, bytes.NewBuffer(requestBody))
	authHeader := fmt.Sprintf("ApiKey %s:%s", User, ApiKey)
	req.Header.Add("Authorization", authHeader)
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

func LoadAuth() error {
	if User != "" && ApiKey != "" {
		// Configuration already loaded
		return nil
	}

	// Process ~/.tutum configuration file first
	if usr, err := user.Current(); err == nil {
		var conf config
		confFilePath := usr.HomeDir + "/.tutum"
		if _, err := os.Stat(confFilePath); !os.IsNotExist(err) {
			if _, err := toml.DecodeFile(confFilePath, &conf); err == nil {
				if conf["auth"].User != "" && conf["auth"].Apikey != "" {
					User = conf["auth"].User
					ApiKey = conf["auth"].Apikey
					return nil
				}
			} else {
				return fmt.Errorf("Malformed Tutum configuration file found at %s: %s", confFilePath, err)
			}
		}
	}

	// Load environment variables as an alternative option
	if os.Getenv("TUTUM_USER") != "" && os.Getenv("TUTUM_APIKEY") != "" {
		User = os.Getenv("TUTUM_USER")
		ApiKey = os.Getenv("TUTUM_APIKEY")
		return nil
	}

	return fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
}

func IsAuthenticated() bool {
	return (User != "" && ApiKey != "")
}
