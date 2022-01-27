package integrationtest

import (
	"testing"

	"github.com/andreyvit/diff"
)

func TestGetTuningStep(t *testing.T) {

	want := `{
    "success": true,
    "error": "",
    "raw": "get_ts:\nTuning Step: 0\nRPRT 0\n",
    "data": {
        "step": 0
    }
}`

	got, err := httpGet("/tuning_step")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestSetTuningStep(t *testing.T) {

	want := `{
    "success": true,
    "error": "",
    "raw": "set_ts: 1000\nRPRT 0\n",
    "data": {
        "step": 1000
    }
}`

	got, err := httpPutJSON("/tuning_step", []byte(`{"step":1000}`))
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}
