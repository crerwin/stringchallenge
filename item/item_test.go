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
