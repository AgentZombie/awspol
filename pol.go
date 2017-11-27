package awspol

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type PolicyDocument struct {
	Version   string           `json:",omitempty"`
	Statement []StatementEntry `json:",omitempty"`
}

type StatementEntry struct {
	Sid       string          `json:",omitempty"`
	Effect    string          `json:",omitempty"`
	Principal json.RawMessage `json:",omitempty"`
	Action    MultiString     `json:",omitempty"`
	Resource  MultiString     `json:",omitempty"`
}

func (e StatementEntry) Equals(o StatementEntry) bool {
	if e.Effect != o.Effect {
		return false
	}
	if len(e.Resource) != len(o.Resource) {
		return false
	}
	for _, eRes := range e.Resource {
		matched := false
		for _, oRes := range o.Resource {
			if eRes == oRes {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	if len(e.Action) != len(o.Action) {
		return false
	}
	for _, eAct := range e.Action {
		matched := false
		for _, oAct := range o.Action {
			if eAct == oAct {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func (d PolicyDocument) Equals(o PolicyDocument) bool {
	if d.Version != o.Version {
		return false
	}
	if len(d.Statement) != len(o.Statement) {
		return false
	}
	for _, dStat := range d.Statement {
		matched := false
		for _, oStat := range o.Statement {
			if dStat.Equals(oStat) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func ParsePolicyDocument(s string) (PolicyDocument, error) {
	pd := PolicyDocument{}
	err := json.Unmarshal([]byte(s), &pd)
	if err != nil {
		return pd, errors.Wrap(err, "unmarshaling policy JSON")
	}
	return pd, nil
}
