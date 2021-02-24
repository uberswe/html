package html

import (
	"strings"
	"testing"
)

func TestFormRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<form method=\"POST\" action=\"/\">"
	expected3 := "<div class=\"test\">"
	expected4 := "<label for=\"test\">Test</label>"
	expected5 := "<input type=\"text\" id=\"test\" name=\"test\" value=\"\" placeholder=\"\">"
	expected6 := "<button class=\"button\">This is a button</button>"
	expected7 := "</form>"
	expected8 := "</div>"

	textField := Text().Label("Test").Class("test").Id("test").Name("test")

	button := Button().Class("button").Content("This is a button")

	f := Form().Class("test").Method("POST").Action("/").Fields(textField, button)

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

	if !strings.Contains(f.Render(), expected5) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected5)
	}

	if !strings.Contains(f.Render(), expected6) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected6)
	}

	if !strings.Contains(f.Render(), expected7) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected7)
	}

	if !strings.Contains(f.Render(), expected8) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected8)
	}
}
