package test

func ReverseString(s string) string {
	left, right := 0, len(s)-1

	str := []byte(s)
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return string(str)
}
