package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)
	if ! errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected ErrNoAuthHeaderIncluded")
	}

	header = http.Header{}
	header.Add("Authorization", "ApiKey 123456")

	want := "123456"

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected no error but got error")
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	header = http.Header{}
	header.Add("Authorization", "badvalue")

	_, err = GetAPIKey(header)
	wanterr := errors.New("malformed authorization header")

	if errors.Is(wanterr, err) {
		t.Fatalf("expected: %v, got: %v", wanterr, err)
	}

}
