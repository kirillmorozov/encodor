package utils

import (
	"testing"
)

func Test_isSpecialWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Hashtag", args: args{"#nofilter"}, want: true},
		{name: "Username", args: args{"@username"}, want: true},
		{name: "Regular word", args: args{"word"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSpecialWord(tt.args.word); got != tt.want {
				t.Errorf("isSpecialWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
