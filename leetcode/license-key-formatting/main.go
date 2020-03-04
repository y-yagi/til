package main

func licenseKeyFormatting(S string, K int) string {
	var res []byte
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] != '-' {
			if len(res)%(K+1) == K {
				res = append(res, '-')
			}
			res = append(res, toUpper(S[i]))
		}
	}
	reverse(res)
	return string(res)
}

func toUpper(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func licenseKeyFormatting_old(S string, K int) string {
	b := make([]byte, 0, len(S))

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] != '-' {
			if len(b)%(K+1) == K {
				b = append([]byte{'-'}, b...)
			}
			b = append([]byte{toUpper(S[i])}, b...)
		}
	}

	s := string(b)
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
	}

	return s
}
