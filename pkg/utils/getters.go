package utils

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetString(d *schema.ResourceData, key string) string {
	switch v := d.Get(key).(type) {
	case *string:
		return *v
	case string:
		return v
	default:
		panic("cannot GetString on a non string type")
	}
}

func GetStringP(d *schema.ResourceData, key string) *string {
	s := GetString(d, key)
	if reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface()) {
		return nil
	}
	return &s
}

func GetBool(d *schema.ResourceData, key string) bool {
	switch v := d.Get(key).(type) {
	case *bool:
		return *v
	case bool:
		return v
	default:
		panic("cannot GetString on a non string type")
	}
}

func GetBoolP(d *schema.ResourceData, key string) *bool {
	s := GetBool(d, key)
	return &s
}

func GetFloat64(d *schema.ResourceData, key string) float64 {
	switch v := d.Get(key).(type) {
	case *float64:
		return *v
	case float64:
		return v
	default:
		panic("cannot GetFloat on a non float type")
	}
}

func GetFloat64P(d *schema.ResourceData, key string) *float64 {
	s := GetFloat64(d, key)
	return &s
}

func GetInt(d *schema.ResourceData, key string) int {
	switch v := d.Get(key).(type) {
	case *int:
		return *v
	case int:
		return v
	default:
		panic("cannot GetInt on a non int type")
	}
}

func GetIntP(d *schema.ResourceData, key string) *int {
	s := GetInt(d, key)
	return &s
}

func GetBlob(d *schema.ResourceData, key string) map[string]interface{} {
	var blob map[string]interface{}

	switch b := d.Get(key).(type) {
	case []interface{}:
		// we're a list of blobs, not what we want.  so
		// let's take the first element of the list
		blob = b[0].(map[string]interface{})
	case map[string]interface{}:
		blob = b
	}

	return blob
}
