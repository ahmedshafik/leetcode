
func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}

	matching := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			expectedOpener := matching[char]
			if top == expectedOpener {
				stack = stack[:len(stack)-1]
			} else {
				return false

			}
		}

	}
	if len(stack) > 0 {
		return false
	}
	return true
}