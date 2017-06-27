package item

import (
	"errors"
	"sort"

	"github.com/crerwin/stringchallenge/chunk"
)

// Item is our object for this exercise
type Item struct {
	rawText string
	chunks  chunk.Chunks
}

// CreateItem is a constructor.  Returns an Item loaded with your input
func CreateItem(input string) (Item, error) {
	if validateText(input) {
		newItem := Item{rawText: input}
		newItem.extractAllChunks()
		return newItem, nil
	}
	return Item{}, errors.New("Invalid Input")
}

// GetRawText returns the raw text from the Item.  By this point it has been
// validated
func (i *Item) GetRawText() string {
	return i.rawText
}

// GetOutput returns the final output, not alphabetized.
func (i *Item) GetOutput() string {
	return createOutput(i.chunks, false)
}

// GetOutputAlphabetical returns the final output alphabetized.
func (i *Item) GetOutputAlphabetical() string {
	return createOutput(i.chunks, true)
}

// createOutput returns a string in the required format from a slice of chunks.
// It recursively calls itself if a chunk has children, and the final output
// bubbles up.
func createOutput(chunks chunk.Chunks, alphabetical bool) string {
	finalstring := ""
	if alphabetical {
		sort.Sort(chunks)
	}
	for i := range chunks {
		finalstring += "\n"
		for j := 0; j < chunks[i].Depth; j++ {
			finalstring += "-"
		}
		if chunks[i].Depth > 0 {
			finalstring += " "
		}
		finalstring += chunks[i].Value
		if len(chunks[i].Children) > 0 {
			finalstring += createOutput(chunks[i].Children, alphabetical)
		}
	}
	return finalstring
}

// extractAllChunks passes rawText to extractChunks and assigns the return slice
// to chunks.  All the work is done in extractChunks.
func (i *Item) extractAllChunks() {
	i.chunks, _ = extractChunks(i.rawText, 0)
}

// extractChunks iterates through an input string and separates the chunks
// (words).  If it hits a (, it calls itself, attaching what's
// received back as the children of the last chunk in the list it's building.
//  I know, I KNOW, recursion bad, but it works here.
func extractChunks(input string, depth int) (chunk.Chunks, int) {
	start := 1 // tracks beginning of each chunk
	end := 1
	var chunks chunk.Chunks // temporary slice of chunks
	// Using this instead of a nice i, char in range loop because I need to move
	// i under certain conditions.  We've already validated the text so we'll
	// skip the first (
	for i := 1; i < len(input); i++ {
		char := input[i]
		if char == ',' {
			// if we hit a comma, add the word to chunks
			if input[start] != ')' {
				chunks = append(chunks, chunk.CreateChunk(input[start:i], depth))
			}
			start = i + 1
		} else if char == '(' {
			chunks = append(chunks, chunk.CreateChunk(input[start:i], depth))
			tempchunks, newi := extractChunks(input[i:], depth+1)
			i += newi
			chunks[len(chunks)-1].AddChildren(tempchunks)
			start = i
			continue
		} else if char == ')' && input[start] != ')' {
			chunks = append(chunks, chunk.CreateChunk(input[start:i], depth))
			end = i
			break
		}
	}
	return chunks, end
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
		if char == ' ' {
			return false
		} else if char == '(' {
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
