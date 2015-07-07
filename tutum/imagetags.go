package tutum

import "encoding/json"

func GetImageTag(name string, tag string) (ImageTags, error) {
	url := ""
	if string(name[0]) == "/" {
		url = name[8:]
	} else {
		url = "image/" + name + "/tag/" + tag + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ImageTags

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

func (self *ImageTags) Build() (ImageTags, error) {

	url := "image/" + self.Image[14:len(self.Image)-1] + "/tag/" + self.Name + "/build/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response ImageTags

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
