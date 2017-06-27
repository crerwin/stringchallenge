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

// func TestConvertText(t *testing.T) {
// 	cases := []struct {
// 		in           string
// 		want         string
// 		alphabetical bool
// 	}{
// 		{"(id,created,employee(id,firstname,employeeType(id), lastname),location)",
// 			"id\ncreated\nemployee\n- id\n- firstname\n- employeeType\n-- id\n- lastname\nlocation",
// 			false},
// 		{"(one, one, one(two, two(three)), one(two(three)))",
// 			"one\none\none\n- two\n- two\n-- three\none\n- two\n-- three",
// 			false},
// 		{"(id,created,employee(id,firstname,employeeType(id), lastname),location)",
// 			"created\nemployee\n- employeeType\n-- id\n- firstname\n- id\n- lastname\nid\nlocation",
// 			true},
// 		{"(one, one, one(two, two(three)), one(two(three)))",
// 			"one\none\none\n- two\n- two\n-- three\none\n- two\n-- three",
// 			true},
// 	}
// 	for _, c := range cases {
// 		got := convertText(c.in, c.alphabetical)
// 		if got != c.want {
// 			t.Errorf("convertText(%v) failed.  Got: %v, Expected: %v", c.in, got, c.want)
// 		}
// 	}
// }
