package utils

import "reflect"

func StringRef(s string) *string {
	return &s
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func F64Ref(f float64) *float64 {
	return &f
}

func DerefF64(f *float64) float64 {
	if f != nil {
		return *f
	}

	return 0.0
}

// IsSlicey returns true if the argument is either a slice, or an array
// we need to treat some configuration blocks as a list (pg_config)
// and some configuration blocks as a map (architecture)
func IsSlicey(o any) bool {
	return reflect.TypeOf(o).Kind() == reflect.Slice || reflect.TypeOf(o).Kind() == reflect.Array
}
