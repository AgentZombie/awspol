package awspol

// http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	EffectAllow = "Allow"
	EffectDeny  = "Deny"
)

type StatementEntry struct {
	Condition Condition   `json:",omitempty"`
	Sid       string      `json:",omitempty"`
	Effect    string      `json:",omitempty"`
	Principal *Principal  `json:",omitempty"`
	Action    MultiString `json:",omitempty"`
	Resource  MultiString `json:",omitempty"`
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

type statementEntryJSON struct {
	Condition    Condition   `json:",omitempty"`
	Sid          string      `json:",omitempty"`
	Effect       string      `json:",omitempty"`
	Principal    *Principal  `json:",omitempty"`
	NotPrincipal *Principal  `json:",omitempty"`
	Action       MultiString `json:",omitempty"`
	Resource     MultiString `json:",omitempty"`
}

func (e StatementEntry) MarshalJSON() ([]byte, error) {
	sej := statementEntryJSON{
		ConditionJSON: e.ConditionJSON,
		Sid:           e.Sid,
		Effect:        e.Effect,
		Action:        e.Action,
		Resource:      e.Resource,
	}
	if e.Principal != nil && e.Principal.Invert {
		sej.NotPrincipal = e.Principal
	} else {
		sej.Principal = e.Principal
	}
	return json.Marshal(sej)
}

func (e *StatementEntry) UnmarshalJSON(b []byte) error {
	sej := statementEntryJSON{}
	if err := json.Unmarshal(b, &sej); err != nil {
		return errors.Wrap(err, "unmarshalling StatementEntry")
	}
	if sej.Principal != nil && sej.NotPrincipal != nil {
		return errors.New("Statement cannot have both Principal and NotPrincipal")
	}
	e.ConditionJSON = sej.ConditionJSON
	e.Sid = sej.Sid
	e.Effect = sej.Effect
	e.Action = sej.Action
	e.Resource = sej.Resource
	if sej.Principal != nil {
		e.Principal = sej.Principal
	} else if sej.NotPrincipal != nil {
		e.Principal = sej.NotPrincipal
		e.Principal.Invert = true
	}
	return nil
}
