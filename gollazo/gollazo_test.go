package gollazo

import (
	"reflect"
	"testing"
)

func TestIsCollazoCipher(t *testing.T) {
	type args struct {
		cipher string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Check given example from programming challenge",
			args: args{"84581248O6096095854123337"},
			want: true,
		},
		{
			name: "Check given example from programming challenge with missing first character",
			args: args{"4581248O6096095854123337"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCollazoCipher(tt.args.cipher); got != tt.want {
				t.Errorf("IsCollazoCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractAB(t *testing.T) {
	type args struct {
		cipher string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:    "Check given example from programming challenge",
			args:    args{"84581248O6096095854123337"},
			want:    "5412333",
			want1:   "84581248O60960958",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := extractAB(tt.args.cipher)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractAB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractAB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("extractAB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
