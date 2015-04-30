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

/*
func ListWebhooks
Returns : Array of Webhook objects
*/
func (self *Service) ListWebhooks() (WebhookListResponse, error) {
	url := "service/" + self.Uuid + "/webhook/handler/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response WebhookListResponse

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
func GetWebhook
Argument : service uuid and webhook uuid
Returns : Webhook JSON object
*/
func (self *Service) GetWebhook(webhook_uuid string) (Webhook, error) {

	url := ""
	if string(webhook_uuid[0]) == "/" {
		url = webhook_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/webhook/handler/" + webhook_uuid + "/"
	}

	request := "GET"
	body := []byte(`{}`)
	var response Webhook

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
func CreateWebhook
Argument : service uuid and webhook JSON object
Returns : Array of Webhook objects
*/
func (self *Service) CreateWebhook(requestBody string) ([]Webhook, error) {

	url := "service/" + self.Uuid + "/webhook/handler/"
	request := "POST"

	newWebhook := []byte(requestBody)

	var response []Webhook

	data, err := TutumCall(url, request, newWebhook)
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
func DeleteWebhook
Argument : service uuid and webhook uuid
*/
func (self *Service) DeleteWebhook(webhook_uuid string) error {
	url := ""
	if string(webhook_uuid[0]) == "/" {
		url = webhook_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/webhook/handler/" + webhook_uuid + "/"
	}

	request := "DELETE"
	body := []byte(`{}`)
	var response Webhook

	data, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return nil
}

/*
func CallWebhook
Argument : service uuid and webhook uuid
Returns : Webhook JSON object
*/
func (self *Service) CallWebhook(webhook_uuid string) (Webhook, error) {
	url := ""
	if string(webhook_uuid[0]) == "/" {
		url = webhook_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/webhook/handler/" + webhook_uuid + "/"
	}

	request := "POST"
	body := []byte(`{}`)
	var response Webhook

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
