package tutum

import "encoding/json"

type Token struct {
	Token string `json:"token"`
}

/*
	func CreateToken
	Returns : Token JSON object
*/
func CreateToken() (Token, error) {
	url := "token/"
	request := "POST"
	body := []byte(`{}`)
	var response Token

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
