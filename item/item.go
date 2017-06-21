package item

import (
	"errors"
	"strings"

	"github.com/crerwin/stringchallenge/chunk"
)

// Item is our object for this exercise
type Item struct {
	rawText string
	chunks  []chunk.Chunk
}

// CreateItem is a constructor.  Returns an Item loaded with your input
func CreateItem(input string) (Item, error) {
	if validateText(input) {
		return Item{rawText: input}, nil
	}
	return Item{}, errors.New("Invalid Input")
}

// GetRawText returns the raw text from the Item.  By this point it has been
// validated
func (i *Item) GetRawText() string {
	return i.rawText
}

// extractAllChunks passes rawText to extractChunks and assigns the return slice
// to chunks.  All the work is done in extractChunks.
func (i *Item) extractAllChunks() {
	i.chunks = extractChunks(i.rawText, 0)
}

// extractChunks iterates through an input string and separates the chunks
// (words).  If it hits a (, it calls itself, attaching what's
// received back as the children of the last chunk in the list it's building.
//  I know, I KNOW, recursion bad, but it works really well here.
func extractChunks(input string, depth int) []chunk.Chunk {
	start := 0               // tracks beginning of each chunk
	var chunks []chunk.Chunk // temporary slice of chunks
	for i, char := range input {
		if char == ',' {
			// if we hit a comma, add the word to chunks
			chunks = append(chunks, chunk.CreateChunk(input[start:i], depth))
			start = i
		} else if char == '(' && depth > 0 {
			chunks[len(chunks)-1].AddChildren(extractChunks(input[i:], depth+1))
		} else if char == ')' {
			break
		}
	}
	return chunks
}

func (i *Item) GetConvertedText() string {
	return convertText(i.rawText, false)
}

func (i *Item) GetConvertedTextAlphabetical() string {
	return convertText(i.rawText, true)
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

func convertText(input string, alphabetical bool) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		switch r {
		case ',', '(', ')':
			return true
		}
		return false
	})
	return words[1]
}
