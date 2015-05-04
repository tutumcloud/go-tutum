package tutum

import "encoding/json"

type VolumeGroupListResponse struct {
	Objects []VolumeGroup `json:"objects"`
}

type VolumeGroup struct {
	Name         string   `json:"name"`
	Resource_uri string   `json:"resource_uri"`
	Services     []string `json:"services"`
	State        string   `json:"state"`
	Uuid         string   `json:"uuid"`
	Volume       []string `json:"volume"`
}

/*
func ListVolumeGroups
Returns : Array of VolumeGroup objects
*/
func ListVolumeGroups() (VolumeGroupListResponse, error) {

	url := "volumegroup/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeGroupListResponse

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
func GetVolumeGroup
Argument : uuid
Returns : VolumeGroup JSON object
*/
func GetVolumeGroup(uuid string) (VolumeGroup, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "volumegroup/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeGroup

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
