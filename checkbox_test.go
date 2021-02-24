package html

import (
	"strings"
	"testing"
)

func TestCheckboxRender(t *testing.T) {
	expected1 := "<input type=\"checkbox\" id=\"test\" name=\"test\" value=\"test\" checked>"
	expected2 := "<div class=\"test\">"
	expected3 := "<label for=\"test\">Test</label>"
	expected4 := "</div>"

	c := Checkbox().Class("test").Content("Test").Name("test").Id("test").Value("test").Check()

	if !strings.Contains(c.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", c.Render(), expected1)
	}

	if !strings.Contains(c.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", c.Render(), expected2)
	}

	if !strings.Contains(c.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", c.Render(), expected3)
	}

	if !strings.Contains(c.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", c.Render(), expected4)
	}
}
