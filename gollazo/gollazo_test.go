package gollazo

import (
	"reflect"
	"testing"
)

func Test_splitAtoIntArray(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Check against given example",
			args: args{"5412333"},
			want: []int{5, 4, 1, 2, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitAtoIntArray(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitAtoIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitBtoStrArray(t *testing.T) {
	type args struct {
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Check against given example",
			args: args{"84581248O60960958"},
			want: []string{"84", "58", "12", "48", "O60", "960", "958"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitBtoStrArray(tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBtoStrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateAB2Roman(t *testing.T) {
	type args struct {
		num_numerals int
		sum_decimal  string
		private_key  []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check against given example",
			args: args{5, "84", []int{24, 22, 12}},
			want: "XXIII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateAB2Roman(tt.args.num_numerals, tt.args.sum_decimal, tt.args.private_key); got != tt.want {
				t.Errorf("translateAB2Roman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCipher(t *testing.T) {
	type args struct {
		cipher string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 string
		want3 string
	}{
		{
			name:  "Check given example from programming challenge",
			args:  args{"84581248O6096095854123337"},
			want:  true,
			want1: 7,
			want2: "5412333",
			want3: "84581248O60960958",
		},
		{
			name:  "Check given example from programming challenge with missing first character",
			args:  args{"4581248O6096095854123337"},
			want:  false,
			want1: -1,
			want2: " ",
			want3: " ",
		},
		{
			name:  "One digit U toy example",
			args:  args{"8743O2396854126543216"},
			want:  true,
			want1: 6,
			want2: "654321",
			want3: "8743O239685412",
		},
		{
			name:  "Two digit U toy example",
			args:  args{"11223344556677881110111213777777777777713"},
			want:  true,
			want1: 13,
			want2: "7777777777777",
			want3: "11223344556677881110111213",
		},
		{
			name:  "One digit U toy example invalid triplet",
			args:  args{"11O2334455667777776"},
			want:  false,
			want1: -1,
			want2: " ",
			want3: " ",
		},
		{
			name:  "Two digit U toy example invalid triplet",
			args:  args{"1122334455667788991011777777777777713"},
			want:  false,
			want1: -1,
			want2: " ",
			want3: " ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := CheckCipher(tt.args.cipher)
			if got != tt.want {
				t.Errorf("CheckCipher() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckCipher() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckCipher() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("CheckCipher() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
