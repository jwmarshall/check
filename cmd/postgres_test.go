package cmd

import "testing"

func TestCheckPostgres(t *testing.T) {
	host := "localhost"
	port := 5432
	user := "postgres"
	pass := "postgres"
	database := "postgres"
	err := checkPostgres(host, port, user, pass, database)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckPostgresConnectionError(t *testing.T) {
	Expected := "dial tcp 127.0.0.1:5434: connect: connection refused"
	host := "localhost"
	port := 5434
	user := "postgres"
	pass := "postgres"
	database := "postgres"
	err := checkPostgres(host, port, user, pass, database)
	if err != nil {
		if err.Error() != Expected {
			t.Errorf("Expected error: %v, got %v", Expected, err)
		}
	} else {
		t.Errorf("Expected error: %v, got %v", Expected, err)
	}
}
