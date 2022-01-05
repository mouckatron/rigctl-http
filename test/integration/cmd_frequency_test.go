package integrationtest

import (
	"testing"

	"github.com/andreyvit/diff"
)

func TestGetFrequency(t *testing.T) {

	want := `{
    "success": true,
    "error": "",
    "raw": "get_freq:\nFrequency: 145000000\nRPRT 0\n",
    "data": {
        "frequency": "145000000"
    }
}`

	got, err := httpGet("/frequency")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}

	// j, err := GetJSON(body)
	// if err != nil {
	// 	t.Errorf("Error calling API, %s", err)
	// }

	// if data, ok := j["data"].(map[string]string); ok {
	// 	if data["frequency"] != "145000000" {
	// 		t.Errorf("Did not get a frequency, %+v", j)
	// 	}
	// }
}

func TestSetFrequency(t *testing.T) {

	want := `{
    "success": true,
    "error": "",
    "raw": "set_freq: 145500000\nRPRT 0\n",
    "data": {
        "frequency": "145500000"
    }
}`

	got, err := httpPostJSON("/frequency", []byte(`{"frequency":"145500000"}`))
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}
