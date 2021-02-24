package html

import (
	"strings"
	"testing"
)

func TestTextRender(t *testing.T) {
	expected1 := "<button class=\"test2\">Test2</button>"
	expected2 := "<a href=\"https://beubo.com\">"

	b := Button().Class("test2").Content("Test2").Link("https://beubo.com")

	if !strings.Contains(b.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected1)
	}

	if !strings.Contains(b.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected2)
	}
}
