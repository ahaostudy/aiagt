package types

import (
	"encoding/json"
	"fmt"

	"github.com/aiagt/aiagt/pkg/jsonutil"
	"github.com/pkg/errors"
)

type Any struct {
	value any
}

func NewAny(val ...any) *Any {
	if len(val) > 0 {
		return &Any{value: val[0]}
	}

	return &Any{value: new(any)}
}

func (a *Any) MarshalJSON() ([]byte, error) {
	if a == nil || a.value == nil {
		return []byte("null"), nil
	}

	return json.Marshal(a.value)
}

func (a *Any) UnmarshalJSON(data []byte) error {
	if a == nil {
		return errors.New("unmarshal json on nil pointer")
	}

	return jsonutil.Unmarshal(data, &a.value)
}

func (a *Any) Value() any {
	return a.value
}

func (a *Any) IsNil() bool {
	return a == nil || a.value == nil
}

func (a *Any) InitDefault() {
}

func (a *Any) String() string {
	if a == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Any(%+v)", *a)
}

func (a *Any) Decode(out any) error {
	if a == nil || a.value == nil {
		return errors.New("decode on nil pointer or nil value")
	}

	if out == nil {
		return errors.New("decode to nil pointer")
	}

	data, err := a.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal value to json error")
	}

	err = jsonutil.Unmarshal(data, out)
	if err != nil {
		return errors.Wrap(err, "unmarshal json to value error")
	}

	return nil
}

func GetValue[T any](v *Any) (T, error) {
	var result T

	if v.IsNil() {
		return result, nil
	}

	return result, v.Decode(result)
}
