package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	str := []rune(input)
	for i := len(str) - 1; i >= 0; i-- {
		output += string(str[i])
	}
	return output
}
