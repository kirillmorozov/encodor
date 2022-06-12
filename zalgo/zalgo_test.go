package zalgo

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/kirillmorozov/encodor/utils"
)

func TestEncode(t *testing.T) {
	type args struct {
		reader     io.Reader
		diacritics int8
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "Hashtag are not encoded",
			args:       args{reader: strings.NewReader("#hashtag"), diacritics: 1},
			wantWriter: "#hashtag",
			wantErr:    false,
		},
		{
			name:       "Usernames are not encoded",
			args:       args{reader: strings.NewReader("@username"), diacritics: 1},
			wantWriter: "@username",
			wantErr:    false,
		},
		{
			name:       "Diacritics < minDiacritics",
			args:       args{reader: strings.NewReader("henlo"), diacritics: 0},
			wantWriter: "",
			wantErr:    true,
		},
		{
			name:       "Diacritics > maxDiacritics",
			args:       args{reader: strings.NewReader("henlo"), diacritics: 6},
			wantWriter: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := Encode(tt.args.reader, writer, tt.args.diacritics); (err != nil) != tt.wantErr {
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
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := bytes.NewReader(bench.Text)
				var buf bytes.Buffer
				_ = Encode(r, &buf, 5)
			}
		})
	}
}

func FuzzEncode(f *testing.F) {
	f.Add("zalgo", int8(5))
	f.Fuzz(func(t *testing.T, input string, diacritics int8) {
		r := strings.NewReader(input)
		var w bytes.Buffer
		err := Encode(r, &w, diacritics)
		if err != nil && w.String() != "" {
			t.Errorf("%q, %v", err, w.String())
		}
	})
}
