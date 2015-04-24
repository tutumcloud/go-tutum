package tutum

import "encoding/json"

type ProviderListResponse struct {
	Objects []Provider `json:"objects"`
}

type Provider struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Regions      []string `json:"regions"`
	Resource_uri string   `json:"resource_uri"`
}

func ListProviders() ([]Provider, error) {

	url := "provider/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ProviderListResponse

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

func GetProvider(name string) (Provider, error) {

	url := "provider/" + name + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Provider

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
