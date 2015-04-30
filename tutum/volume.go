package tutum

import "encoding/json"

type VolumeListResponse struct {
	Objects []Volume `json:"objects"`
}

type Volume struct {
	Containers   []string `json:"containers"`
	Node         string   `json:"node"`
	Resource_uri string   `json:"resource_uri"`
	State        string   `json:"state"`
	Uuid         string   `json:"uuid"`
	Volume_group string   `json:"volume_group"`
}

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

	url := "volume/" + uuid + "/"
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
