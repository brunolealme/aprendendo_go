package morestrings

func ReverseRunes(s string) string {
	reverse := []rune(s)

	// 3 var					 ; enquanto
	for i, j := 0, len(reverse)-1; i < len(reverse)/2; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}
