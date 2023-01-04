package utils

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func StructFromProps[S any](blobs any) (S, error) {
	lst, ok := blobs.([]interface{})
	var s S

	if !ok {
		return s, fmt.Errorf("wrong type of block, need list, got %T", blobs)
	}

	if reflect.TypeOf(s).Kind() == reflect.Slice || reflect.TypeOf(s).Kind() == reflect.Array {
		err := mapstructure.Decode(lst, &s)
		return s, err
	}

	if len(lst) != 1 {
		return s, fmt.Errorf("%T needs exactly one block", s)
	}

	err := mapstructure.Decode(lst[0], &s)
	return s, err
}

func Set(d *schema.ResourceData, key string, value any) error {
	// there are a bunch of ways that data can be sent to this function
	// the data that can come in are
	// - primitive types
	// such as strings, floats, bools, ints, etc
	//
	// we can also get pointers to these primative types
	// eg, a *bool needs to be treated differently than a bool
	//
	// as well, we need to be able to set
	// structs
	// arrays, slices of either structs, or primitive types
	// and finally
	// we need to be able to translate structs into strings
	// eg, a pointintime struct in our API needs to be cast into a string
	// for the terraform schema
	//
	// so we do the following:

	// if it's a pointer we dereference it right away
	// references that get passed in will be confused by the
	// stringer check below, as stringer isn't implemented with
	// pointer receivers
	// we also make sure we don't try to dereference a zero value,
	// because that will break
	if reflect.ValueOf(value).Kind() == reflect.Ptr && !reflect.ValueOf(value).IsZero() {
		value = reflect.Indirect(reflect.ValueOf(value)).Interface()
	}

	// if it's a stringer, then stringify it
	if stringer, ok := value.(fmt.Stringer); ok {

		// and we can't stringify nil's.  so we skip that case
		if reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(stringer).IsNil() {
			value = ""
		} else {
			value = stringer.String()
		}

		// we can set the value here and return
		return d.Set(key, value)
	}

	switch reflect.ValueOf(value).Kind() {

	// in the case of a slice or an array, we need to
	// decode each element in the array and then set the
	// entire array at once
	case reflect.Slice, reflect.Array:
		pl := []interface{}{}
		items := reflect.ValueOf(value)
		for i := 0; i < items.Len(); i++ {
			obj := items.Index(i).Interface()

			var m interface{}
			err := mapstructure.Decode(obj, &m)
			if err != nil {
				return err
			}
			pl = append(pl, m)
		}
		value = pl

	// however if we're a simple structure then we
	// decode that scruct and set it.
	case reflect.Struct:
		m := map[string]interface{}{}
		pl := []interface{}{}

		err := mapstructure.Decode(value, &m)
		if err != nil {
			return err
		}

		pl = append(pl, m)
		value = pl
	}

	return d.Set(key, value)
}

func SetOrPanic(d *schema.ResourceData, key string, value any) {
	err := Set(d, key, value)

	if err != nil {
		panic(fmt.Sprintf("unable to set %s = %v (%T) : %v", key, value, value, err))
	}
}
