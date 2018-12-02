package main

import "testing"

func TestGreet(t *testing.T) {
	t.Helper()
	cases := []struct{
		name string
		input int
		expect string
	}{
		{"0:00", 0, "Morning"},
		{"12:00", 12, "Afternoon"},
		{"19:00", 19, "Evening"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Greet(c.input)
			if actual != c.expect {
				t.Errorf("%s should be %s, but actuail is %s", actual, c.expect, actual)
			}
		})
	}
}
