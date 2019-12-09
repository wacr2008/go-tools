package email

import (
	"testing"
)

func TestSendEmailCode(t *testing.T) {
	type args struct {
		to   string
		code string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				to:   "liulei5522@qq.com",
				code: "123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendEmailCode(tt.args.to, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendEmailCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
