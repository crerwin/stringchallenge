package item

import "testing"

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
