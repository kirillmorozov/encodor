package utils

import (
	"reflect"
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
			word := []byte(tt.args.word)
			if got := IsSpecialWord(word); got != tt.want {
				t.Errorf("isSpecialWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseLetters(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Empty buf",
			args: args{buf: []byte("")},
			want: []byte(""),
		},
		{
			name: "Single letter",
			args: args{buf: []byte("a")},
			want: []byte("a"),
		},
		{
			name: "Two letters",
			args: args{buf: []byte("ab")},
			want: []byte("ba"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLetters(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseLines(t *testing.T) {
	type args struct {
		slice [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Empty slice",
			args: args{slice: [][]byte{}},
			want: [][]byte{},
		},
		{
			name: "Single line",
			args: args{slice: [][]byte{[]byte("abc")}},
			want: [][]byte{[]byte("abc")},
		},
		{
			name: "Two lines",
			args: args{slice: [][]byte{[]byte("abc"), []byte("xyz")}},
			want: [][]byte{[]byte("xyz"), []byte("abc")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLines(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
