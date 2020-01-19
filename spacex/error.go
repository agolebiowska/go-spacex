package spacex

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrInvalidID   = errors.New("invalid ID")
	ErrNoResults   = errors.New("results are empty")
	ErrInvalidJSON = errors.New("invalid JSON")
)

var (
	ErrBadRequest = ServerError{
		Status: http.StatusBadRequest,
		Msg:    "bad request: check query parameters",
	}
	ErrNotFound = ServerError{
		Status: http.StatusNotFound,
		Msg:    "nof found",
	}
	ErrUnauthorized = ServerError{
		Status: http.StatusUnauthorized,
		Msg:    "authentication failed: check for valid API key in user-key header",
	}
	ErrForbidden = ServerError{
		Status: http.StatusForbidden,
		Msg:    "authentication failed: check for valid API key in user-key header",
	}
	ErrInternalError = ServerError{
		Status: http.StatusInternalServerError,
		Msg:    "internal error: report bug",
	}
)

type ServerError struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (e ServerError) Error() string {
	return "spaceX server error: status: " + strconv.Itoa(e.Status) + " message: " + e.Msg
}

func checkResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusInternalServerError:
		return ErrInternalError
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(b))

	var e ServerError
	err = json.Unmarshal(b, &e)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal server error message")
	}

	return e
}

const (
	// openBracketASCII represents the ASCII code for an open bracket.
	openBracketASCII = 91
	// closedBracketASCII represents the ASCII code for a closed bracket.
	closedBracketASCII = 93
)

func isBracketPair(b []byte) bool {
	if len(b) != 2 {
		return false
	}

	if b[0] == openBracketASCII && b[1] == closedBracketASCII {
		return true
	}

	return false
}
