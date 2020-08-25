/**
 *
 * @author liangjf
 * @create on 2020/8/25
 * @version 1.0
 */
package utils

import (
	"reflect"
	"testing"
)

func TestDiffSlice(t *testing.T) {
	type args struct {
		old []string
		new []string
	}
	tests := []struct {
		name     string
		args     args
		wantDiff []string
	}{
		{
			name: "TestDiffSlice",
			args: args{
				old: []string{"1", "2", "3", "4"},
				new: []string{"1", "2", "3", "4", "5"},
			},
			wantDiff: []string{"5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := DiffSlice(tt.args.old, tt.args.new); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("DiffSlice() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
