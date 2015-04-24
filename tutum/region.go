package tutum

import "encoding/json"

type RegionListResponse struct {
	Objects []Region `json:"objects"`
}

type Region struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Node_types   []string `json:"node_types"`
	Provider     string   `json:"provider"`
	Resource_uri string   `json:"resource_uri"`
}

func ListRegions() ([]Region, error) {

	url := "region/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response RegionListResponse

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

func GetRegion(provider string, name string) (Region, error) {

	url := "region/" + provider + "/" + name + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Region

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
