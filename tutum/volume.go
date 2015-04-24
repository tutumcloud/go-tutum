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

func ListVolumes() ([]Volume, error) {
	url := "volume/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Objects, nil
}

func GetVolume(uuid string) (Volume, error) {

	url := "volume/" + uuid + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Volume

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}
