package helper

import "testing"

func TestStr2Int(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"00:00:10,000", 10000},
		{"00:00:00,000", 0},
		{"00:04:04,286", 244286},
		{"01:40:19,000", 6019000},
		{"01:40:29,000", 6029000},
	}

	for _, c := range cases {
		got := int(Str2milliseconds(c.in))
		if got != c.want {
			t.Errorf("milliseconds return error with %s, want %d but got: %d.\n", c.in, c.want, got)
		}
	}
}

func TestInt2Str(t *testing.T) {

	cases := []struct {
		in   int
		want string
	}{
		{10000, "00:00:10,000"},
		{0, "00:00:00,000"},
		{244286, "00:04:04,286"},
		{6019000, "01:40:19,000"},
		{6029000, "01:40:29,000"},
	}

	for _, c := range cases {
		got := Milliseconds2str(c.in)
		if got != c.want {
			t.Errorf("milliseconds return error with %d, want %s but got: %s.\n", c.in, c.want, got)
		}
	}
}
