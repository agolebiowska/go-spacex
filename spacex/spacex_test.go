package spacex

import (
	"os/exec"
	"reflect"
	"testing"
)

var (
	client *Client
)

func setup() {
	client = NewClient(nil)
}

func WantJSON(u string, t *testing.T) []byte {
	want, err := exec.Command("curl", "-s", u).Output()
	if err != nil {
		t.Fatal(err)
	}

	return want
}

func Compare(got []byte, want []byte, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v\n want: %+v", string(got), string(want))
	}
}
