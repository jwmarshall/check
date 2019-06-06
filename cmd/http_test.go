package cmd

import "testing"

func TestCheckHttp(t *testing.T) {
	url := "https://www.google.ca"
	status := 200
	err := checkHTTP(url, status)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckHTTPRedirect(t *testing.T) {
	url := "http://google.ca"
	status := 301
	err := checkHTTP(url, status)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckHTTPInvalidUrl(t *testing.T) {
	Expected := "parse invalid-url: invalid URI for request"
	url := "invalid-url"
	status := 200
	err := checkHTTP(url, status)
	if err.Error() != Expected {
		t.Errorf("Expected error: %v, got %v", Expected, err)
	}
}
