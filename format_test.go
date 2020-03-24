package lib

import (
	"strings"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
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
	formats := []struct {
		name   string
		format *dataFormat
	}{
		{"Deflate", Deflate},
		{"LZW-LSB", LzwLSB},
		{"LZW-MSB", LzwMSB},
		{"Plain", Plain},
	}

	for _, tt := range tests {
		for _, ff := range formats {
			tName := tt.name + "_" + ff.name
			t.Run(tName, func(t *testing.T) {
				encoded, err := ff.format.Encode([]byte(tt.raw))
				if (err != nil) != tt.wantErr {
					t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				decoded, err := ff.format.Decode(encoded)
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
}
