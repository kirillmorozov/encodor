package zalgo

import (
	"testing"

	"github.com/kirillmorozov/encodor/utils"
)

func TestEncode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hashtag are not encoded",
			args: args{text: "#hashtag"},
			want: "#hashtag",
		},
		{
			name: "Usernames are not encoded",
			args: args{text: "@username"},
			want: "@username",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.text); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for _, bench := range utils.EncodeBenchmarks {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Encode(bench.Text)
			}
		})
	}
}
