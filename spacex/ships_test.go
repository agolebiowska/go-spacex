package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestShipsService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/ships/"
	var tests = []struct {
		testName string
		ID       string
		wantErr  error
		want     []byte
	}{
		{"Valid response", "AMERICANCHAMPION", nil, WantJSON(u+"AMERICANCHAMPION", t)},
		{"Non existent ID", "non-existent-id", ErrNotFound, WantJSON(u+"non-existent-id", t)},
		{"Invalid ID", "", ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.Ships.Get(test.ID)
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
