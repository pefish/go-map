package go_map

import (
	"reflect"
	"testing"
)

func TestMapClass_Remain(t *testing.T) {
	type args struct {
		map_   map[string]interface{}
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut map[string]interface{}
	}{
		{
			name: `test`,
			args: args{
				map_: map[string]interface{}{
					`a`: 1,
					`b`: 2,
					`c`: 3,
				},
				fields: []string{`a`, `b`},
			},
			wantOut: map[string]interface{}{
				`a`: 1,
				`b`: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := Map.Remain(tt.args.map_, tt.args.fields); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("MapClass.Remain() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
