package stringutil

func reverseTwo (input string) string {
	rev := []rune(input)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev [i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}