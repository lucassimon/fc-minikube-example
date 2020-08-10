package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		parameter string
		expected  string
	}{
		{"", "Code.education Rocks!"},
		{"Hello World", "Hello World"},
	}

	for _, table := range tables {
		var greeting string = Greeting(table.parameter)
		if greeting != table.expected {
			t.Errorf("Result error: %s. An error occurred when call Greeting with %s", table.expected, table.parameter)
		} else {
			t.Logf("Success")
		}
	}
}
