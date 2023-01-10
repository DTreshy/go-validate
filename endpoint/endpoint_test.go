package endpoint_test

import (
	"testing"

	"github.com/DTreshy/go-validate/endpoint"
	"github.com/stretchr/testify/require"
)

var validateTestCases = []struct {
	name    string
	input   string
	wantErr bool
}{
	{
		name:    "valid IP address and port",
		input:   "192.0.2.1:80",
		wantErr: false,
	},
	{
		name:    "invalid IP address and valid port",
		input:   "300.0.2.1:80",
		wantErr: true,
	},
	{
		name:    "valid IP address and invalid port",
		input:   "192.0.2.1:808080",
		wantErr: true,
	},
	{
		name:    "valid domain name and port",
		input:   "example.com:80",
		wantErr: false,
	},
	{
		name:    "invalid domain name and valid port",
		input:   "invalid..domain:80",
		wantErr: true,
	},
	{
		name:    "valid domain name and invalid port",
		input:   "example.com:65536",
		wantErr: true,
	},
	{
		name:    "missing port",
		input:   "example.com",
		wantErr: true,
	},
	{
		name:    "valid domain name with hyphens and valid port",
		input:   "example-domain.com:80",
		wantErr: false,
	},
	{
		name:    "valid domain name with digits and valid port",
		input:   "example123.com:80",
		wantErr: false,
	},
	{
		name:    "empty endpoint",
		input:   "",
		wantErr: true,
	},
	{
		name:    "port only",
		input:   ":80",
		wantErr: true,
	},
	{
		name:    "maximum length domain label and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.com:80",
		wantErr: false,
	},
	{
		name:    "maximum length plus one domain label and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl.com:80",
		wantErr: true,
	},
	{
		name:    "maximum length domain and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk:80",
		wantErr: false,
	},
	{
		name:    "maximum length plus one domain and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl:80",
		wantErr: true,
	},
	{
		name:    "maximum length domain name and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk:80",
		wantErr: false,
	},
	{
		name:    "maximum length domain name plus one and valid port",
		input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijks:80",
		wantErr: true,
	},
	{
		name:    "domain name with invalid character and valid port",
		input:   "example$domain.com:80",
		wantErr: true,
	},
	{
		name:    "label with invalid rune and valid port",
		input:   "example\uFFFDdomain.com:80",
		wantErr: true,
	},
	{
		name:    "label with hyphen on the beggining and valid port",
		input:   "example.-com.invalid:80",
		wantErr: true,
	},
	{
		name:    "label with hyphen on the end and valid port",
		input:   "example.com-.invalid:80",
		wantErr: true,
	},
	{
		name:    "domain beginning with hyphen and valid port",
		input:   "-example:80",
		wantErr: true,
	},
	{
		name:    "domain ending with hyphen and valid port",
		input:   "example-:80",
		wantErr: true,
	},
	{
		name:    "domain beggining with digit and valid port",
		input:   "9example:80",
		wantErr: true,
	},
	{
		name:    "domain ending with a dot and valid port",
		input:   "example.:80",
		wantErr: true,
	},
	{
		name:    "domain without dot with valid port",
		input:   "example:80",
		wantErr: false,
	},
	{
		name:    "domain without dot with invalid port",
		input:   "example:-1",
		wantErr: true,
	},
}

func TestValidate(t *testing.T) {
	for _, tt := range validateTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := endpoint.Validate(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inputNum := i % len(validateTestCases)
		endpoint.Validate(validateTestCases[inputNum].input)
	}
}
