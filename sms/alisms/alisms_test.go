package alisms

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/liuchonglin/go-utils"
)

var aliSmsClient Sms

func init() {
	var err error
	aliSmsClient, err = NewAliSms()
	if err != nil {
		panic(err)
	}
}

func TestNewAliSms(t *testing.T) {
	tests := []struct {
		name    string
		want    *AliSms
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAliSms()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAliSms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAliSms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliSms_SendSmsCode(t *testing.T) {
	type args struct {
		phoneNumber string
		code        string
	}
	tests := []struct {
		name    string
		args    args
		want    *SendSmsResponse
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := aliSmsClient.SendSmsCode(tt.args.phoneNumber, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("AliSms.SendSmsCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AliSms.SendSmsCode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AliSms.SendSmsCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAliSms_GetSmsDetails(t *testing.T) {
	type fields struct {
		Client *sdk.Client
	}
	type args struct {
		phoneNumber string
		sendDate    string
		bizId       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *QuerySendDetailsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aliSmsClient.GetSmsDetails(tt.args.phoneNumber, tt.args.sendDate, tt.args.bizId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AliSms.GetSmsDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AliSms.GetSmsDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliSms_GetSmsDetailsList(t *testing.T) {
	type args struct {
		phoneNumber string
		sendDate    string
		bizId       string
		pageSize    int
		currentPage int
	}
	tests := []struct {
		name    string
		args    args
		want    *QuerySendDetailsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aliSmsClient.GetSmsDetailsList(tt.args.phoneNumber, tt.args.sendDate, tt.args.bizId, tt.args.pageSize, tt.args.currentPage)
			if (err != nil) != tt.wantErr {
				t.Errorf("AliSms.GetSmsDetailsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AliSms.GetSmsDetailsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliSms_SendSms(t *testing.T) {
	type args struct {
		phoneNumber   string
		signName      string
		templateCode  string
		templateParam string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				phoneNumber:   "13096403377",
				signName:      "亖堂小镇",
				templateCode:  "SMS_99430013",
				templateParam: "{\"code\":\"123456\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSuccess, resp, err := aliSmsClient.SendSms(tt.args.phoneNumber, tt.args.signName, tt.args.templateCode, tt.args.templateParam)
			if err != nil {
				t.Errorf("a.SendSms() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("send sms status=", isSuccess)
			utils.PrintlnJson(resp)
		})
	}
}

func TestMns(t *testing.T) {
	Mns()
}
