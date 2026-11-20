package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMIMETypeIgnoresDetectedCharset(t *testing.T) {
	for _, text := range []string{"plain ASCII text\n", "Unicode text: café\n"} {
		filename := filepath.Join(t.TempDir(), "text.txt")
		if err := os.WriteFile(filename, []byte(text), 0600); err != nil {
			t.Fatal(err)
		}
		for _, tag := range []string{"mimetype=text/plain", "mimetype=text/*"} {
			if err := New().Var(filename, tag); err != nil {
				t.Errorf("%s rejected text file: %v", tag, err)
			}
		}
		if err := New().Var(filename, "mimetype=image/png"); err == nil {
			t.Error("text file accepted as image")
		}
	}
}

func TestMIMETypeMalformedDetectedType(t *testing.T) {
	if matchesMIMEType("text/plain; charset=", "text/plain") {
		t.Error("malformed detected type accepted")
	}
}
