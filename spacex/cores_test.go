package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestCoresService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/cores/"
	var tests = []struct {
		testName string
		ID       string
		wantErr  error
		want     []byte
	}{
		{"Valid response", "B1042", nil, WantJSON(u+"B1042", t)},
		{"Non existent ID", "non-existent-id", ErrNotFound, WantJSON(u+"non-existent-id", t)},
		{"Invalid ID", "", ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Cores.Get(test.ID)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: %v\n want: %v", errors.Cause(err), test.wantErr)
			}

			if test.wantErr == nil {
				got, err := json.Marshal(res)
				if err != nil {
					t.Fatal(err)
				}
				Compare(got, test.want, t)
			}
		})
	}
}

func TestCoresService_ListAll(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/cores?limit=4", t)
	res, err := client.Cores.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}

func TestCoresService_ListUpcoming(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/cores/upcoming?limit=4", t)
	res, err := client.Cores.ListUpcoming(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}

func TestCoresService_ListPast(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/cores/past?limit=4", t)
	res, err := client.Cores.ListPast(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
