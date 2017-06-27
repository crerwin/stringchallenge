package item_test

import (
	"testing"

	"github.com/crerwin/stringchallenge/item"
)

// Test CreateItem with valid inputs
func TestCreateItem(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"(id,created,employee(id,firstname,employeeType(id),lastname),location)"},
	}
	for _, c := range cases {
		got, err := item.CreateItem(c.in)
		if err != nil {
			t.Errorf("createItem(%v) failed.  Error: %v", c.in, err.Error())
		}
		if got.GetRawText() != c.in {
			t.Errorf("createItem(%v) failed.  Got: %v, Expected: %v", c.in, got.GetRawText(), c.in)
		}
	}
}

// Test CreateItem with invalid inputs (must throw error)
func TestCreateItemInvalid(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"(id,created,employee(id,firstname,employeeType(id),lastname,location)"},
		{""},
		{"()()"},
		{"(()"},
	}
	for _, c := range cases {
		_, err := item.CreateItem(c.in)
		if err.Error() != "Invalid Input" {
			t.Errorf("CreateItem(%v) failed.  Expected Invalid Input error.", c.in)
		}
	}
}

// Test GetRawText (after creating item with valid input)
func TestGetRawText(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"()", "()"},
		{"(id,created,employee(id,firstname,employeeType(id),lastname),location)", "(id,created,employee(id,firstname,employeeType(id),lastname),location)"},
	}
	for _, c := range cases {
		tempItem, _ := item.CreateItem(c.in)
		got := tempItem.GetRawText()
		if got != c.want {
			t.Errorf("GetRawText failed.  Got: %v, Wanted: %v", got, c.want)
		}
	}
}

func TestGetOutput(t *testing.T) {
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
		tempItem, _ := item.CreateItem(c.in)
		got := tempItem.GetOutput(c.alphabetical)
		if got != c.want {
			t.Errorf("createOutput failed.  In: %v, alphabetical: %v, Got: %v, Expected: %v", c.in, c.alphabetical, got, c.want)
		}
	}
}
