package html

import (
	"strings"
	"testing"
)

func TestTextRender(t *testing.T) {
	expected1 := "<div class=\"test\">"
	expected2 := "<label for=\"test\">Test</label>"
	expected3 := "<input type=\"email\" id=\"test\" name=\"test\" value=\"test\" placeholder=\"markus@beubo.com\">"
	expected4 := "</div"

	b := Text().Value("test").Name("test").Label("Test").Type("email").Placeholder("markus@beubo.com").ID("test").Class("test")

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
}
