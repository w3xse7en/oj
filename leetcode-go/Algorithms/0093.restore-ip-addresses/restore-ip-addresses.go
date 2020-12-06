package problem0093

import (
	"strconv"
	"strings"
)

var ips []string

func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	ips = make([]string, 0)
	loop(s, []string{})
	return ips
}

func loop(s string, ipaddr []string) {
	if len(ipaddr) > 3 {
		return
	}
	if len(ipaddr) == 3 && len(s) <= 3 {
		if len(s) > 1 && s[0] == '0' {
			return
		}
		sip, _ := strconv.Atoi(s)
		if sip <= 255 {
			ipaddr = append(ipaddr, s)
			ips = append(ips, strings.Join(ipaddr, "."))
		}
	}
	for i := 1; i <= 3 && i < len(s); i++ {
		sub := s[:i]
		if len(sub) > 1 && sub[0] == '0' {
			continue
		}
		sip, _ := strconv.Atoi(sub)
		if sip <= 255 {
			loop(s[i:], append(ipaddr, sub))
		}
	}
}
