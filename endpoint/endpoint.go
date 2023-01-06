package endpoint

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Validate return nil if address have a valid domain or an ip address and a port, else it returns an error.
func Validate(address string) error {
	host := strings.Split(address, ":")
	if len(host) != 2 {
		return fmt.Errorf("invalid address format: %v", address)
	}

	isIP, err := regexp.Match(`^\d+\.\d+\.\d+\.\d+$`, []byte(host[0]))
	if err != nil {
		return err
	}

	if isIP {
		if r := net.ParseIP(host[0]); r == nil {
			return fmt.Errorf("invalid ip address: %v", host[0])
		}
	} else {
		if err := isValidDomain(host[0]); err != nil {
			return err
		}
	}

	port, err := strconv.Atoi(host[1])
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %v", host[1])
	}

	return nil
}

func isValidDomain(name string) error {
	switch {
	case name == "":
		return fmt.Errorf("domain can't be empty")
	case len(name) > 255:
		return fmt.Errorf("domain length is %d, can't exceed 255", len(name))
	}

	var l int

	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			switch {
			case i == l:
				return fmt.Errorf("invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}

			l = i + 1

			continue
		}

		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("invalid rune at offset %d", i)
			}

			return fmt.Errorf("invalid character '%c' at offset %d", c, i)
		}
	}

	switch {
	case l == len(name):
		return fmt.Errorf("missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("byte length of top level domain '%s' is %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}

	return nil
}
