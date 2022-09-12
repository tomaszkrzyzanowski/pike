package pike

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCompare(t *testing.T) {
	type args struct {
		directory string
		arn       string
		init      bool
		exclude   *cli.StringSlice
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compare(tt.args.directory, tt.args.arn, tt.args.init, tt.args.exclude); (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompareIAMPolicy(t *testing.T) {
	type args struct {
		Policy    string
		OldPolicy string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"same",
			args{
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}"},
			true,
			false},
		{"different",
			args{
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:bogus\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}"},
			false,
			false},
		{"not-json",
			args{
				"guff",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:bogus\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}"},
			false,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareIAMPolicy(tt.args.Policy, tt.args.OldPolicy)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareIAMPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareIAMPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
