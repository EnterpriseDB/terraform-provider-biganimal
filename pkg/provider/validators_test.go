package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"reflect"
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

func Test_validateARN(t *testing.T) {
	tests := []struct {
		name string
		arn  string
		want diag.Diagnostics
	}{
		{
			name: "invalid ARN",
			arn:  "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "invalid arn",
				Detail:   fmt.Sprintf("%v is a invalid aws role arn", "arn:aws:rds:eu-west-1:123456789012:db:mysql-db"),
			}},
		},
		{
			name: "valid ARN",
			arn:  "arn:aws:iam::123456789012:role/biganimal-role",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateARN(tt.arn, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateARN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateUUID(t *testing.T) {
	tests := []struct {
		name string
		uuid string
		want diag.Diagnostics
	}{
		{
			name: "invalid uuid",
			uuid: "xyz",
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "invalid uuid",
				Detail:   fmt.Sprintf("%v is a invalid uuid", "xyz"),
			}},
		},
		{
			name: "invalid uuid",
			uuid: "2c276c84-b756-11ed-afa1-0242ac120002",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateUUID(tt.uuid, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
