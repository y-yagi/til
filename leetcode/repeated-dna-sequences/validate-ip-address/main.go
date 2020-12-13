package main

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		if validIPv4(IP) {
			return "IPv4"
		}
	} else {
		if validIPv6(IP) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validIPv4(IP string) bool {
	nums := strings.Split(IP, ".")
	if len(nums) != 4 {
		return false
	}
	for _, n := range nums {
		if tmp, err := strconv.Atoi(n); err != nil || tmp < 0 || tmp > 255 {
			return false
		} else if strconv.Itoa(tmp) != n {
			return false
		}
	}
	return true
}

func validIPv6(IP string) bool {
	nums := strings.Split(IP, ":")
	if len(nums) != 8 {
		return false
	}
	for _, n := range nums {
		if !valid(n) {
			return false
		}
	}
	return true
}

func valid(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
