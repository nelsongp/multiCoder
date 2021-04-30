package encoder

import (
	"reflect"
	"testing"
)

func TestMorse(t *testing.T){
	tests := []struct {
		name string
		want string
	}{
		{
			name: "testing morseParse",
			want: ".--....-.-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CodeMorse("abc"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cannotparse() = %v, want %v", got, tt.want)
			}
		})
	}
}