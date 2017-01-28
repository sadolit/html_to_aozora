package main

import "testing"

var testRubyToBunkoData = []struct {
	in       string
	expected string
}{
	{"", ""},
}

func TestRubyToBunko(t *testing.T) {
	for _, testData := range testRubyToBunkoData {
		actualBunko := RubyToBunko(testData.in)
		if testData.expected != actualBunko {
			t.Errorf("Input: [%s].Expected [%s], but it was [%s] instead.", testData.in, testData.expected, actualBunko)
		}
	}
}
