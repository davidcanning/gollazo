package gollazo

import (
	"reflect"
	"testing"
)

func TestCheckCipher(t *testing.T) {
	type args struct {
		cipher string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name:    "Check given example from programming challenge",
			args:    args{"84581248O6096095854123337"},
			want:    7,
			want1:   "5412333",
			want2:   "84581248O60960958",
			wantErr: false,
		},
		{
			name:    "Check given example from programming challenge with missing first character",
			args:    args{"4581248O6096095854123337"},
			want:    -1,
			want1:   " ",
			want2:   " ",
			wantErr: true,
		},
		{
			name:    "One digit U toy example",
			args:    args{"8743O2396854126543216"},
			want:    6,
			want1:   "654321",
			want2:   "8743O239685412",
			wantErr: false,
		},
		{
			name:    "Two digit U toy example",
			args:    args{"11223344556677881110111213777777777777713"},
			want:    13,
			want1:   "7777777777777",
			want2:   "11223344556677881110111213",
			wantErr: false,
		},
		{
			name:    "One digit U toy example invalid triplet",
			args:    args{"11O2334455667777776"},
			want:    -1,
			want1:   " ",
			want2:   " ",
			wantErr: true,
		},
		{
			name:    "Two digit U toy example invalid triplet",
			args:    args{"1122334455667788991011777777777777713"},
			want:    -1,
			want1:   " ",
			want2:   " ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := CheckCipher(tt.args.cipher)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckCipher() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckCipher() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckCipher() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_splitAtoIntArray(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Check against given example",
			args:    args{"5412333"},
			want:    []int{5, 4, 1, 2, 3, 3, 3},
			wantErr: false,
		},
		{
			name:    "Check for non-integer in string",
			args:    args{"5412A33"},
			want:    []int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitAtoIntArray(tt.args.A)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitAtoIntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Check against given example",
			args:    args{"84581248O60960958"},
			want:    []string{"84", "58", "12", "48", "O60", "960", "958"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitBtoStrArray(tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitBtoStrArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBtoStrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
