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

	m := f.(map[string]interface{})
	// log.Println(m)
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		log.Println(k, "is string", vv)
	// 	case int:
	// 		log.Println(k, "is int", vv)
	// 	case []interface{}:
	// 		log.Println(k, "is an array:")
	// 		//Loops through containers.  We want an object. made out of u.
	// 		for i, u := range vv {
	// 			log.Println(i)
	// 			log.Println(u)
	// 			log.Println("Gets here")
	// 		}
	// 	default:
	// 		log.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }

	// for k, v := range m {
	// 	if v.(type) == []interface{} {
	// 		log.Println(v)
	// 	}
	// }
	for k, v := range m {
		// log.Println()
		// log.Println(k)
		// log.Println(v)
		if k == "objects" {
			return v
		}
	}
	return nil
}
