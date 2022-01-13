package integrationtest

import (
	"testing"

	"github.com/andreyvit/diff"
)

func TestGetPowerstat(t *testing.T) {
	want := `{
    "success": true,
    "error": "",
    "raw": "get_powerstat:\nPower Status: 1\nRPRT 0\n",
    "data": {
        "status": 1
    }
}`

	got, err := httpGet("/powerstat")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestSetPowerstat(t *testing.T) {
	want := `{
    "success": true,
    "error": "",
    "raw": "set_powerstat: 1\nRPRT 0\n",
    "data": {
        "status": 1
    }
}`

	got, err := httpPutJSON("/powerstat", []byte(`{"status": 1}`))
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}
