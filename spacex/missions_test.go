package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestMissionsService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/missions/"
	var tests = []struct {
		testName string
		ID       string
		wantErr  error
		want     []byte
	}{
		{"Valid response", "F3364BF", nil, WantJSON(u+"F3364BF", t)},
		{"Non existent ID", "non-existent-id", ErrNotFound, WantJSON(u+"non-existent-id", t)},
		{"Invalid ID", "", ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Missions.Get(test.ID)
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

func TestMissionsService_ListAll(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/missions?limit=4", t)
	res, err := client.Missions.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
