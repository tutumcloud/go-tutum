package tutum

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
)

var User string
var Apikey string

//Used to unpack the config file.
type Auth struct {
	User   string
	Apikey string
}
type config map[string]Auth

func TutumCall(url string, requestType string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(requestType, BaseUrl()+url, nil)
	req.Header.Add("Authorization", AuthForUser())
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed API call " + response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AuthForUser() string {
	//Check the init file first.  See if it exists?
	if User == "" || Apikey == "" {
		//Get the current user, so we can access home directory.
		usr, err := user.Current()
		if err != nil {
			//What to do if the current user cant be found?

		}

		var conf config
		if _, err := toml.DecodeFile(usr.HomeDir+"/.tutum", &conf); err != nil {
			// handle error
		}

		if User == "" {
			if conf["auth"].User != "" {
				User = conf["auth"].User
			}
		}
		if Apikey == "" {
			if conf["auth"].Apikey != "" {
				Apikey = conf["auth"].Apikey
			}
		}
	}

	if User == "" {
		if os.Getenv("TUTUM_USER") != "" {
			User = os.Getenv("TUTUM_USER")
		}
	}

	if Apikey == "" {
		if os.Getenv("TUTUM_APIKEY") != "" {
			Apikey = os.Getenv("TUTUM_APIKEY")
		}
	}
	return "ApiKey " + User + ":" + Apikey
}

func BaseUrl() string {
	var baseUrl string
	if os.Getenv("TUTUM_BASE_URL") != "" {
		baseUrl = os.Getenv("TUTUM_BASE_URL")
	} else {
		baseUrl = "https://app.tutum.co/api/v1/"
	}
	return baseUrl
}
