package tutum

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
)

var version string = "0.9.8"

var (
	User       string
	ApiKey     string
	BaseUrl    = "https://dashboard.tutum.co/api/v1/"
	StreamUrl  = "wss://stream.tutum.co:443/v1/"
	AuthHeader string
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
		return nil
	} else {
		if os.Getenv("TUTUM_AUTH") != "" {
			AuthHeader = os.Getenv("TUTUM_AUTH")
		}
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
	return ((User != "" && ApiKey != "") || os.Getenv("TUTUM_AUTH") != "")
}

func FetchByResourceUri(id string) interface{} {
	words := strings.Split(id, "/")
	switch words[3] {
	case "action":
		action, err := GetAction(id)

		if err != nil {
			return err
		}

		return action
		break
	case "nodecluster":
		nodecluster, err := GetNodeCluster(id)

		if err != nil {
			return err
		}

		return nodecluster
		break
	case "provider":
		provider, err := GetProvider(id)

		if err != nil {
			return err
		}

		return provider
		break
	case "region":
		region, err := GetRegion(id)

		if err != nil {
			return err
		}

		return region
		break
	case "service":
		service, err := GetService(id)

		if err != nil {
			return err
		}

		return service
		break
	case "stack":
		stack, err := GetStack(id)

		if err != nil {
			return err
		}

		return stack
		break
	case "volume":
		volume, err := GetVolume(id)

		if err != nil {
			return err
		}

		return volume
		break
	case "volumegroup":
		volumegroup, err := GetVolumeGroup(id)

		if err != nil {
			return err
		}

		return volumegroup
		break
	case "node":
		node, err := GetNode(id)

		if err != nil {
			return err
		}

		return node
		break
	case "container":
		container, err := GetContainer(id)

		if err != nil {
			return err
		}

		return container
		break
	}
	return 0
}
