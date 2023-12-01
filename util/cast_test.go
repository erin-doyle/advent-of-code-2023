package util

import (
	"testing"
)

func TestToInt(t *testing.T) {
	if got := ToInt("123"); got != 123 {
		t.Errorf("Int(123) = %v, want 123", got)
	}
	if got := ToInt("9835"); got != 9835 {
		t.Errorf("Int(9835) = %v, want 9835", got)
	}
}

func TestToString(t *testing.T) {
	byteTests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"byte", byte('a'), "a"},
		{"byte", byte('x'), "x"},
		{"int", 1234, "1234"},
		{"int", 512, "512"},
		{"rune", rune(65), "A"},
		{"rune", rune(97), "a"},
	}
	for _, tt := range byteTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.input); got != tt.want {
				t.Errorf("ToString(byte) = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToASCIIConstants(t *testing.T) {
	if ASCIICodeCapA != 65 {
		t.Errorf("Expected ASCIICodeCapA to be 65, got %d", ASCIICodeCapA)
	}
	if ASCIICodeLowerA != 97 {
		t.Errorf("Expected ASCIICodeLowerA to be 97, got %d", ASCIICodeLowerA)
	}
}

func TestToASCIICode(t *testing.T) {
	type args struct {
		arg interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example_string", args{"a"}, ASCIICodeLowerA},
		{"example_string", args{"b"}, ASCIICodeLowerA + 1},
		{"example_string", args{"z"}, ASCIICodeLowerA + 25},
		{"example_string", args{"C"}, ASCIICodeCapA + 2},
		{"example_rune", args{rune(97)}, 97},
		{"example_byte", args{'a'}, 97},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToASCIICode(tt.args.arg); got != tt.want {
				t.Errorf("ToASCIICode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestASCIIIntToChar(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{97}, "a"},
		{"example", args{98}, "b"},
		{"example", args{65}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ASCIIIntToChar(tt.args.code); got != tt.want {
				t.Errorf("ASCIIIntToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
