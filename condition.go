package awspol

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
	Value MultiString
}

func (c CondOp) SameContentsAs(o CondOp) bool {
	if c.Key != o.Key {
		return false
	}
	if !c.Value.ExactlyEquals(o.Value) {
		return false
	}
	return true
}

type Condition map[CondType]CondOp

func (c Condition) SameContentsAs(o Condition) bool {
	if len(c) != len(o) {
		return false
	}
	for k, v := range c {
		ov, ok := o[k]
		if !ok {
			return false
		}
		if !v.SameContentsAs(ov) {
			return false
		}
	}
	return true
}
