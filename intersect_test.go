package glob_intersection

import "testing"

func TestIntersection(t *testing.T) {
	Debug = true

	tcs := []struct {
		id   int
		l, r string
		m    bool
	}{
		{1, `a`, `a`, true},
		{2, `ab`, `a`, false},

		{3, `*`, `*.log`, true},
		{4, `*`, `abc`, true},
		{5, `*.log`, `a.log`, true},

		{6, `*-live-*`, `-live-`, true},
		{7, `*-live-*`, `a-live-`, true},
		{8, `*-live-*`, `-live-b`, true},
		{9, `*-live-*`, `a-live-b`, true},

		{10, `ab*d`, `a*cd`, true},
		{11, `abc*def`, `abc.def`, true},
		{12, `abcd*g`, `a*defg`, true},
		{13, `*bc`, `ab*`, true},

		{14, `*`, `**`, true},
		{15, `a*b*c*d`, `**`, true},
		{16, `a*b*c*d`, `*********`, true},
	}

	debug := -1
	for _, tc := range tcs {
		if debug != -1 && tc.id != debug {
			continue
		}
		if x := Intersect(tc.l, tc.r); x != tc.m {
			t.Logf("%-3d: |%s| intersects |%s| ? got: %v, want: %v", tc.id, tc.l, tc.r, x, tc.m)
			t.Fail()
		}
	}
}
