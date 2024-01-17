package stack

// BracketMatch 括号匹配
func BracketMatch(str string) bool {
	stack := NewSimpleStack()
	for _, i := range str {
		if i == '(' || i == '[' || i == '{' {
			stack.Push(i)
		} else {
			pop, err := stack.Pop()
			if err != nil {
				return false
			}
			if !matches(pop, i) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func matches(open any, close rune) bool {
	switch open.(rune) {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	}
	return false
}
