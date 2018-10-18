package awspol

import (
	"encoding/json"

	"github.com/AgentZombie/multistring"
	"github.com/pkg/errors"
)

const (
	CondOpStringEquals                      CondType = "StringEquals"
	CondOpStringNotEquals                            = "StringNotEquals"
	CondOpStringEqualsIgnoreCase                     = "StringEqualsIgnoreCase"
	CondOpStringNotEqualsIgnoreCase                  = "StringNotEqualsIgnoreCase"
	CondOpStringLike                                 = "StringLike"
	CondOpStringNotLike                              = "StringNotLike"
	CondOpNumericEquals                              = "NumericEquals"
	CondOpNumericNotEquals                           = "NumericNotEquals"
	CondOpNumericLessThan                            = "NumericLessThan"
	CondOpNumericLessThanEquals                      = "NumericLessThanEquals"
	CondOpNumericGreaterThan                         = "NumericGreaterThan"
	CondOpNumericGreaterThanEquals                   = "NumericGreaterThanEquals"
	CondOpDateEquals                                 = "DateEquals"
	CondOpDateNotEquals                              = "DateNotEquals"
	CondOpDateLessThan                               = "DateLessThan"
	CondOpDateLessThanEquals                         = "DateLessThanEquals"
	CondOpDateGreaterThan                            = "DateGreaterThan"
	CondOpDateGreaterThanEquals                      = "DateGreaterThanEquals"
	CondOpBool                                       = "Bool"
	CondOpBinaryEquals                               = "BinaryEquals"
	CondOpIpAddress                                  = "IpAddress"
	CondOpNotIpAddress                               = "NotIpAddress"
	CondOpArnEquals                                  = "ArnEquals"
	CondOpArnLike                                    = "ArnLike"
	CondOpArnNotEquals                               = "ArnNotEquals"
	CondOpArnNotLike                                 = "ArnNotLike"
	CondOpStringEqualsIfExists                       = "StringEqualsIfExists"
	CondOpStringNotEqualsIfExists                    = "StringNotEqualsIfExists"
	CondOpStringEqualsIgnoreCaseIfExists             = "StringEqualsIgnoreCaseIfExists"
	CondOpStringNotEqualsIgnoreCaseIfExists          = "StringNotEqualsIgnoreCaseIfExists"
	CondOpStringLikeIfExists                         = "StringLikeIfExists"
	CondOpStringNotLikeIfExists                      = "StringNotLikeIfExists"
	CondOpNumericEqualsIfExists                      = "NumericEqualsIfExists"
	CondOpNumericNotEqualsIfExists                   = "NumericNotEqualsIfExists"
	CondOpNumericLessThanIfExists                    = "NumericLessThanIfExists"
	CondOpNumericLessThanEqualsIfExists              = "NumericLessThanEqualsIfExists"
	CondOpNumericGreaterThanIfExists                 = "NumericGreaterThanIfExists"
	CondOpNumericGreaterThanEqualsIfExists           = "NumericGreaterThanEqualsIfExists"
	CondOpDateEqualsIfExists                         = "DateEqualsIfExists"
	CondOpDateNotEqualsIfExists                      = "DateNotEqualsIfExists"
	CondOpDateLessThanIfExists                       = "DateLessThanIfExists"
	CondOpDateLessThanEqualsIfExists                 = "DateLessThanEqualsIfExists"
	CondOpDateGreaterThanIfExists                    = "DateGreaterThanIfExists"
	CondOpDateGreaterThanEqualsIfExists              = "DateGreaterThanEqualsIfExists"
	CondOpBoolIfExists                               = "BoolIfExists"
	CondOpBinaryEqualsIfExists                       = "BinaryEqualsIfExists"
	CondOpIpAddressIfExists                          = "IpAddressIfExists"
	CondOpNotIpAddressIfExists                       = "NotIpAddressIfExists"
	CondOpArnEqualsIfExists                          = "ArnEqualsIfExists"
	CondOpArnLikeIfExists                            = "ArnLikeIfExists"
	CondOpArnNotEqualsIfExists                       = "ArnNotEqualsIfExists"
	CondOpArnNotLikeIfExists                         = "ArnNotLikeIfExists"
	CondOpNull                                       = "Null"
)

type CondType string

type CondOp struct {
	Key   string
	Value multistring.MultiString
}

func (c CondOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]multistring.MultiString{
		c.Key: c.Value,
	})
}

func (c *CondOp) UnmarshalJSON(b []byte) error {
	m := map[string]multistring.MultiString{}
	if err := json.Unmarshal(b, &m); err != nil {
		return errors.Wrap(err, "unmarshalling CondOp")
	}
	if len(m) > 1 {
		return errors.New("invalid CondOp")
	}
	for k, v := range m {
		c.Key = k
		c.Value = v
	}
	return nil
}

func (c CondOp) EquivalentTo(o CondOp) bool {
	if c.Key != o.Key {
		return false
	}
	if !c.Value.EquivalentTo(o.Value) {
		return false
	}
	return true
}

func (c CondOp) ExactlyEquals(o CondOp) bool {
	if c.Key != o.Key {
		return false
	}
	if !c.Value.ExactlyEquals(o.Value) {
		return false
	}
	return true
}

type Condition map[CondType]CondOp

func (c Condition) ExactlyEquals(o Condition) bool {
	if len(c) != len(o) {
		return false
	}
	for k, v := range c {
		ov, ok := o[k]
		if !ok {
			return false
		}
		if !v.ExactlyEquals(ov) {
			return false
		}
	}
	return true
}

func (c Condition) EquivalentTo(o Condition) bool {
	if len(c) != len(o) {
		return false
	}
	for k, v := range c {
		ov, ok := o[k]
		if !ok {
			return false
		}
		if !v.EquivalentTo(ov) {
			return false
		}
	}
	return true
}
