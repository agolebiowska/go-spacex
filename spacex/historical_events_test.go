package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestHistoricalEventsService_Get(t *testing.T) {
	setup()

	u := "https://api.spacexdata.com/v3/history/"
	var tests = []struct {
		testName string
		ID       int
		wantErr  error
		want     []byte
	}{
		{"Valid response", 1, nil, WantJSON(u+"1", t)},
		{"Non existent ID", 50, ErrNotFound, WantJSON(u+"50", t)},
		{"Invalid ID", 0, ErrInvalidID, nil},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := client.HistoricalEvents.Get(test.ID)
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

func TestHistoricalEventsService_ListAll(t *testing.T) {
	setup()
	want := WantJSON("https://api.spacexdata.com/v3/history?limit=4", t)
	res, err := client.HistoricalEvents.ListAll(NewLimitOption(4), nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	Compare(got, want, t)
}
