package html

import (
	"strings"
	"testing"
)

func TestTextAreaRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<label for=\"test\">Text area</label>"
	expected3 := "<textarea id=\"test\" name=\"test\" rows=\"10\" cols=\"50\">"
	expected4 := "TestField"
	expected5 := "</textarea>"
	expected6 := "</div"

	b := TextArea().Name("test").Content("TestField").Label("Text area").Id("test").Class("test").Cols(50).Rows(10)

	if !strings.Contains(b.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected1)
	}

	if !strings.Contains(b.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected2)
	}

	if !strings.Contains(b.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected3)
	}

	if !strings.Contains(b.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected4)
	}

	if !strings.Contains(b.Render(), expected5) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected5)
	}

	if !strings.Contains(b.Render(), expected6) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected6)
	}
}
