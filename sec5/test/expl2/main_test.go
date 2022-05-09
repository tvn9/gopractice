// Testing in Go example
package main

import "testing"

func TestValidEmail(t *testing.T) {
	data := "email@example.com"
	if !ValidEmail(data) {
		t.Errorf("ValidEmail(%v)=false, expected true", data)
	}
}

func TestValidEmailTable(t *testing.T) {
	table := []struct {
		email string
		want  bool
	}{
		{"email@example.com", true},
		{"missing@tld", false},
	}

	for _, data := range table {
		result := ValidEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %t, want: %t", data.email, result, data.want)
		}
	}
}
