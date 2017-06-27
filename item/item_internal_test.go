package item

import (
	"testing"

	"github.com/crerwin/stringchallenge/chunk"
)

func TestValidateText(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"(", false},
		{"f", false},
		{"", false},
		{"()", true},
		{"(id,created,employee(id,firstname,employeeType(id),lastname),location)", true},
		{"(thing,thing2(thing3,thing4(thing5)))", true},
		{"(thing,thing2(thing3,thing4(thing5))", false},
		{"thing,thing2(thing3,thing4(thing5)))", false},
		{"()()(())", false},
		{"(()()(()))", true},
		{"()()(()", false},
		{")()(", false},
		{"(test1, test2)", false},
	}
	for _, c := range cases {
		got := validateText(c.in)
		if got != c.want {
			t.Errorf("validateText(%v) failed.  Got: %v, expected: %v", c.in, got, c.want)
		}
	}
}

func TestExtractChunks(t *testing.T) {
	cases := []struct {
		in   string
		want []chunk.Chunk
	}{
		{"(test)", []chunk.Chunk{chunk.Chunk{Value: "test", Depth: 1}}},
		{"(test,test2)", []chunk.Chunk{chunk.Chunk{Value: "test", Depth: 1}, chunk.Chunk{Value: "test2", Depth: 1}}},
		{"(test(test3))", []chunk.Chunk{chunk.Chunk{Value: "test", Depth: 1, Children: []chunk.Chunk{chunk.Chunk{Value: "test3", Depth: 2}}}}},
	}
	for _, c := range cases {
		got, _ := extractChunks(c.in, 1)
		if len(got) != len(c.want) {
			t.Errorf("extractChunks(%v, %v) failed.  Lengths different.  Got %v, Expected %v", c.in, 1, got, c.want)
		}
		for i := range got {
			if got[i].Value != c.want[i].Value {
				t.Errorf("extractChunks(%v, %v) failed.  Got %v, Expected %v", c.in, 1, got, c.want)
			}
		}
	}
}

// func TestCreateOutput(t *testing.T) {
// 	cases := []struct {
// 		in           string
// 		alphabetical bool
// 		want         string
// 	}{
// 		{"(test)", false, "test"},
// 		{"(test1,test2(test3),test4)", false, "test1\ntest2\n- test3\ntest4"},
// 		{"(test1,test2(test3),test4)", true, "test1\ntest2\n- test3\ntest4"},
// 		{"(ctest1,atest2(test3),btest4)", true, "atest2\n- test3\nbtest4\nctest1"},
// 	}
// 	for _, c := range cases {
// 		tempItem, _ := CreateItem(c.in)
// 		got := tempItem.createOutput(c.alphabetical)
// 		if got != c.want {
// 			t.Errorf("createOutput failed.  In: %v, alphabetical: %v, Got: %v, Expected: %v", c.in, c.alphabetical, got, c.want)
// 		}
// 	}
// }
