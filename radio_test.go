package html

import (
	"strings"
	"testing"
)

func TestRadioRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<input type=\"radio\" id=\"test\" name=\"test\" value=\"test\" checked>"
	expected3 := "<label for=\"test\">Test</label>"
	expected4 := "</div>"

	r := Radio().Label("Test").Id("test").Class("test").Value("test").Name("test").Check()

	if !strings.Contains(r.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", r.Render(), expected1)
	}

	if !strings.Contains(r.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", r.Render(), expected2)
	}

	if !strings.Contains(r.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", r.Render(), expected3)
	}

	if !strings.Contains(r.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", r.Render(), expected4)
	}
}
