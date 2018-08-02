package awspol

import "testing"

func TestConditionEquivalentTo(t *testing.T) {
	for _, tc := range []struct {
		label string
		a     Condition
		b     Condition
		want  bool
	}{
		{
			label: "exactly the same",
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
			label: "only condop differs",
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
			label: "differing key",
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
			label: "differing value order",
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
			want: true,
		},
	} {
		t.Logf("Testing %#v", tc)
		if got := tc.a.EquivalentTo(tc.b); got != tc.want {
			t.Fatalf("got %v, want %v in %#v.EquivalentTo(%#v) for %v", got, tc.want, tc.a, tc.b, tc.label)
		}
	}
}
