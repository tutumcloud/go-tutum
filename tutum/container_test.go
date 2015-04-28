package tutum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_ListContainers(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("listcontainers.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response CListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"

	test_response, err := ListContainers()
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_GetContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/" + fake_uuid_container

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := GetContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

/*func Test_StartContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/" + fake_uuid_container + "/start"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := StartContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_StopContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/" + fake_uuid_container + "/stop/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := StopContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_RedeployContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/" + fake_uuid_container + "/redeploy/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := RedeployContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_TerminateContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/" + fake_uuid_container

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"

	test_response, err := TerminateContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}*/
