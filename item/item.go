package item

import "errors"

// Item is our object for this exercise
type Item struct {
	rawText string
}

// CreateItem is a constructor.  Returns an Item loaded with your input
func CreateItem(input string) (Item, error) {
	if validateText(input) {
		return Item{rawText: input}, nil
	} else {
		return Item{}, errors.New("Invalid Input")
	}
}

// GetRawText returns the raw text from the Item.  By this point it has been
// validated
func (i *Item) GetRawText() string {
	return i.rawText
}

// validateText returns false if the given text does not fit the criteria as
// follows: 1) the given string must start with an open paren and end with a
// close paren, and the initial open paren must not be closed until the end (ie
// all text must fall within the outer parens).  2) all subsequent parens must
// be closed at some point, meaning our final depth should be exactly 0.
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

func GetConvertedText() string {
	return ""
}

func GetConvertedTextAlphabetical() string {
	return ""
}
