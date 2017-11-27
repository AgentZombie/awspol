package awspol

import (
	"encoding/json"
	"testing"
)

func TestStatementEntryUnmarshal(t *testing.T) {
	msArray := `{"Sid":"testSid","Effect":"Allow","Principal":{"AWS":"testArn"},"Action":["action1","action2"],"Resource":["res1","res2"]}`
	se := &StatementEntry{}
	if err := json.Unmarshal([]byte(msArray), &se); err != nil {
		t.Fatalf("unexpected error unmarshalling msArray: %q", err)
	}
	if len(se.Resource) != 2 || se.Resource[0] != "res1" || se.Resource[1] != "res2" {
		t.Fatalf("got Resource: %#v, want Resource; %#v on msStr", se.Resource, []string{"res1", "res2"})
	}

	msStr := `{"Sid":"testSid","Effect":"Allow","Principal":{"AWS":"testArn"},"Action":["action1","action2"],"Resource":"res1"}`
	se = &StatementEntry{}
	if err := json.Unmarshal([]byte(msStr), &se); err != nil {
		t.Fatalf("unexpected error unmarshalling msStr: %q", err)
	}
	if len(se.Resource) != 1 || se.Resource[0] != "res1" {
		t.Fatalf("got Resource: %#v, want Resource; %#v on msStr", se.Resource, []string{"res1"})
	}

	msArrayEmpty := `{"Sid":"testSid","Effect":"Allow","Principal":{"AWS":"testArn"},"Action":["action1","action2"],"Resource":[]}`
	se = &StatementEntry{}
	if err := json.Unmarshal([]byte(msArrayEmpty), &se); err != nil {
		t.Fatalf("unexpected error unmarshalling msArrayEmpty: %q", err)
	}

	msStrEmpty := `{"Sid":"testSid","Effect":"Allow","Principal":{"AWS":"testArn"},"Action":["action1","action2"],"Resource":""}`
	se = &StatementEntry{}
	if err := json.Unmarshal([]byte(msStrEmpty), &se); err != nil {
		t.Fatalf("unexpected error unmarshalling msStrEmpty: %q", err)
	}
}

func TestStatementEntryMarshal(t *testing.T) {
	for _, tc := range []struct {
		in     MultiString
		expect string
	}{
		{in: MultiString{"res1", "res2"}, expect: `{"Resource":["res1","res2"]}`},
		{in: MultiString{"res1"}, expect: `{"Resource":"res1"}`},
		{in: MultiString{}, expect: `{}`},
		{in: nil, expect: `{}`},
	} {
		se := &StatementEntry{Resource: tc.in}
		got, err := json.Marshal(se)
		if err != nil {
			t.Fatalf("unexected error mashalling %#v: %q", tc.in, err)
		}
		if string(got) != tc.expect {
			t.Fatalf("got %q, want %q on case %#v", string(got), tc.expect, tc.in)
		}
	}
}
