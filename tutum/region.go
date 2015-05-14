package tutum

import "encoding/json"

/*
func ListRegions
Returns : Array of Region objects
*/
func ListRegions() (RegionListResponse, error) {

	url := "region/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response RegionListResponse

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

/*
func GetRegion
Argument : provider name and location name
Returns : Region JSON object
*/
func GetRegion(id string) (Region, error) {

	url := ""
	if string(id[0]) == "/" {
		url = id[8:]
	} else {
		url = "region/" + id + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Region

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
