package validate_test

import (
	"testing"

	"github.com/DTreshy/go-validate"
	"github.com/stretchr/testify/require"
)

func TestEndpoint(t *testing.T) {
	var endpointTestCases = []struct {
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
			name:    "invalid IP address as valid domain and valid port",
			input:   "300.0.2.1:80",
			wantErr: false,
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
			name:    "empty",
			input:   "",
			wantErr: true,
		},
		{
			name:    "port only",
			input:   ":80",
			wantErr: true,
		},
		{
			name:    "valid domain and invalid format port",
			input:   "example.com:invalid",
			wantErr: true,
		},
	}

	for _, tt := range endpointTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Endpoint(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestPortString(t *testing.T) {
	var portStringTestCases = []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid port",
			input:   "80",
			wantErr: false,
		},
		{
			name:    "too big port",
			input:   "808080",
			wantErr: true,
		},
		{
			name:    "zero port",
			input:   "0",
			wantErr: true,
		},
		{
			name:    "negative port",
			input:   "-1",
			wantErr: true,
		},
		{
			name:    "max port",
			input:   "65535",
			wantErr: false,
		},
		{
			name:    "max plus 1 port",
			input:   "65536",
			wantErr: true,
		},
		{
			name:    "not a number",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range portStringTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.PortString(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestPort(t *testing.T) {
	var portTestCases = []struct {
		name    string
		input   int
		wantErr bool
	}{
		{
			name:    "valid port",
			input:   80,
			wantErr: false,
		},
		{
			name:    "too big port",
			input:   808080,
			wantErr: true,
		},
		{
			name:    "zero port",
			input:   0,
			wantErr: true,
		},
		{
			name:    "negative port",
			input:   -1,
			wantErr: true,
		},
		{
			name:    "max port",
			input:   65535,
			wantErr: false,
		},
		{
			name:    "max plus 1 port",
			input:   65536,
			wantErr: true,
		},
	}

	for _, tt := range portTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Port(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestIP(t *testing.T) {
	var IPTestCases = []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid ip",
			input:   "192.0.2.1",
			wantErr: false,
		},
		{
			name:    "RFC2732 compatible IPv6 endpoint representation",
			input:   "[2001:db8::68]",
			wantErr: false,
		},
		{
			name:    "RFC2732 compatible IPv6 endpoint representation (for IPv4 address)",
			input:   "[::ffff:192.0.2.1]",
			wantErr: false,
		},
		{
			name:    "invalid ip",
			input:   "300.0.2.1",
			wantErr: true,
		},
		{
			name:    "too short format ip",
			input:   "192.0.2",
			wantErr: true,
		},
		{
			name:    "too long",
			input:   "192.0.2.1.9",
			wantErr: true,
		},
		{
			name:    "invalid format ip",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range IPTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.IP(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDomain(t *testing.T) {
	var DomainTestCases = []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid domain",
			input:   "example.com",
			wantErr: false,
		},
		{
			name:    "invalid ipv4 is valid domain",
			input:   "300.0.2.1",
			wantErr: false,
		},
		{
			name:    "domain beginning with a dot",
			input:   ".invalid.domain",
			wantErr: true,
		},
		{
			name:    "domain beginning with a hyphen",
			input:   "-invalid.domain",
			wantErr: true,
		},
		{
			name:    "domain beginning with a digit",
			input:   "9invalid.domain",
			wantErr: false,
		},
		{
			name:    "domain being a number",
			input:   "999",
			wantErr: false,
		},
		{
			name:    "domain name with hyphens",
			input:   "example-domain.com",
			wantErr: false,
		},
		{
			name:    "domain name with digits",
			input:   "example123.com",
			wantErr: false,
		},
		{
			name:    "empty domain",
			input:   "",
			wantErr: true,
		},
		{
			name:    "maximum length domain label",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.com",
			wantErr: false,
		},
		{
			name:    "maximum length plus one domain label",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl.com",
			wantErr: true,
		},
		{
			name:    "maximum length top level domain label",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk",
			wantErr: false,
		},
		{
			name:    "maximum length plus one top level domain label",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl",
			wantErr: true,
		},
		{
			name:    "maximum length domain name",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghi",
			wantErr: false,
		},
		{
			name:    "maximum length domain name plus one",
			input:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghij",
			wantErr: true,
		},
		{
			name:    "domain name with invalid character",
			input:   "example$domain.com",
			wantErr: true,
		},
		{
			name:    "label with invalid rune",
			input:   "example\uFFFDdomain.com",
			wantErr: true,
		},
		{
			name:    "label with hyphen on the beginning",
			input:   "example.-com.invalid",
			wantErr: true,
		},
		{
			name:    "label with hyphen on the end",
			input:   "example.com-.invalid",
			wantErr: true,
		},
		{
			name:    "domain ending with a hyphen",
			input:   "example-",
			wantErr: true,
		},
		{
			name:    "domain beginning with a hyphen",
			input:   "-example",
			wantErr: true,
		},
		{
			name:    "domain ending with a dot",
			input:   "example.",
			wantErr: true,
		},
	}

	for _, tt := range DomainTestCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Domain(tt.input)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
