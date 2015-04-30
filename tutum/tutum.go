package tutum

import (
	"fmt"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

var version string = "0.9.8"

var (
	User    string
	ApiKey  string
	BaseUrl = "https://dashboard.tutum.co/api/v1/"
)

type config map[string]Auth

type Auth struct {
	User   string
	Apikey string
}

type TutumObject interface {
	Start() error
	Stop() error
	Redeploy() error
	Terminate() error
	Deploy() error
	Upgrade() error
	Update([]byte) error
}

func init() {
	LoadAuth()
}

func LoadAuth() error {
	if User != "" && ApiKey != "" {
		// Configuration already loaded
		return nil
	}

	// Process ~/.tutum configuration file first
	if usr, err := user.Current(); err == nil {
		var conf config
		confFilePath := usr.HomeDir + "/.tutum"
		if _, err := os.Stat(confFilePath); !os.IsNotExist(err) {
			if _, err := toml.DecodeFile(confFilePath, &conf); err == nil {
				if conf["auth"].User != "" && conf["auth"].Apikey != "" {
					User = conf["auth"].User
					ApiKey = conf["auth"].Apikey
					return nil
				}
			} else {
				return fmt.Errorf("Malformed Tutum configuration file found at %s: %s", confFilePath, err)
			}
		}
	}

	// Load environment variables as an alternative option
	if os.Getenv("TUTUM_USER") != "" && os.Getenv("TUTUM_APIKEY") != "" {
		User = os.Getenv("TUTUM_USER")
		ApiKey = os.Getenv("TUTUM_APIKEY")
		return nil
	}

	return fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
}

func IsAuthenticated() bool {
	return (User != "" && ApiKey != "")
}
