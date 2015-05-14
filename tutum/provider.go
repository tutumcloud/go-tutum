package tutum

import "encoding/json"

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

	url := ""
	if string(name[0]) == "/" {
		url = name[8:]
	} else {
		url = "provider/" + name + "/"
	}

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
