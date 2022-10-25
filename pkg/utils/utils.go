package utils

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
