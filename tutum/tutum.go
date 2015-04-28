package tutum

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
	Start() string
	Stop() string
	Redeploy()
	Terminate()
	Deploy()
	Update([]byte)
}

func init() {
	LoadAuth()
}
