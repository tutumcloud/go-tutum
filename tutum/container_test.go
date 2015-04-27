package tutum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var (
	fake_uuid_container = "dcbe16b4-21a1-474b-a814-131a3626b1de"
)

func MockupResponse(response_file string) (string, error) {

	file, e := ioutil.ReadFile("json_test_output/" + response_file)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	fake_response := string(file)

	return fake_response, nil
}

func Test_ListContainers(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("listcontainers.json")

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response CListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	test_response, err := GetContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_StartContainer(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("container.json")

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

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

	BaseUrl = server.URL + "/api/v1/"
	if err != nil {
		t.Fatal(err)
	}

	var response Container
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	test_response, err := TerminateContainer(fake_uuid_container)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}
