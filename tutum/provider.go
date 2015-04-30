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

/*
func ListProviders
Returns : Array of Provider objects
*/
func ListProviders() (ProviderListResponse, error) {

	url := "provider/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ProviderListResponse

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
func GetProvider
Argument : name of the provider
Returns : Provider JSON object
*/
func GetProvider(name string) (Provider, error) {

	url := "provider/" + name + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Provider

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
