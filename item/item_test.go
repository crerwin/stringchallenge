package item

import (
	"testing"
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
		{"(id,created,employee(id,firstname,employeeType(id), lastname),location)", true},
		{"(thing, thing2(thing3, thing4(thing5)))", true},
		{"(thing, thing2(thing3, thing4(thing5))", false},
		{"thing, thing2(thing3, thing4(thing5)))", false},
		{"()()(())", false},
		{"(()()(()))", true},
		{"()()(()", false},
		{")()(", false},
	}
	for _, c := range cases {
		got := validateText(c.in)
		if got != c.want {
			t.Errorf("validateText(%v) failed.  Got: %v, expected: %v", c.in, got, c.want)
		}
	}
}

// Test CreateItem with valid inputs
func TestCreateItem(t *testing.T) {
	cases := []struct {
		in   string
		want Item
	}{
		{"(id,created,employee(id,firstname,employeeType(id), lastname),location)", Item{rawText: "(id,created,employee(id,firstname,employeeType(id), lastname),location)"}},
	}
	for _, c := range cases {
		got, _ := CreateItem(c.in)
		if got != c.want {
			t.Errorf("createItem(%v) failed.  Got: %v, Expected: %v", c.in, got, c.want)
		}
	}
}

// Test CreateItem with invalid inputs (must throw error)
func TestCreateItemInvalid(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"(id,created,employee(id,firstname,employeeType(id), lastname,location)"},
		{""},
		{"()()"},
		{"(()"},
	}
	for _, c := range cases {
		_, err := CreateItem(c.in)
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
		{"(id,created,employee(id,firstname,employeeType(id), lastname),location)", "(id,created,employee(id,firstname,employeeType(id), lastname),location)"},
	}
	for _, c := range cases {
		tempItem, _ := CreateItem(c.in)
		got := tempItem.GetRawText()
		if got != c.want {
			t.Errorf("GetRawText failed.  Got: %v, Wanted: %v", got, c.want)
		}
	}
}
