package main

func isPalindrome(s string) bool {
	bs := []rune{}
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			bs = append(bs, v)
		} else if v >= 'A' && v <= 'Z' {
			bs = append(bs, v+'a'-'A')
		} else if v >= '0' && v <= '9' {
			bs = append(bs, v)
		}
	}

	for i := 0; i < len(bs)/2; i++ {
		if bs[i] != bs[len(bs)-1-i] {
			return false
		}
	}

	return true
}
