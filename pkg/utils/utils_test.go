package utils

import (
	"reflect"
	"testing"
)

func TestCopyObjectJson(t *testing.T) {
	t.Parallel()

	testValue := "testMe"

	testSucceedIn := map[string]interface{}{
		"test": testValue,
	}

	testSucceedOut := struct {
		Test string `json:"test"`
	}{}

	testFailIn := map[string]interface{}{
		"test": testValue,
	}

	testFailOut := struct {
		Test string `json:"wrong_annotation"`
	}{}

	type args struct {
		in  any
		out any
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		expected   interface{}
		expectFail bool
	}{
		{
			name:       "Expect succeed",
			wantErr:    false,
			expectFail: false,
			args: args{
				in:  testSucceedIn,
				out: &testSucceedOut,
			},
			expected: &struct {
				Test string `json:"test"`
			}{
				Test: testValue,
			},
		},
		{
			name:       "Expect fail",
			wantErr:    false,
			expectFail: true,
			args: args{
				in:  testFailIn,
				out: &testFailOut,
			},
			expected: &struct {
				Test string `json:"wrong_annotation"`
			}{
				Test: testValue,
			},
		},
		{
			name:    "Expect error",
			wantErr: true,
			args: args{
				in:  struct{}{},
				out: struct{}{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := CopyObjectJson(tt.args.in, tt.args.out)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CopyObjectJson() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				if tt.expectFail && reflect.DeepEqual(tt.args.out, tt.expected) {
					t.Fatalf("CopyObjectJson() expected fail but succeeded, got %v expected %v", tt.args.out, tt.expected)
				} else if !tt.expectFail && !reflect.DeepEqual(tt.args.out, tt.expected) {
					t.Fatalf("CopyObjectJson() expected success but failed, field(s) not mapped, got %v expected %v", tt.args.out, tt.expected)
				}
			}
		})
	}
}
