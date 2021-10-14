package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Birinci case",
			args: args{
				name: "Meltem",
			},
			want: "Hello, Meltem\n",
		},

		{
			name: "İkinci case",
			args: args{
				name: "Mert",
			},
			want: "Hello, Mert\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depart(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Birinci case",
			args: args{
				name: "Meltem",
			},
			want: "Goodbye, Meltem\n",
		},

		{
			name: "İkinci case",
			args: args{
				name: "Mert",
			},
			want: "Goodbye, Mert\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depart(tt.args.name); got != tt.want {
				t.Errorf("depart() = %v, want %v", got, tt.want)
			}
		})
	}
}
