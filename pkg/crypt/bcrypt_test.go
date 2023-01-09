package crypt

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "123456",
			args: args{
				hashedPassword: "$2a$10$5UgbrSkN2sAzUbvf362HReUjXGNE/rdeU0QrDDbb9f87s7CEQGp9a",
				password:       "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compare(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	fmt.Println(Encrypt("123456LYPqyi9wv64akuc5"))
}
