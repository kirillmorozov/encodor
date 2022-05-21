package beghilosz

import (
	"testing"

	"github.com/kirillmorozov/encodor/utils"
)

func TestEncode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Basic",
			args: args{"beghilosz"},
			want: "250714638",
		},
		{
			name: "Two words",
			args: args{"two words"},
			want: "5DR0W 0WT",
		},
		{
			name: "Hashtags untouched, but capitalized",
			args: args{"#hashtag"},
			want: "#HASHTAG",
		},
		{
			name: "Usernames untouched, but capitalized",
			args: args{"@username"},
			want: "@USERNAME",
		},
		{
			name: "Word order in a sentence with hashtag",
			args: args{"Sentence with a #hashtag"},
			want: "#HASHTAG A 4T1W 3CN3TN35",
		},
		{
			name: "Multiline text",
			args: args{"Line 1\nLine 2"},
			want: "2 3N17\n1 3N17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.input); got != tt.want {
				t.Errorf("Beghilosz() = %v, want %v", got, tt.want)
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
