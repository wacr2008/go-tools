package jsonWebToken

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestJwt_CreateToken(t *testing.T) {
	priKeyData, _ := ioutil.ReadFile("./rsa_private_key.pem")
	jsonWebToken, err := New(nil, priKeyData, nil)
	if err != nil {
		panic(err)
	}
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				data: map[string]interface{}{"name": "admin", "phone": "13300220033", "role": "admin"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonWebToken.CreateToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwt.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("Jwt.CreateToken() token = %v \n", got)
		})
	}
}

func TestJwt_ParseToken(t *testing.T) {
	priKeyData, _ := ioutil.ReadFile("./rsa_private_key.pem")
	pubKeyData, _ := ioutil.ReadFile("./rsa_public_key.pem")
	jsonWebToken, err := New(nil, priKeyData, pubKeyData)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{"name": "admin", "phone": "13300220033", "role": "admin"}
	token, err := jsonWebToken.CreateToken(data)
	if err != nil {
		panic(err)
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				tokenString: token,
			},
			want:    data,
			wantErr: false,
		}, {
			name: "format error",
			args: args{
				tokenString: token + "a",
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "is expired",
			args: args{
				tokenString: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjp7Im5hbWUiOiJhZG1pbiIsInBob25lIjoiMTMzMDAyMjAwMzMiLCJyb2xlIjoiYWRtaW4ifSwiZXhwIjoxNTYyMDIwMTgyLCJpc3MiOiJnby10b29scyIsIm5iZiI6MTU2MjAxNTU4Mn0.jEt7Lw-pgN7MwLz1q-JGCZtD8TkJPvA83f9BJxIaR8PD7d1HcAyNxkoLf10USMWUg_tay556DrUa-zBmdl68faJR8Fp40eCYMWk46ee724ZZtly0sa6xXaqC_mwM4IUKJ9QAZVvjpqa39Yccin18k5gb-DZM7b5gTJu-UohGWaI",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonWebToken.ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwt.ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Jwt.ParseToken() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Jwt.ParseToken() claims = %v\n", got)
			}
		})
	}
}

func TestJwt_RefreshToken(t *testing.T) {
	priKeyData, _ := ioutil.ReadFile("./rsa_private_key.pem")
	pubKeyData, _ := ioutil.ReadFile("./rsa_public_key.pem")
	jsonWebToken, err := New(nil, priKeyData, pubKeyData)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{"name": "admin", "phone": "13300220033", "role": "admin"}
	token, err := jsonWebToken.CreateToken(data)
	if err != nil {
		panic(err)
	}

	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				tokenString: token,
			},
			wantErr: false,
		}, {
			name: "format error",
			args: args{
				tokenString: token + "a",
			},
			wantErr: true,
		}, {
			name: "is expired",
			args: args{
				tokenString: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjp7Im5hbWUiOiJhZG1pbiIsInBob25lIjoiMTMzMDAyMjAwMzMiLCJyb2xlIjoiYWRtaW4ifSwiZXhwIjoxNTYyMDIwMTgyLCJpc3MiOiJnby10b29scyIsIm5iZiI6MTU2MjAxNTU4Mn0.jEt7Lw-pgN7MwLz1q-JGCZtD8TkJPvA83f9BJxIaR8PD7d1HcAyNxkoLf10USMWUg_tay556DrUa-zBmdl68faJR8Fp40eCYMWk46ee724ZZtly0sa6xXaqC_mwM4IUKJ9QAZVvjpqa39Yccin18k5gb-DZM7b5gTJu-UohGWaI",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonWebToken.RefreshToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwt.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("Jwt.CreateToken() token = %v \n", got)
		})
	}
}
