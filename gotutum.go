package tutum

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var User string
var Apikey string

func AuthForUser() string {
	//Still need to check init file.  Look at python to see how to do it.
	//First init file, then ENV VAR, then none.
	//Check if the User and Apikey are set.  If not grab from the environment variables.
	if User == "" {
		if os.Getenv("TUTUM_USER") != "" {
			User = os.Getenv("TUTUM_USER")
		}
	}
	if Apikey == "" {
		if os.Getenv("TUTUM_APIKEY") != "" {
			Apikey = os.Getenv("TUTUM_APIKEY")
		}
	}
	return "ApiKey " + User + ":" + Apikey
}

func BaseUrl() string {
	var baseUrl string
	if os.Getenv("TUTUM_BASE_URL") != "" {
		url = os.Getenv("TUTUM_BASE_URL")
	} else {
		baseUrl = "https://app.tutum.co/api/v1/"
	}
	return baseUrl
}

func TutumCall(url string, requestType string) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(requestType, BaseUrl()+url, nil)
	req.Header.Add("Authorization", AuthForUser())
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	data, err := ioutil.ReadAll(response.Body)

	return data
}
