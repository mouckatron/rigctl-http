package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var testServerURL string = "http://localhost:8080"

// func TestEmptyEndpoint(t *testing.T) {
// 	_, err := httpgetJSON("/nil")
// 	if err != nil {
// 		t.Errorf("Expected error, %s", err)
// 	}
// }

// #############################################################################
// help functions

func httpGet(path string) ([]byte, error) {
	got, err := http.Get(url(path))

	if err != nil {
		fmt.Println(got)
		return nil, err
	}

	defer got.Body.Close()
	body, err := io.ReadAll(got.Body)

	return body, err
}

func httpGetJSON(path string) (map[string]interface{}, error) {

	body, err := httpGet(path)
	if err != nil {
		return nil, err
	}

	j, err := getJSON(body)
	if err != nil {
		fmt.Printf("%+v\n", j)
	}
	return j, err
}

func httpPostJSON(path string, data []byte) ([]byte, error) {
	got, err := http.Post(url(path), "application/json", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(got)
		return nil, err
	}

	defer got.Body.Close()
	body, err := io.ReadAll(got.Body)

	return body, err
}

func getJSON(response []byte) (map[string]interface{}, error) {
	decoded := make(map[string]interface{})
	err := json.Unmarshal(response, &decoded)

	return decoded, err
}

func url(path string) string {
	return fmt.Sprintf("%s%s", testServerURL, path)
}
