package tutum

import "encoding/json"

type TriggerListResponse struct {
	Objects []Trigger `json:"objects"`
}

type Trigger struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	Operation    string `json:"operation"`
	Resource_uri string `json:"resource_uri"`
}

/*
func ListTriggers
Returns : Array of Trigger objects
*/
func (self *Service) ListTriggers() (TriggerListResponse, error) {
	url := "service/" + self.Uuid + "/trigger/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response TriggerListResponse

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
func GetTrigger
Argument : service uuid and Trigger uuid
Returns : Trigger JSON object
*/
func (self *Service) GetTrigger(trigger_uuid string) (Trigger, error) {

	url := ""
	if string(trigger_uuid[0]) == "/" {
		url = trigger_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/trigger/" + trigger_uuid + "/"
	}

	request := "GET"
	body := []byte(`{}`)
	var response Trigger

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
func CreateTrigger
Argument : service uuid and Trigger JSON object
Returns : Array of Trigger objects
*/
func (self *Service) CreateTrigger(requestBody string) ([]Trigger, error) {

	url := "service/" + self.Uuid + "/trigger/"
	request := "POST"

	newTrigger := []byte(requestBody)

	var response []Trigger

	data, err := TutumCall(url, request, newTrigger)
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
func DeleteTrigger
Argument : service uuid and Trigger uuid
*/
func (self *Service) DeleteTrigger(trigger_uuid string) error {
	url := ""
	if string(trigger_uuid[0]) == "/" {
		url = trigger_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/trigger/" + trigger_uuid + "/"
	}

	request := "DELETE"
	body := []byte(`{}`)
	var response Trigger

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
func CallTrigger
Argument : service uuid and Trigger uuid
Returns : Trigger JSON object
*/
func (self *Service) CallTrigger(trigger_uuid string) (Trigger, error) {
	url := ""
	if string(trigger_uuid[0]) == "/" {
		url = trigger_uuid[8:]
	} else {
		url = "service/" + self.Uuid + "/trigger/" + trigger_uuid + "/call/"
	}

	request := "POST"
	body := []byte(`{}`)
	var response Trigger

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
