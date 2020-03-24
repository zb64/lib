package lib

import (
	"strings"
	"testing"
)

func TestDeflate_EncodeAndDecode(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		wantErr bool
	}{
		{"Empty", "", false},
		{"Single char", "A", false},
		{"Words", "Hello World", false},
		{"More words", "Strength does not come from winning. Your struggles develop your strengths.", false},
		{"Repeated words", strings.Repeat("1234567890", 1024), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := Deflate.Encode([]byte(tt.raw))
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			decoded, err := Deflate.Decode(encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.raw != string(decoded) {
				t.Errorf("Decode() gotResult = %v, want %v", string(decoded), tt.raw)
			}
		})
	}
}
