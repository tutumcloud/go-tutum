package tutum

import "encoding/json"

type WebhookListResponse struct {
	Objects []Webhook `json:"objects"`
}

type Webhook struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
}

func ListWebhooks(uuid string) ([]Webhook, error) {
	url := "service/" + uuid + "/webhook/handler/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response WebhookListResponse

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

func GetWebhook(uuid string, webhook_uuid string) (Webhook, error) {
	url := "service/" + uuid + "/webhook/handler/" + webhook_uuid + "/"
	request := "GET"
	body := []byte(`{}`)
	var response Webhook

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

func CreateWebhook(uuid string, newWebhook []byte) ([]Webhook, error) {

	url := "service/" + uuid + "/webhook/handler/"
	request := "POST"

	var response []Webhook

	data, err := TutumCall(url, request, newWebhook)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil

}

func DeleteWebhook(uuid string, webhook_uuid string) error {
	url := "service/" + uuid + "/webhook/handler/" + webhook_uuid + "/"
	request := "DELETE"
	body := []byte(`{}`)
	var response Webhook

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return nil
}

func CallWebhook(uuid string, webhook_uuid string) (Webhook, error) {
	url := "service/" + uuid + "/webhook/handler/" + webhook_uuid + "/call/"
	request := "POST"
	body := []byte(`{}`)
	var response Webhook

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