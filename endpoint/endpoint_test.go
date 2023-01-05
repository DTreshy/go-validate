package endpoint_test

import (
	"testing"

	"github.com/DTreshy/go-validate/endpoint"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
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
			name:    "maximum length domain name",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.com:80",
			wantErr: false,
		},
		{
			name:    "maximum length plus one domain name and valid port",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl.com:80",
			wantErr: true,
		},
		{
			name:    "domain name with invalid character and valid port",
			input:   "example$domain.com:80",
			wantErr: true,
		},
		{
			name:    "domain name with invalid label and valid port",
			input:   "example.com..invalid:80",
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

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := endpoint.Validate(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
