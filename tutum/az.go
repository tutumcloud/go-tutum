package tutum

import "encoding/json"

func ListAZ() (AZListResponse, error) {
	url := "az/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response AZListResponse

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
