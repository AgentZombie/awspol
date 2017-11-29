package awspol

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type MultiString []string

func (ms *MultiString) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	str := string(b)
	switch str[0] {
	case '[':
		data := []string{}
		if err := json.Unmarshal(b, &data); err != nil {
			return errors.Wrap(err, "unmarshalling MultiString array")
		}
		if len(data) > 0 {
			*ms = data
		}
	case '"':
		data := ""
		if err := json.Unmarshal(b, &data); err != nil {
			return errors.Wrap(err, "unmarshalling MultiString string")
		}
		if len(data) > 0 {
			*ms = []string{data}
		}
	default:
		return errors.New("expecting JSON array or string")
	}

	return nil
}

func (ms MultiString) MarshalJSON() ([]byte, error) {
	switch len(ms) {
	case 0:
		return nil, nil
	case 1:
		return json.Marshal(ms[0])
	default:
		return json.Marshal([]string(ms))
	}
}

func (ms MultiString) ExactlyEquals(o MultiString) bool {
	if len(ms) != len(o) {
		return false
	}
	for i := range ms {
		if ms[i] != o[i] {
			return false
		}
	}
	return true
}
