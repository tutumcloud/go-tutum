package tutum

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	fake_api_url          = ""
	fake_uuid_container   = "6ae5e8f7-e1ae-4825-b208-c459b6e333e9"
	containerListResponse = "[{ OFF  OFF [{ 80 0 tcp}]  0 Thu, 23 Apr 2015 10:37:13 +0000   0  tutum/hello-world:latest /api/v1/image/tutum/hello-world/tag/latest/ my-new-app-1 my-new-app-1.maximeheckel.cont.tutum.io /api/v1/container/6ae5e8f7-e1ae-4825-b208-c459b6e333e9/ /run.sh Thu, 23 Apr 2015 10:37:13 +0000 Running   6ae5e8f7-e1ae-4825-b208-c459b6e333e9}]"
	getContainerResponse  = "{ OFF  OFF [{ 80 0 tcp}]  0 Thu, 23 Apr 2015 10:37:13 +0000   0  tutum/hello-world:latest /api/v1/image/tutum/hello-world/tag/latest/ my-new-app-1 my-new-app-1.maximeheckel.cont.tutum.io /api/v1/container/6ae5e8f7-e1ae-4825-b208-c459b6e333e9/ /run.sh Thu, 23 Apr 2015 10:37:13 +0000 Running   6ae5e8f7-e1ae-4825-b208-c459b6e333e9}"
)

func Test_ListContainers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, containerListResponse)
	}))

	defer server.Close()
	url := server.URL + "/container/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Get: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	testResp := string(body)
	fmt.Println(testResp)

}

func Test_getContainer(t *testing.T) {
	/*server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, getContainerResponse)
	}))

	defer server.Close()
	url := server.URL + "/container/" + fake_uuid + "/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Get: %v", err)
	}*/

}
