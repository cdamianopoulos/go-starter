package utl

import (
	"fmt"
	"strconv"
	"strings"
)

// HostPort joins a host string and port number that can be accepted as an address.
func HostPort(host string, port uint16) string {
	if port == 0 {
		return host + ":"
	}

	return host + ":" + strconv.FormatUint(uint64(port), 10)
}

// Sprint prints a []string in a comma separated human-readable string.
func Sprint(s []string) string {
	switch {
	case s == nil:
		return "<nil>"
	case len(s) == 0:
		return "[]"
	default:
		return fmt.Sprintf("[`%s`]", strings.Join(s, "`, `"))
	}
}
