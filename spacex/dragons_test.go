package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestDragonsService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/dragons/"
	var tests = []struct {
		testName string
		ID       string
		wantErr  error
		want     []byte
	}{
		{"Valid response", "dragon1", nil, WantJSON(u+"dragon1", t)},
		{"Non existent ID", "non-existent-id", ErrNotFound, WantJSON(u+"non-existent-id", t)},
		{"Invalid ID", "", ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Dragons.Get(test.ID)
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

func TestDragonsService_ListAll(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/dragons?limit=4", t)
	res, err := client.Dragons.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
