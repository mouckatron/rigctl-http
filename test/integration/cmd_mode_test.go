package integrationtest

import (
  "testing"

  "github.com/andreyvit/diff"
)

func TestGetModeNoop(t *testing.T) {
  want := `{
    "success": false,
    "error": "Option not available on this rig, see data for available options",
    "raw": "",
    "data": {
        "options": [
            "AM",
            "CW",
            "USB",
            "LSB",
            "RTTY",
            "FM",
            "WFM",
            "CWR",
            "RTTYR"
        ]
    }
}`

  got, err := httpPostJSON("/mode", []byte(`{"mode":"NOOP","passband":0}`))
  if err != nil {
    t.Error(err)
  }

  if string(got) != want {
    t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
  }
}

func TestMode_options(t *testing.T) {
  want := `{
    "success": true,
    "error": "",
    "raw": "set_mode: ?\nAM CW USB LSB RTTY FM WFM CWR RTTYR \nRPRT 0\n",
    "data": {
        "options": [
            "AM",
            "CW",
            "USB",
            "LSB",
            "RTTY",
            "FM",
            "WFM",
            "CWR",
            "RTTYR"
        ]
    }
}`

  got, err := httpGet("/mode/_options")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestGetMode(t *testing.T) {
  want := `{
    "success": true,
    "error": "",
    "raw": "get_mode:\nMode: FM\nPassband: 15000\nRPRT 0\n",
    "data": {
        "mode": "FM",
        "passband": 15000
    }
}`

  got, err := httpGet("/mode")
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
	}
}

func TestSetModeFM(t *testing.T) {
  want := `{
    "success": true,
    "error": "",
    "raw": "set_mode: AM 0\nRPRT 0\n",
    "data": {
        "mode": "AM",
        "passband": 0
    }
}`

  got, err := httpPostJSON("/mode", []byte(`{"mode":"AM","passband":0}`))
  if err != nil {
    t.Error(err)
  }

  if string(got) != want {
    t.Errorf("Did not get correct response\n%v", diff.LineDiff(want, string(got)))
  }
}
