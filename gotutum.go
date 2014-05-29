package gotutum

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var User string
var Apikey string

func AuthForUser() string {

	return "ApiKey " + User + ":" + Apikey
}

type Objects struct {
	Status string
}

func ListContainers() interface{} {

	url := "https://app.tutum.co/api/v1/container/"
	request := "GET"

	f := TutumCall(url, request)
	m := f.(map[string]interface{})

	//Relies on the array of items being named objects.
	for k, v := range m {
		if k == "objects" {
			return v
		}
	}
	return nil
}

func TutumCall(url string, requestType string) interface{} {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", AuthForUser())
	req.Header.Add("Accept", "application/json")

	log.Println(url)
	// log.Println(req.Header)
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	// log.Println(string(data))

	//Response has a tag named objects.  Need to unmarshal objects, then work with objects.

	var f interface{}
	err = json.Unmarshal(data, &f)
}
