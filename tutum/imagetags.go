package tutum

import (
	"encoding/json"
	"log"
)

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

func GetImageBuildSetting(name string, tag string) (BuildSettings, error) {

	url := ""
	if string(name[0]) == "/" {
		url = name[8:]
	} else {
		url = "image/" + name + "/buildsetting/" + tag + "/"
	}

	log.Println(url)

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response BuildSettings

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

func (self *BuildSettings) Build() (BuildSettings, error) {
	url := "image/" + self.Image[14:len(self.Image)-1] + "/buildsetting/" + self.Tag + "/build/"
	request := "POST"

	body := []byte(`{}`)
	var response BuildSettings

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
