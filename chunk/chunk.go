package chunk

// Chunk is a word from the input item, split on , ( and ), but contains other
// information needed to fulfill the requirements.
// Truffle Shuffle not implemented.
type Chunk struct {
	Value    string
	Depth    int
	Children []Chunk
}

// CreateChunk returns a new Chunk with the given values
func CreateChunk(value string, depth int) Chunk {
	return Chunk{Value: value, Depth: depth}
}

// AddChildren appends a slice of chunks to children
func (c *Chunk) AddChildren(children []Chunk) {
	c.Children = append(c.Children, children...)
}
