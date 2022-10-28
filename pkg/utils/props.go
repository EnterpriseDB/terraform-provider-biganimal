package utils

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/thoas/go-funk"
)

func NewPropList(o any) []interface{} {
	var (
		pl  = []interface{}{}
		err error
	)

	if !IsSlicey(o) {
		// we're not a list, so we can just build
		// a single map
		m := map[string]interface{}{}
		err = mapstructure.Decode(o, &m)
		pl = append(pl, m)
		if err != nil {
			panic(err)
		}

		return pl
	}

	// we're listy, but we don't know
	// what we're a list of.
	// let's iterate over the reflect.Value, rather than
	// trying to typeswitch and cast it to []interface{} or []KeyValue
	items := reflect.ValueOf(o)
	for i := 0; i < items.Len(); i++ {
		obj := items.Index(i).Interface()

		m := map[string]interface{}{}
		err = mapstructure.Decode(obj, &m)
		if err != nil {
			panic(err)
		}
		pl = append(pl, m)
	}

	return pl
}

func StructFromProps[S any](blobs any) (S, error) {
	lst, ok := blobs.([]interface{})
	var s S

	if !ok {
		return s, fmt.Errorf("wrong type of block, need list, got %T", blobs)
	}

	if IsSlicey(s) {
		err := mapstructure.Decode(lst, &s)
		return s, err
	}

	if len(lst) != 1 {
		return s, fmt.Errorf("%T needs exactly one block", s)
	}

	err := mapstructure.Decode(lst[0], &s)
	return s, err
}

func SetOrPanic(d *schema.ResourceData, key string, value interface{}) {
	if funk.IsEmpty(value) {
		return // empty value
	}

	var err error
	switch v := value.(type) {
	case []interface{}:
		err = d.Set(key, v)
	case int, float64, bool, string:
		err = d.Set(key, v)
	case *float64:
		err = d.Set(key, *v)
	case *int:
		err = d.Set(key, *v)
	case *bool:
		err = d.Set(key, *v)
	case *string:
		err = d.Set(key, DerefString(v))

	default:
		stringer, ok := value.(fmt.Stringer)
		if !ok {
			panic(fmt.Sprintf(" don't know how to handle %T", value))
		}

		err = d.Set(key, stringer.String())
	}

	// if d.Set fails
	if err != nil {
		panic(err)
	}
}
