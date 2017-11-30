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

type Condition map[CondType]CondOp
