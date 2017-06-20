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

func validateText(text string) bool {
	return true
}
