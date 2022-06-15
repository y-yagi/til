package main

func canConstruct(ransomNote string, magazine string) bool {
	dict := make([]int, 26)

	for i := 0; i < len(magazine); i++ {
		dict[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if dict[ransomNote[i]-'a'] == 0 {
			return false
		}
		dict[ransomNote[i]-'a']--
	}
	return true
}
