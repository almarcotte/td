package todo

import (
	"runtime"
	"strings"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	// Just to simplify things
	if runtime.GOOS != "darwin" {
		t.SkipNow()
	}

	config := NewConfiguration()

	if !strings.HasPrefix(config.GlobalFile, "/Users/") && !strings.Contains(config.GlobalFile, "/home/") {
		t.Fatalf("Expected global file path to be an OS X / Unix homedir path, got %s", config.GlobalFile)
	}

	if !strings.HasSuffix(config.GlobalFile, ".td") {
		t.Fatalf("Expected global file path to end in `.td`, got %s", config.GlobalFile)
	}
}
