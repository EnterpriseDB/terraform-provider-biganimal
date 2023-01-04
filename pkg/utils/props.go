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
	// if it's a pointer we dereference it right away
	// references that get passed in will be confused by the
	// stringer check below, as stringer isn't implemented with
	// pointer receivers
	if reflect.ValueOf(value).Kind() == reflect.Ptr && !reflect.ValueOf(value).IsZero() {
		value = reflect.Indirect(reflect.ValueOf(value)).Interface()
	}

	// if it's a stringer, then stringify it
	if stringer, ok := value.(fmt.Stringer); ok {
		value = stringer.String()
		err := d.Set(key, value)
		if err != nil {
			return err
		}
	}

	switch reflect.ValueOf(value).Kind() {
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
