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

func TestCreateOutput(t *testing.T) {
	cases := []struct {
		in           string
		alphabetical bool
		want         string
	}{
		{"(test)", false, "\ntest"},
		{"(test1,test2(test3),test4)", false, "\ntest1\ntest2\n- test3\ntest4"},
		{"(test1,test2(test3),test4)", true, "\ntest1\ntest2\n- test3\ntest4"},
		{"(ctest1,atest2(test3),btest4)", true, "\natest2\n- test3\nbtest4\nctest1"},
		{"(test1,test2(test3(test4(test5,test6),test7),test8),test9,test10)", false, "\ntest1\ntest2\n- test3\n-- test4\n--- test5\n--- test6\n-- test7\n- test8\ntest9\ntest10"},
		{"(dtest1,btest2(btest3(atest4(btest5,atest6),btest7),atest8),atest9,ctest10)", true, "\natest9\nbtest2\n- atest8\n- btest3\n-- atest4\n--- atest6\n--- btest5\n-- btest7\nctest10\ndtest1"},
	}
	for _, c := range cases {
		tempItem, _ := CreateItem(c.in)
		got := createOutput(tempItem.chunks, c.alphabetical)
		if got != c.want {
			t.Errorf("createOutput failed.  In: %v, alphabetical: %v, Got: %v, Expected: %v", c.in, c.alphabetical, got, c.want)
		}
	}
}
