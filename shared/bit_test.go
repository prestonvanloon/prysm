package shared

import (
	"bytes"
	"testing"
)

func TestCheckBit(t *testing.T) {
	tests := []struct {
		a []byte
		b int
		c bool
	}{
		{a: []byte{200}, b: 4, c: true},  //11001000
		{a: []byte{148}, b: 5, c: true},  //10010100
		{a: []byte{146}, b: 4, c: false}, //10010010
		{a: []byte{179}, b: 7, c: true},  //10110011
		{a: []byte{49}, b: 6, c: false},  //00110001
	}
	for _, tt := range tests {
		set := CheckBit(tt.a, tt.b)
		if set != tt.c {
			t.Errorf("Test check bit set failed with %v and location %v", tt.a, tt.b)
		}
	}
}

func TestBitSetCount(t *testing.T) {
	tests := []struct {
		a byte
		b int
	}{
		{a: 200, b: 3}, //11001000
		{a: 148, b: 3}, //10010100
		{a: 146, b: 3}, //10010010
		{a: 179, b: 5}, //10110011
		{a: 49, b: 3},  //00110001
	}
	for _, tt := range tests {
		if int(BitSetCount([]byte{tt.a})) != tt.b {
			t.Errorf("BitSetCount(%d) = %v, want = %d", tt.a, int(BitSetCount([]byte{tt.a})), tt.b)
		}
	}
}

func TestByteLength(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{a: 200, b: 25},     //11001000
		{a: 34324, b: 4291}, //10010100
		{a: 146, b: 19},     //10010010
		{a: 179, b: 23},     //10110011
		{a: 49, b: 7},       //00110001
	}
	for _, tt := range tests {
		if BitLength(tt.a) != tt.b {
			t.Errorf("BitLength(%d) = %d, want = %d", tt.a, BitLength(tt.a), tt.b)
		}
	}
}

func TestBitSet(t *testing.T) {
	tests := []struct {
		a int
		b []byte
	}{
		{a: 0, b: []byte{128}},    //10000000
		{a: 1, b: []byte{64}},     //01000000
		{a: 5, b: []byte{4}},      //00000100
		{a: 10, b: []byte{0, 32}}, //00000000 00100000
		{a: 100, b: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8}},
	}
	for _, tt := range tests {
		if !bytes.Equal(SetBitfield(tt.a), tt.b) {
			t.Errorf("SetBitfield(%v) = %d, want = %v", tt.a, SetBitfield(tt.a), tt.b)
		}
	}
}
