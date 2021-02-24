package html

import (
	"strings"
	"testing"
)

func TestFileRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<label for=\"test\">Test</label>"
	expected3 := "<input type=\"file\" id=\"test\" name=\"test\" accept=\"image/png, image/jpeg\">"
	expected4 := "</div>"

	f := File().Class("test").Content("Test").Name("test").Id("test").Accept("image/png, image/jpeg")

	if !strings.Contains(f.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected1)
	}

	if !strings.Contains(f.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected2)
	}

	if !strings.Contains(f.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected3)
	}

	if !strings.Contains(f.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected4)
	}
}
