package container

import (
	"log"
	"tutum"
)

func List() interface{} {

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
