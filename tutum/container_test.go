package tutum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	fake_api_url          = ""
	fake_uuid_container   = "6ae5e8f7-e1ae-4825-b208-c459b6e333e9"
	containerListResponse = `{
    "meta": {
        "limit": 25,
        "next": null,
        "offset": 0,
        "previous": null,
        "total_count": 2
    },
    "objects": [
        {
            "autodestroy": "OFF",
            "autorestart": "OFF",
            "container_ports": [
                {
                    "endpoint_uri": null,
                    "inner_port": 80,
                    "outer_port": null,
                    "port_name": "http",
                    "protocol": "tcp",
                    "published": false,
                    "uri_protocol": "http"
                }
            ],
            "cpu_shares": null,
            "deployed_datetime": "Fri, 24 Apr 2015 16:30:47 +0000",
            "destroyed_datetime": null,
            "docker_id": "9ab356fd7aa7b7e23f31a0eb952e7c31ff4182136767370e15cb6321651a0885",
            "entrypoint": "",
            "exit_code": null,
            "exit_code_msg": null,
            "image_name": "tutum/hello-world:latest",
            "image_tag": "/api/v1/image/tutum/hello-world/tag/latest/",
            "layer": "/api/v1/layer/529c404c672f7370316815c0ccd6b544a8b1baf72741a3de601457a58f205cf9/",
            "memory": null,
            "name": "my-new-app-1",
            "node": "/api/v1/node/89226618-4cbf-44a7-b354-0edd6e251068/",
            "private_ip": "10.7.0.1",
            "privileged": false,
            "public_dns": "my-new-app-1.maximeheckel.cont.tutum.io",
            "resource_uri": "/api/v1/container/dcbe16b4-21a1-474b-a814-131a3626b1de/",
            "run_command": "/run.sh",
            "service": "/api/v1/service/02522970-a79a-46d6-8a64-475bf52e4258/",
            "started_datetime": "Fri, 24 Apr 2015 16:30:47 +0000",
            "state": "Running",
            "stopped_datetime": null,
            "synchronized": true,
            "uuid": "dcbe16b4-21a1-474b-a814-131a3626b1de"
        },
        {
            "autodestroy": "OFF",
            "autorestart": "OFF",
            "container_ports": [
                {
                    "endpoint_uri": null,
                    "inner_port": 80,
                    "outer_port": null,
                    "port_name": "http",
                    "protocol": "tcp",
                    "published": false,
                    "uri_protocol": "http"
                }
            ],
            "cpu_shares": null,
            "deployed_datetime": "Fri, 24 Apr 2015 16:30:56 +0000",
            "destroyed_datetime": null,
            "docker_id": "1c09ffe60df381ff854f5625da632d152f109b0f2962be10bb43919d3de23fd4",
            "entrypoint": "",
            "exit_code": null,
            "exit_code_msg": null,
            "image_name": "tutum/hello-world:latest",
            "image_tag": "/api/v1/image/tutum/hello-world/tag/latest/",
            "layer": "/api/v1/layer/529c404c672f7370316815c0ccd6b544a8b1baf72741a3de601457a58f205cf9/",
            "memory": null,
            "name": "my-new-app-2",
            "node": "/api/v1/node/89226618-4cbf-44a7-b354-0edd6e251068/",
            "private_ip": "10.7.0.2",
            "privileged": false,
            "public_dns": "my-new-app-2.maximeheckel.cont.tutum.io",
            "resource_uri": "/api/v1/container/3d84762f-46de-4c6a-b81c-2a11923df9db/",
            "run_command": "/run.sh",
            "service": "/api/v1/service/02522970-a79a-46d6-8a64-475bf52e4258/",
            "started_datetime": "Fri, 24 Apr 2015 16:30:56 +0000",
            "state": "Running",
            "stopped_datetime": null,
            "synchronized": true,
            "uuid": "3d84762f-46de-4c6a-b81c-2a11923df9db"
        }
    ]
}`
	getContainerResponse = "{ OFF  OFF [{ 80 0 tcp}]  0 Thu, 23 Apr 2015 10:37:13 +0000   0  tutum/hello-world:latest /api/v1/image/tutum/hello-world/tag/latest/ my-new-app-1 my-new-app-1.maximeheckel.cont.tutum.io /api/v1/container/6ae5e8f7-e1ae-4825-b208-c459b6e333e9/ /run.sh Thu, 23 Apr 2015 10:37:13 +0000 Running   6ae5e8f7-e1ae-4825-b208-c459b6e333e9}"
)

func Test_ListContainers(t *testing.T) {
	User = "test"
	ApiKey = "test"
	os.Setenv("TUTUM_ENV", "test")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, containerListResponse)
	}))

	defer server.Close()
	url := server.URL + "/api/v1/container/"
	BaseUrl = server.URL + "/api/v1/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Get: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response CListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(response)

	list, err := ListContainers()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(list)

	os.Setenv("TUTUM_ENV", "")
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
