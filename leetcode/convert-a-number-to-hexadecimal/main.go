package main

func toHex(num int) string {
	dict := map[uint64]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}
	hex := ""
	var n uint64

	if num == 0 {
		return "0"
	}

	if num < 0 {
		n = uint64(num + 4294967296)
	} else {
		n = uint64(num)
	}

	for n > 0 {
		d := n % 16
		n /= 16

		hex = dict[d] + hex
	}
	return hex
}
