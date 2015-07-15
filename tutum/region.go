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
	var finalResponse RegionListResponse

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
			var nextResponse RegionListResponse
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
