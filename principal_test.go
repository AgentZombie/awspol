package awspol

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPrincipalExactlyEquals(t *testing.T) {
	for _, tc := range []struct {
		a    Principal
		b    Principal
		want bool
	}{
		{
			a: Principal{
				Invert: true,
				All:    false,
				AWS:    MultiString{"a", "b", "c"},
			},
			b: Principal{
				Invert: true,
				All:    false,
				AWS:    MultiString{"a", "b", "c"},
			},
			want: true,
		},
		{
			a: Principal{
				Invert: true,
				All:    false,
				AWS:    MultiString{"a", "b", "c"},
			},
			b: Principal{
				Invert:  true,
				All:     false,
				Service: MultiString{"a", "b", "c"},
			},
			want: false,
		},
		{
			a: Principal{
				Invert: true,
				All:    false,
				AWS:    MultiString{"a", "b", "c"},
			},
			b: Principal{
				Invert: true,
				All:    false,
				AWS:    MultiString{"a", "c", "b"},
			},
			want: false,
		},
	} {
		if got := tc.a.ExactlyEquals(tc.b); got != tc.want {
			t.Fatalf("got %v, want %v on %#v.ExactlyEquals(%#v)", got, tc.want, tc.a, tc.b)
		}
	}
}

func TestPrincipalJSONSymmetry(t *testing.T) {
	for _, tc := range []Principal{
		Principal{
			All:     false,
			AWS:     MultiString{"a", "b", "c"},
			Service: MultiString{"c"},
		},
		Principal{
			All: true,
		},
	} {
		t.Logf("Testing %#v", tc)
		fstEnc, err := json.Marshal(tc)
		if err != nil {
			t.Fatalf("unexpected error marshaling original Principal: %q", err)
		}
		p2 := Principal{}
		err = json.Unmarshal(fstEnc, &p2)
		if err != nil {
			t.Fatalf("unexpected error unmarshaling first encoding: %q", err)
		}
		if !tc.ExactlyEquals(p2) {
			t.Logf("unmarshalled: %s", string(fstEnc))
			t.Fatalf("got %#v, want %#v from unmarshal", p2, tc)
		}
		sndEnc, err := json.Marshal(p2)
		if err != nil {
			t.Fatalf("unexpected error marshaling reconstituted Principal: %q", err)
		}
		if !bytes.Equal(fstEnc, sndEnc) {
			t.Fatalf("got %s, want %s on second marshal", string(sndEnc), string(fstEnc))
		}
	}
}
