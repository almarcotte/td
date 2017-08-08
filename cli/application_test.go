package cli

import (
	"runtime"
	"strings"
	"testing"
)

func TestNewApplication(t *testing.T) {
	// Just to simplify things
	if runtime.GOOS != "darwin" {
		t.SkipNow()
	}

	app := NewApplication()

	if !strings.HasPrefix(app.GlobalFile, "/Users/") && !strings.Contains(app.GlobalFile, "/home/") {
		t.Fatalf("Expected global file path to be an OS X / Unix homedir path, got %s", app.GlobalFile)
	}

	if !strings.HasSuffix(app.GlobalFile, ".td") {
		t.Fatalf("Expected global file path to end in `.td`, got %s", app.GlobalFile)
	}
}
