package tutum

import "encoding/json"

func ListImages() (ImageListResponse, error) {
	url := "image/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ImageListResponse
	var finalResponse ImageListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse ImageListResponse
			data, err := TutumCall(response.Meta.Next[8:], request, body)
			if err != nil {
				return nextResponse, err
			}
			err = json.Unmarshal(data, &nextResponse)
			if err != nil {
				return nextResponse, err
			}
			finalResponse.Objects = append(finalResponse.Objects, nextResponse.Objects...)
			response = nextResponse

		} else {
			break Loop
		}
	}

	return finalResponse, nil
}

func GetImage(name string) (Image, error) {

	url := ""
	if string(name[0]) == "/" {
		url = name[8:]
	} else {
		url = "image/" + name + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Image

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

func CreateImage(createRequest ImageCreateRequest) (Image, error) {

	url := "image/"
	request := "POST"
	var response Image

	newImage, err := json.Marshal(createRequest)
	if err != nil {
		return response, err
	}

	data, err := TutumCall(url, request, newImage)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*func (self *Image) Update(createRequest ImageCreateRequest) error {

	//url := "image/" + self.Name + "/"
	//request := "PATCH"

	updatedImage, err := json.Marshal(createRequest)
	if err != nil {
		return err
	}

	log.Println(createRequest)
	log.Println(string(updatedImage))
	_, err = TutumCall(url, request, updatedImage)
	if err != nil {
		return err
	}

	return nil
}*/

func (self *Image) Remove() error {
	url := "image/" + self.Name + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
