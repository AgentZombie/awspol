package awspol

import "testing"

func TestConditionSameContentsAs(t *testing.T) {
	for _, tc := range []struct {
		a    Condition
		b    Condition
		want bool
	}{
		{
			a: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			b: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			want: true,
		},
		{
			a: Condition{
				CondOpStringNotEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			b: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			want: false,
		},
		{
			a: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			b: Condition{
				CondOpStringEquals: CondOp{
					Key:   "foo",
					Value: MultiString{"a", "b"},
				},
			},
			want: false,
		},
		{
			a: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"a", "b"},
				},
			},
			b: Condition{
				CondOpStringEquals: CondOp{
					Key:   "blarg",
					Value: MultiString{"b", "a"},
				},
			},
			want: false,
		},
	} {
		t.Logf("Testing %#v", tc)
		if got := tc.a.SameContentsAs(tc.b); got != tc.want {
			t.Fatalf("got %v, want %v in %#v.SameContentsAs(%#v)", got, tc.want, tc.a, tc.b)
		}
	}
}
