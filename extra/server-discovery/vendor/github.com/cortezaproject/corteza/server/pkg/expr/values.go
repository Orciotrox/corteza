package expr

import (
	"reflect"
)

type (
	Type interface {
		Type() string
		Cast(interface{}) (TypedValue, error)
	}

	TypedValue interface {
		Type
		Assigner
		Clone() (TypedValue, error)
		Get() interface{}
	}

	Comparable interface {
		Compare(TypedValue) (int, error)
	}

	typedValueWrap struct {
		Value interface{} `json:"@value"`
		Type  string      `json:"@type"`
	}

	Assigner interface {
		Assign(interface{}) error
	}

	FieldSelector interface {
		Has(k string) bool
		Select(k string) (TypedValue, error)
	}

	FieldAssigner interface {
		AssignFieldValue(string, TypedValue) error
	}

	DeepFieldAssigner interface {
		AssignFieldValue(Pather, TypedValue) error
	}

	Iterator interface {
		Each(func(k string, v TypedValue) error) error
	}

	TypeValueDecoder interface {
		Decode(dst reflect.Value) error
	}

	Dict interface {
		Dict() map[string]interface{}
	}

	Slice interface {
		Slice() []interface{}
	}
)

func UntypedValue(val interface{}) interface{} {
	if tv, is := val.(TypedValue); is {
		return tv.Get()
	}

	return val
}

func Must(v TypedValue, err error) TypedValue {
	if err != nil {
		panic(err)
	}
	return v
}
