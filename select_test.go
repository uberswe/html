package html

import (
	"strings"
	"testing"
)

func TestSelectRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<label for=\"test\">Test</label>"
	expected3 := "<select name=\"test\" id=\"test\">"
	expected4 := "<option value=\"option1\">Option 1</option>"
	expected5 := "<option value=\"option1\">Option 1</option>"
	expected6 := "</select>"
	expected7 := "</div>"

	option1 := Option().Label("Option 1").Value("option1")
	option2 := Option().Label("Option 1").Value("option1")

	s := Select().Class("test").Name("test").Label("Test").Id("test").Options(option1, option2)

	if !strings.Contains(s.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected1)
	}

	if !strings.Contains(s.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected2)
	}

	if !strings.Contains(s.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected3)
	}

	if !strings.Contains(s.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected4)
	}

	if !strings.Contains(s.Render(), expected5) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected5)
	}

	if !strings.Contains(s.Render(), expected6) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected6)
	}

	if !strings.Contains(s.Render(), expected7) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", s.Render(), expected7)
	}
}
