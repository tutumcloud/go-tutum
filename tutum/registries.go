package tutum

import (
	"encoding/json"
	"log"
)

func ListRegistries() (RegistryListResponse, error) {
	url := "registry/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response RegistryListResponse
	var finalResponse RegistryListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Println("hello")
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse RegistryListResponse
			data, err := TutumCall(response.Meta.Next[8:], request, body)
			if err != nil {
				return nextResponse, err
			}
			err = json.Unmarshal(data, &nextResponse)
			if err != nil {
				return nextResponse, err
			}
			finalResponse.Objects = append(finalResponse.Objects, nextResponse.Objects...)
			response = nextResponse

		} else {
			break Loop
		}
	}

	return finalResponse, nil
}

func GetRegistry(registry string) (Registry, error) {

	url := "registry/" + registry + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Registry

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
