package validate

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const MaxPortNumber int = 65535
const MaxLabelLength int = 63
const MaxDomainLength int = 255

// Validate return nil if address have a valid domain or an ip address and a port, else it returns an error.
func Endpoint(address string) error {
	host := strings.Split(address, ":")
	if len(host) != 2 {
		return fmt.Errorf("invalid address format: %s", address)
	}

	port, err := strconv.Atoi(host[1])
	if err != nil {
		return fmt.Errorf("port %s is not a number", host[1])
	}

	if err = Port(port); err != nil {
		return err
	}

	isIPv4, _ := regexp.Match(`^\d+\.\d+\.\d+\.\d+$`, []byte(host[0]))
	if isIPv4 {
		return IPv4(host[0])
	}

	return Domain(host[0])
}

func Domain(name string) error {
	if name == "" {
		return fmt.Errorf("domain can't be empty")
	}

	if len(name) > 255 {
		return fmt.Errorf("domain length is %d, can't exceed %d", len(name), MaxDomainLength)
	}

	var l int

	for i, char := range name {
		if char == '.' {
			if i == l {
				return fmt.Errorf("invalid character '%c' at offset %d: label can't begin with a period", rune(char), i)
			}

			if i-l > MaxLabelLength {
				return fmt.Errorf("byte length of label '%s' is %d, can't exceed %d", name[l:i], i-l, MaxLabelLength)
			}

			if name[l] == '-' {
				return fmt.Errorf("label '%s' at offset %d begins with a hyphen", name[l:i], l)
			}

			if name[i-1] == '-' {
				return fmt.Errorf("label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}

			l = i + 1

			continue
		}

		if !(char >= 'a' && char <= 'z' || char >= '0' && char <= '9' || char == '-' || char >= 'A' && char <= 'Z') {
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("invalid rune at offset %d", i)
			}

			return fmt.Errorf("invalid character '%c' at offset %d", c, i)
		}
	}

	if l == len(name) {
		return fmt.Errorf("missing top level domain, domain can't end with a period")
	}

	if len(name)-l > MaxLabelLength {
		return fmt.Errorf("byte length of top level domain '%s' is %d, can't exceed %d", name[l:], len(name)-l, MaxLabelLength)
	}

	if name[l] == '-' {
		return fmt.Errorf("top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	}

	if name[len(name)-1] == '-' {
		return fmt.Errorf("top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	}

	return nil
}

func Port(port int) error {
	if port < 1 {
		return fmt.Errorf("port number is not positive: %d", port)
	}

	if port > MaxPortNumber {
		return fmt.Errorf("port number is %d, can't exeed %d", port, MaxPortNumber)
	}

	return nil
}

func IPv4(ip string) error {
	if r := net.ParseIP(ip); r == nil {
		return fmt.Errorf("invalid ip address: %s", ip)
	}

	return nil
}
