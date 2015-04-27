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

var fake_uuid_stack = "09cbcf8d-a727-40d9-b420-c8e18b7fa55b"

func Test_ListStacks(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("liststacks.json")
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response StackListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"

	test_response, err := ListStacks()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_GetStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := GetStack(fake_uuid_stack)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_StartStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack + "/start/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := StartStack(fake_uuid_stack)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_CreateStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := CreateStack([]byte(`{"image": "tutum/hello-world", "name": "my-new-app", "target_num_containers": 2}`))
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_StopStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack + "/stop/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := StopStack(fake_uuid_stack)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_UpdateStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack + "/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := UpdateStack(fake_uuid_stack, []byte(`{"image": "tutum/hello-world", "name": "my-new-app", "target_num_containers": 2}`))
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_RedeployStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack + "/redeploy/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := RedeployStack(fake_uuid_stack)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_TerminateStack(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("stack.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/stack/" + fake_uuid_stack + "/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Stack
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := TerminateStack(fake_uuid_stack)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}
