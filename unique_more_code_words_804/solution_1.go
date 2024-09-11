package unique_more_code_words_804

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--", "--..",
	}

	morseMap := make(map[string]bool)

	for _, word := range words {
		morse := ""
		for _, c := range word {
			morse += morseCodes[c-'a']
		}
		morseMap[morse] = true
	}

	return len(morseMap)

}
