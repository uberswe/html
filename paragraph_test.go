package html

import (
	"strings"
	"testing"
)

func TestParagraphRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "Hello world!"
	expected3 := "</div>"

	p := Paragraph().Content("Hello world!").Class("test")

	if !strings.Contains(p.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", p.Render(), expected1)
	}

	if !strings.Contains(p.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", p.Render(), expected2)
	}

	if !strings.Contains(p.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", p.Render(), expected3)
	}
}
