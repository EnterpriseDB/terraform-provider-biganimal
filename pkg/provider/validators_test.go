package provider

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func Test_validateProjectId(t *testing.T) {
	v := "prj_abcdABCD01234567"
	if diags := validateProjectId(v, *new(cty.Path)); len(diags) != 0 {
		t.Fatalf("%q should be prj_abcdABCD01234567", v)
	}

	bad_names := []string{
		"BAD",                    // Obviously silly
		"Prj_something",          // Capital letter
		"prj_ab^dABCD01234567",   // Symbol
		"prj_abcdABCD0123",       // Too short
		"prj_abcdABCD0123456789", // Too long
		"p_abcdABCD01234567",     // Doesn't start with prj_
		"prjabcdABCD01234567",    // No `_` after prj
	}

	for _, tt := range bad_names {
		if diags := validateProjectId(tt, *new(cty.Path)); len(diags) == 0 {
			t.Fatalf("%q should NOT be prj_abcdABCD01234567", v)
		}
	}

}
