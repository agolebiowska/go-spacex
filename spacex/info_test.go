package spacex

import (
	"encoding/json"
	"testing"
)

func TestInfoService_ListAll(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/info", t)
	res, err := client.Info.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
