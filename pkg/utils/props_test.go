package utils

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gotest.tools/v3/assert"
)

func stringRef(s string) *string {
	return &s
}

var testResource = &schema.Resource{
	Description: "Create a Postgres Cluster",

	CreateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	ReadContext:   func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	DeleteContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },

	Timeouts: &schema.ResourceTimeout{
		Create: schema.DefaultTimeout(45 * time.Minute),
	},

	Schema: map[string]*schema.Schema{
		"list": {
			Description: "a list of models",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"int": {
						Description: "int",
						Type:        schema.TypeInt,
						Required:    true,
					},
					"string": {
						Description: "string",
						Type:        schema.TypeString,
						Computed:    true,
					},
					"float": {
						Description: "float",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
				},
			},
		},
		"int": {
			Description: "int",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"float": {
			Description: "float",
			Type:        schema.TypeFloat,
			Optional:    true,
		},
		"string": {
			Description: "string",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"array": {
			Description: "array",
			Type:        schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
	},
}

type testModel struct {
	I int    `mapstructure:"int"`
	S string `mapstructure:"string"`

	// F is 'optional', so we may not always have this.  we have to omitempty
	// so we don't carry around unwanted default values as we transform these objects
	F float64 `mapstructure:"float,omitempty"`
}

type testStringable struct {
	I int    `mapstructure:"int"`
	S string `mapstructure:"string"`

	// F is 'optional', so we may not always have this.  we have to omitempty
	// so we don't carry around unwanted default values as we transform these objects
	F float64 `mapstructure:"float,omitempty"`
}

func (t testStringable) String() string {
	return t.S
}

type AllowedIpRange struct {
	CidrBlock   string `json:"cidrBlock" mapstructure:"cidr_block"`
	Description string `json:"description" mapstructure:"description"`
}

func TestStructFromProps(t *testing.T) {
	cr := testResource

	testCases := []struct {
		name    string
		in      []interface{}
		want    any
		wantErr bool
	}{
		{
			in: []interface{}{
				map[string]interface{}{
					"string": "string",
					"int":    1,
				},
			},
			want: testModel{
				I: 1,
				S: "string",
			},
			wantErr: false,
		},
		{
			in: []interface{}{
				map[string]interface{}{
					"string": "string",
					"int":    1,
					"float":  2.2,
				},
			},
			want: testModel{
				I: 1,
				S: "string",
				F: 2.2,
			},
			wantErr: false,
		},
	}

	for num, tcase := range testCases {
		t.Logf("testing StructFromProps #%d", num)
		config := map[string]interface{}{
			"list": tcase.in,
		}

		d := schema.TestResourceDataRaw(t, cr.Schema, config)

		a, err := StructFromProps[testModel](d.Get("list"))

		t.Log(err)
		assert.Equal(t, err != nil, tcase.wantErr)
		assert.DeepEqual(t, a, tcase.want)
	}
}

func TestSetOrPanic(t *testing.T) {
	cr := testResource

	testCases := []struct {
		name string
		kind string
		in   any
		out  any
	}{
		{ // simple values
			kind: "int",
			in:   int(1),
			out:  int(1),
		},
		{
			kind: "float",
			in:   1.1,
			out:  float64(1.1),
		},
		{
			kind: "string",
			in:   "string",
			out:  string("string"),
		},
		{ // string pointer
			kind: "string",
			in:   stringRef("randomstring"),
			out:  string("randomstring"),
		},
		{
			kind: "list",
			in:   testModel{I: 9, S: "string"},
			out:  []any{map[string]any{"int": int(9), "float": float64(0), "string": string("string")}},
		},
		{
			kind: "list",
			in:   &testModel{I: 9, S: "string"},
			out:  []any{map[string]any{"int": int(9), "float": float64(0), "string": string("string")}},
		},
		{
			kind: "list",
			in:   []testModel{{S: "hello", I: 1}},
			out:  []any{map[string]any{"float": float64(0), "int": int(1), "string": string("hello")}},
		},
		{
			kind: "list",
			in:   []testModel{},
			out:  []any{},
		},
		{ // stringables
			kind: "string",
			in:   &testStringable{S: "Hello"},
			out:  string("Hello"),
		},
		{ // stringables
			kind: "string",
			in:   testStringable{S: "Hello"},
			out:  string("Hello"),
		},
	}

	for num, tcase := range testCases {
		t.Logf("testing SetOrPanic #%d", num)
		config := map[string]interface{}{}

		d := schema.TestResourceDataRaw(t, cr.Schema, config)
		SetOrPanic(d, tcase.kind, tcase.in)

		out := d.Get(tcase.kind)
		assert.DeepEqual(t, out, tcase.out)
	}
}
