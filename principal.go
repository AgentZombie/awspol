package awspol

import (
	"encoding/json"

	"github.com/AgentZombie/multistring"
	"github.com/pkg/errors"
)

// http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html

type Principal struct {
	Invert    bool
	All       bool
	AWS       multistring.MultiString
	Federated multistring.MultiString
	Service   multistring.MultiString
}

type principalJSON struct {
	AWS       multistring.MultiString `json:",omitempty"`
	Federated multistring.MultiString `json:",omitempty"`
	Service   multistring.MultiString `json:",omitempty"`
}

func (p Principal) MarshalJSON() ([]byte, error) {
	if p.All {
		return []byte(`"*"`), nil
	}
	pp := principalJSON{
		AWS:       p.AWS,
		Federated: p.Federated,
		Service:   p.Service,
	}
	return json.Marshal(pp)
}

func (p *Principal) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	if len(b) < 3 {
		return errors.New("invalid Principal")
	}
	if string(b[:3]) == `"*"` {
		p.All = true
		return nil
	}
	pp := principalJSON{}
	if err := json.Unmarshal(b, &pp); err != nil {
		return errors.Wrap(err, "unmarshalling Principal")
	}
	p.AWS = pp.AWS
	p.Federated = pp.Federated
	p.Service = pp.Service
	return nil
}

func (p *Principal) ExactlyEquals(o *Principal) bool {
	if p == o {
		return true
	}
	if p == nil || o == nil {
		return false
	}
	if p.Invert != o.Invert {
		return false
	}
	if p.All != o.All {
		return false
	}
	if !p.AWS.ExactlyEquals(o.AWS) {
		return false
	}
	if !p.Federated.ExactlyEquals(o.Federated) {
		return false
	}
	if !p.Service.ExactlyEquals(o.Service) {
		return false
	}
	return true
}

func (p *Principal) EquivalentTo(o *Principal) bool {
	if p == o {
		return true
	}
	if p == nil || o == nil {
		return false
	}
	if p.Invert != o.Invert {
		return false
	}
	if p.All != o.All {
		return false
	}
	if !p.AWS.EquivalentTo(o.AWS) {
		return false
	}
	if !p.Federated.EquivalentTo(o.Federated) {
		return false
	}
	if !p.Service.EquivalentTo(o.Service) {
		return false
	}
	return true
}
