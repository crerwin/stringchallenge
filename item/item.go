package item

type item struct {
	rawText string
}

func createItem(input string) item {
	validateText(input)
	return item{rawText: input}
}

func (i *item) GetRawText() string {
	return i.rawText
}

// validateText returns false if the given text does not fit the criteria as
// follows: 1) the given string must start with an open paren and end with a
// close paren, and the initial open paren must not be closed until the end (ie
// all text must fall within the outer parens).  2) all subsequent parens must
// be closed at some point.
func validateText(text string) bool {
	depth := 0
	length := len(text)
	// Since all valid strings must be within at least one level of parens, the
	// minimum lenght of a valid string is 2 "()"
	if length < 2 {
		return false
	}
	for i, char := range text {
		if char == '(' {
			depth++
		} else if char == ')' {
			depth--
		}
		// since all of the string must fall within the outer parens, our depth
		// should never drop below 1 until we get to the end
		if depth < 1 && i < length-1 {
			return false
		}
	}
	// at the end, the depth must be 0.  This could instead be 'if depth == 0
	// return true' and have the last line return false, but I like to hit all
	// false cases first, then have true as the default at the bottom.  If the
	// input has survived the gauntlet, then it is worthy.
	if depth != 0 {
		return false
	}
	return true
}
