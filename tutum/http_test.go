package tutum

import (
	"os"
	"testing"
)

func Test_SetBaseUrl(t *testing.T) {

	url := ""

	os.Setenv("TUTUM_REST_HOST", "https://dashboard.tutum.co")
	url = SetBaseUrl()
	if url != "https://dashboard.tutum.co/api/v1/" {
		t.Fatal("Wrong url set")
	}
	os.Setenv("TUTUM_REST_HOST", "")
	os.Setenv("TUTUM_BASE_URL", "https://dashboard.tutum.co/api/v1/")
	url = SetBaseUrl()
	if url != "https://dashboard.tutum.co/api/v1/" {
		t.Fatal("Wrong url set")
	}
	os.Setenv("TUTUM_BASE_URL", "")
}
