package awspol

// http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html

import (
	"encoding/json"

	"github.com/AgentZombie/multistring"
	"github.com/pkg/errors"
)

const (
	EffectAllow = "Allow"
	EffectDeny  = "Deny"
)

type StatementEntry struct {
	Condition Condition               `json:",omitempty"`
	Sid       string                  `json:",omitempty"`
	Effect    string                  `json:",omitempty"`
	Principal *Principal              `json:",omitempty"`
	Action    multistring.MultiString `json:",omitempty"`
	Resource  multistring.MultiString `json:",omitempty"`
}

func (e StatementEntry) ExactlyEquals(o StatementEntry) bool {
	if e.Effect != o.Effect {
		return false
	}
	if e.Sid != o.Sid {
		return false
	}
	if !e.Resource.ExactlyEquals(o.Resource) {
		return false
	}
	if !e.Action.ExactlyEquals(o.Action) {
		return false
	}
	if !e.Condition.ExactlyEquals(o.Condition) {
		return false
	}
	if !e.Principal.ExactlyEquals(o.Principal) {
		return false
	}
	return true
}

func (e StatementEntry) EquivalentTo(o StatementEntry) bool {
	if e.Effect != o.Effect {
		return false
	}
	if e.Sid != o.Sid {
		return false
	}
	if !e.Resource.EquivalentTo(o.Resource) {
		return false
	}
	if !e.Action.EquivalentTo(o.Action) {
		return false
	}
	if !e.Condition.EquivalentTo(o.Condition) {
		return false
	}
	if !e.Principal.EquivalentTo(o.Principal) {
		return false
	}
	return true
}

type statementEntryJSON struct {
	Condition    Condition               `json:",omitempty"`
	Sid          string                  `json:",omitempty"`
	Effect       string                  `json:",omitempty"`
	Principal    *Principal              `json:",omitempty"`
	NotPrincipal *Principal              `json:",omitempty"`
	Action       multistring.MultiString `json:",omitempty"`
	Resource     multistring.MultiString `json:",omitempty"`
}

func (e StatementEntry) MarshalJSON() ([]byte, error) {
	sej := statementEntryJSON{
		Condition: e.Condition,
		Sid:       e.Sid,
		Effect:    e.Effect,
		Action:    e.Action,
		Resource:  e.Resource,
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
	e.Condition = sej.Condition
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
