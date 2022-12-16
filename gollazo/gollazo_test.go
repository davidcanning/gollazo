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

func Test_sumIntArray(t *testing.T) {
	type args struct {
		int_arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Check against simple example",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumIntArray(tt.args.int_arr); got != tt.want {
				t.Errorf("sumIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateABPair2Plaintext(t *testing.T) {
	type args struct {
		A_elem      int
		B_elem      string
		private_key []int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Check against given example character 1",
			args:    args{5, "84", []int{24, 22, 12}},
			want:    "W",
			wantErr: false,
		},
		{
			name:    "Check against given example character 2",
			args:    args{4, "58", []int{24, 22, 12}},
			want:    "H",
			wantErr: false,
		},
		{
			name:    "Check against given example character 3",
			args:    args{1, "12", []int{24, 22, 12}},
			want:    "A",
			wantErr: false,
		},
		{
			name:    "Check against given example character 4",
			args:    args{2, "48", []int{24, 22, 12}},
			want:    "T",
			wantErr: false,
		},
		{
			name:    "Check against given example character 5",
			args:    args{3, "O60", []int{24, 22, 12}},
			want:    "S",
			wantErr: false,
		},
		{
			name:    "Check against given example character 6",
			args:    args{3, "960", []int{24, 22, 12}},
			want:    "U",
			wantErr: false,
		},
		{
			name:    "Check against given example character 7",
			args:    args{3, "958", []int{24, 22, 12}},
			want:    "P",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translateABPair2Plaintext(tt.args.A_elem, tt.args.B_elem, tt.args.private_key)
			if (err != nil) != tt.wantErr {
				t.Errorf("translateABPair2Plaintext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("translateABPair2Plaintext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumElementwiseProduct(t *testing.T) {
	type args struct {
		int_arr     []int
		private_key []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Toy example",
			args:    args{[]int{1, 2, 3}, []int{1, 2, 3}},
			want:    14,
			wantErr: false,
		},
		{
			name:    "Toy example error",
			args:    args{[]int{1, 2, 3, 4}, []int{1, 2, 3}},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumElementwiseProduct(tt.args.int_arr, tt.args.private_key)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumElementwiseProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sumElementwiseProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
