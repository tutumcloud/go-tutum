package tutum

import "encoding/json"

/*
func ListVolumes
Returns : Array of Volume objects
*/
func ListVolumes() (VolumeListResponse, error) {
	url := "volume/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeListResponse

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
func GetVolume
Argument : uuid
Returns : Volume JSON object
*/
func GetVolume(uuid string) (Volume, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "volume/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Volume

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
