package encoder

import (
	"reflect"
	"testing"
)

func TestMorse(t *testing.T) {
	tests := []struct {
		nameTest string
		name     string
		want     string
	}{
		{
			nameTest: "Morse",
			name:     "abc",
			want:     ".--....-.-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			if got := CodeMorse(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cannotparse() = %v, want %v", got, tt.want)
			}
		})
	}
}
