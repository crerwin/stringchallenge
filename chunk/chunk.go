package chunk

// Chunk is a word from the input item, split on , ( and ), but contains other
// information needed to fulfill the requirements.
// Truffle Shuffle not implemented.
type Chunk struct {
	value    string
	depth    int
	children []Chunk
}

// CreateChunk returns a new Chunk with the given values
func CreateChunk(value string, depth int) Chunk {
	return Chunk{value: value, depth: depth}
}

// AddChildren appends a slice of chunks to children
func (c *Chunk) AddChildren(children []Chunk) {
	c.children = append(c.children, children...)
}

// GetValue returns the Chunk's value
func (c *Chunk) GetValue() string {
	return c.value
}

// GetDepth returns the Chunk's depth
func (c *Chunk) GetDepth() int {
	return c.depth
}
