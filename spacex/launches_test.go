package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestLaunchesService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/launches/"
	var tests = []struct {
		testName string
		ID       int
		wantErr  error
		want     []byte
	}{
		{"Valid response", 65, nil, WantJSON(u+"65", t)},
		{"Non existent ID", 999, ErrNotFound, WantJSON(u+"999", t)},
		{"Invalid ID", 0, ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Launches.Get(test.ID)
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

func TestLaunchesService_ListAll(t *testing.T) {
	setup()
	_, err := client.Launches.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLaunchesService_ListPast(t *testing.T) {
	setup()
	_, err := client.Launches.ListPast(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLaunchesService_ListUpcoming(t *testing.T) {
	setup()
	_, err := client.Launches.ListUpcoming(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
}
