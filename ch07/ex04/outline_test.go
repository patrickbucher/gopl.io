package outline

import (
	"testing"
)

func TestParseHTML(t *testing.T) {
	html := `
	<html>
		<head>
			<title>Title</title>
		</head>
		<body>
			<h1>Title</h1>
			<div>
				<p>Some <em>Content</em></p>
			</div>
		</body>
	</html>
	`
	tags := []string{"html", "head", "title", "body", "h1", "div", "p", "em"}
	stack, err := Parse(html)
	if err != nil {
		t.Errorf("parsing failed: %v", err)
	}
	if len(stack) != len(tags) {
		t.Errorf("html outline not parsed correctly, has %d items", len(stack))
	}
	for i := range stack {
		if stack[i] != tags[i] {
			t.Errorf("stack on level %d must contain %s but is %s", i, tags[i], stack[i])
		}
	}
}
