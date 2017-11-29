package awspol

import "testing"

func TestMultiStringExactlyEquals(t *testing.T) {
	for _, tc := range []struct {
		a      MultiString
		b      MultiString
		expect bool
	}{
		{
			a:      MultiString{"a", "b", "c"},
			b:      MultiString{"a", "b", "c"},
			expect: true,
		},
		{
			a:      MultiString{"a", "b", "c"},
			b:      MultiString{"a", "c", "b"},
			expect: false,
		},
		{
			a:      MultiString{"a", "b", "c"},
			b:      MultiString{"a", "b", "C"},
			expect: false,
		},
		{
			a:      MultiString{"a", "b"},
			b:      MultiString{"a", "b", "c"},
			expect: false,
		},
		{
			a:      MultiString{},
			b:      MultiString{},
			expect: true,
		},
		{
			a:      nil,
			b:      nil,
			expect: true,
		},
		{
			a:      MultiString{},
			b:      nil,
			expect: true,
		},
	} {
		if got := tc.a.ExactlyEquals(tc.b); got != tc.expect {
			t.Fatalf("got %v, want %v in %#v.ExactlyEquals(%#v)", got, tc.expect, tc.a, tc.b)
		}
	}
}
