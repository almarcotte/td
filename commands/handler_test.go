package commands

import (
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	args := []string{"one", "two", "three"}
	handler := NewHandler(args)

	if !reflect.DeepEqual(args, handler.Args) {
		t.Fatalf("Expected handler arguments to be %+v, got %+v", args, handler.Args)
	}
}
