package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	type test struct {
		header http.Header
		result string
		err    error
	}

	header1 := http.Header{}
	header1.Add("Authorization", "a")

	header2 := http.Header{}
	header2.Add("Authorization", "can split")

	header3 := http.Header{}
	header3.Add("Authorization", "ApiKey my-secret-key")

	tests := []test{
		{
			header: http.Header{},
			result: "",
			err:    ErrNoAuthHeaderIncluded,
		},
		{
			header: header1,
			result: "",
			err:    errors.New("malformed authorization header"),
		},
		{
			header: header2,
			result: "",
			err:    errors.New("malformed authorization header"),
		},
		{
			header: header3,
			result: "my-secret-key",
			err:nil,
		},
	}

	for _, tc := range tests {
		result, err := GetAPIKey(tc.header)

		if result != tc.result {
			t.Errorf("Result mismatch %s != %s", result, tc.result)
		}

		if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Error mismatch")
		}
	}

}
