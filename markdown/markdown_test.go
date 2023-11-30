package markdown_test

import (
	"testing"

	"github.com/manyids2/go-tufte/markdown"
)

// TestDocument calls NewDoc on README.md, checking
// for a valid return value.
func TestDocument(t *testing.T) {
	path := "../README.md"

	doc, err := markdown.NewDocument(path)
	if err != nil {
		t.Fatalf(`NewDocument("%s") failed: %s `, path, err)
	}

	if doc.Path != path {
		t.Fatalf(`doc.Path failed: %s, %s ; %s `, doc.Path, path, err)
	}
}
