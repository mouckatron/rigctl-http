package integrationtest

import (
	"testing"

	"github.com/andreyvit/diff"
)

// hoping this runs first to catch the case where the opts are not yet loaded
// though that should be dealt with in unit tests!?
func TestVFOOpNoop(t *testing.T) {
	want := `{
    "success": false,
    "error": "Option not available on this rig, see data for available options",
    "raw": "",
    "data": {
        "options": [
            "CPY",
            "XCHG",
            "FROM_VFO",
            "TO_VFO",
            "MCL",
            "UP",
            "DOWN",
            "BAND_UP",
            "BAND_DOWN",
            "LEFT",
            "RIGHT",
            "TUNE",
            "TOGGLE"
        ]
    }
}`

	got, err := httpPostJSON("/vfo_op", []byte(`{"op":"NOOP"}`))
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestVFOOp_options(t *testing.T) {
	want := `{
    "success": true,
    "error": "",
    "raw": "vfo_op: ?\nCPY XCHG FROM_VFO TO_VFO MCL UP DOWN BAND_UP BAND_DOWN LEFT RIGHT TUNE TOGGLE \nRPRT 0\n",
    "data": {
        "options": [
            "CPY",
            "XCHG",
            "FROM_VFO",
            "TO_VFO",
            "MCL",
            "UP",
            "DOWN",
            "BAND_UP",
            "BAND_DOWN",
            "LEFT",
            "RIGHT",
            "TUNE",
            "TOGGLE"
        ]
    }
}`

	got, err := httpGet("/vfo_op/_opts")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestVFOOpTuneOp(t *testing.T) {
	want := `{
    "success": true,
    "error": "",
    "raw": "vfo_op: TUNE\nRPRT 0\n",
    "data": {
        "op": "TUNE"
    }
}`

	got, err := httpPostJSON("/vfo_op", []byte(`{"op":"TUNE"}`))
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}
