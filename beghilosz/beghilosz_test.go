package beghilosz

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kirillmorozov/encodor/utils"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "Basic",
			input:      "beghilosz",
			wantWriter: "250714638",
		},
		{
			name:       "Two words",
			input:      "two words",
			wantWriter: "5DR0W 0WT",
		},
		{
			name:       "Hashtags untouched, but capitalized",
			input:      "#hashtag",
			wantWriter: "#HASHTAG",
		},
		{
			name:       "Usernames untouched, but capitalized",
			input:      "@username",
			wantWriter: "@USERNAME",
		},
		{
			name:       "Word order in a sentence with hashtag",
			input:      "Sentence with a #hashtag",
			wantWriter: "#HASHTAG A 4T1W 3CN3TN35",
		},
		{
			name:       "Multiline text",
			input:      "Line 1\nLine 2",
			wantWriter: "2 3N17\n1 3N17",
		},
		{
			name:       "Multiline special text",
			input:      "#hashtag\n@username",
			wantWriter: "@USERNAME\n#HASHTAG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			writer := &bytes.Buffer{}
			if err := Encode(reader, writer); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Encode() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for _, bench := range utils.EncodeBenchmarks {
		r := strings.NewReader(bench.Text)
		w := bytes.NewBuffer(make([]byte, len(bench.Text)))
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Encode(r, w)
			}
		})
	}
}
