package tutum

import "encoding/json"

func ListAZ() (AZListResponse, error) {
	url := "az/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response AZListResponse
	var finalResponse AZListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse AZListResponse
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

func GetAZ(az string) (AZ, error) {

	url := "az/" + az + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response AZ

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
