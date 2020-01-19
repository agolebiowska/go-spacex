package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestCapsulesService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/capsules/"
	var tests = []struct {
		testName string
		ID       string
		wantErr  error
		want     []byte
	}{
		{"Valid response", "C112", nil, WantJSON(u+"C112", t)},
		{"Non existent ID", "non-existent-id", ErrNotFound, WantJSON(u+"non-existent-id", t)},
		{"Invalid ID", "", ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Capsules.Get(test.ID)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: %v\n want: %v", errors.Cause(err), test.wantErr)
			}

			if test.wantErr == nil {
				got, err := json.Marshal(res)
				if err != nil {
					t.Fatal(err)
				}

				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("got: %+v\n want: %+v", string(got), string(test.want))
				}
			}
		})
	}
}

func TestCapsulesService_List(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/capsules?limit=4", t)
	res, err := client.Capsules.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}

func TestCapsulesService_ListUpcoming(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/capsules/upcoming?limit=4", t)
	res, err := client.Capsules.ListUpcoming(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}

func TestCapsulesService_ListPast(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/capsules/past?limit=4", t)
	res, err := client.Capsules.ListPast(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
