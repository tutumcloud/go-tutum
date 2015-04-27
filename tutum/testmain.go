package tutum

import (
	"fmt"
	"io/ioutil"
	"os"
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
