package validate

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	MaxLabelLength  int = 63
	MaxDomainLength int = 253
	MaxPort         int = 65535
)

func Endpoint(address string) error {
	idx := strings.LastIndex(address, ":")
	if idx == -1 {
		return fmt.Errorf("no colon in address: %v", address)
	}

	err := PortString(address[idx+1:])
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}

	return Hostname(address[:idx])
}

func Hostname(hostname string) error {
	err := CombineErrors(
		Domain(hostname),
		IP(hostname),
	)

	if err != nil {
		return fmt.Errorf("invalid hostname: %w", err)
	}

	return nil
}

func IP(address string) error {
	if len(address) < 2 {
		return fmt.Errorf("too short for valid IPv4 or IPv6 format: %v", address)
	}

	if address[0] == '[' && address[len(address)-1] == ']' {
		address = address[1 : len(address)-1]
	}

	ip := net.ParseIP(address)
	if ip == nil {
		return fmt.Errorf("should be valid IPv4 or IPv6 format: %v", address)
	}

	return nil
}

func Domain(name string) error {
	if name == "" {
		return fmt.Errorf("domain can't be empty")
	}

	hostname := []byte(name)

	if len(hostname) > MaxDomainLength {
		return fmt.Errorf("byte length of hostname '%s' is %d, can't exeed %d", name, len(hostname), MaxDomainLength)
	}

	if hostname[len(hostname)-1] == '.' {
		return fmt.Errorf("hostname '%s' ends with trailing dot", name)
	}

	for len(hostname) > 0 {
		var (
			label []byte
			err   error
		)

		label, hostname, err = nextLabel(hostname)
		if err != nil {
			return err
		}

		if len(label) == 0 {
			return fmt.Errorf("invalid zero-length label in hostname: %s", name)
		}

		if len(label) > MaxLabelLength {
			return fmt.Errorf("byte length of label '%s' is %d, can't exceed %d", label, len(label), MaxLabelLength)
		}

		if label[0] == '-' {
			return fmt.Errorf("label '%s' begins with a hyphen", label)
		}

		if label[len(label)-1] == '-' {
			return fmt.Errorf("label '%s' ends with a hyphen", label)
		}
	}

	return nil
}

func PortString(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("port is not a number")
	}

	return Port(portNum)
}

func Port(port int) error {
	if port < 1 {
		return fmt.Errorf("port number must be positive")
	}

	if port > MaxPort {
		return fmt.Errorf("port number can't exceed %d", MaxPort)
	}

	return nil
}

func nextLabel(address []byte) (label, remaining []byte, err error) {
	for i, b := range address {
		if b == '.' {
			return address[:i], address[i+1:], nil
		}

		if !(isLetter(b) || isDigit(b) || b == '-') {
			c, _ := utf8.DecodeRuneInString(string(address[i:]))
			if c == utf8.RuneError {
				return nil, address, fmt.Errorf("invalid rune at offset %d", i)
			}

			return nil, address, fmt.Errorf("invalid character '%c' at offset %d", c, i)
		}
	}

	return address, nil, nil
}

func isDigit(digit byte) bool {
	return digit >= '0' && digit <= '9'
}

func isLetter(letter byte) bool {
	return (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z')
}
