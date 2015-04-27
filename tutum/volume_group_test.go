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

var fake_uuid_volumegroup = "1863e34d-6a7d-4945-aefc-8f27a4ab1a9e"

func Test_ListVolumeGroups(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("listvolumegroups.json")
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/volumegroup/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response VolumeGroupListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"

	test_response, err := ListVolumeGroups()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_GetVolumeGroup(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("volumegroup.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/volumegroup/" + fake_uuid_volumegroup

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response VolumeGroup
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/v1/"
	test_response, err := GetVolumeGroup(fake_uuid_volumegroup)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}
