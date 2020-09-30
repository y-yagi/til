package main

func wordSquares(words []string) [][]string {
	var results [][]string

	if len(words) == 0 {
		return results
	} else if len(words) == 1 {
		results = append(results, []string{words[0]})
		return results
	}

	mPrefixArr := make(map[string][]string)

	// save all words by prefix
	for i := range words[0] {
		for _, w := range words {
			mPrefixArr[w[0:i+1]] = append(mPrefixArr[w[0:i+1]], w)
		}
	}

	for _, w := range words {
		var wss []string
		wss = append(wss, w)

		appendWordBySequence(wss, mPrefixArr, &results)
	}

	return results
}

func appendWordBySequence(wss []string, mPrefixArr map[string][]string, results *[][]string) {
	var prefix string
	idx := len(wss)
	// extract perticular column string for the purpose of prefix matching
	for _, ws := range wss {
		prefix += string(ws[idx])
	}

	// exist word array that matching the prefix case
	if seqs, ok := mPrefixArr[prefix]; ok {
		for _, s := range seqs {
			nwss := append([]string(nil), wss...)
			nwss = append(nwss, s)

			if len(nwss) == len(nwss[0]) {
				*results = append(*results, nwss)
				continue
			}

			// keep searching next word sequence
			appendWordBySequence(nwss, mPrefixArr, results)
		}
	}
}
