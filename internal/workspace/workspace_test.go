package workspace

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestInPlaceUsesSourceDirectory(t *testing.T) {
	f := InPlace{Source: "/some/path"}

	ws, err := f.New(context.Background())
	if err != nil {
		t.Fatalf("New returned an error: %s", err)
	}

	if ws.Dir() != "/some/path" {
		t.Errorf("Dir() = %q, want %q", ws.Dir(), "/some/path")
	}

	if err := ws.Close(); err != nil {
		t.Errorf("Close returned an error: %s", err)
	}
}

func TestTempCopyUsesTempDirectory(t *testing.T) {
	src := t.TempDir()

	if err := os.WriteFile(filepath.Join(src, "hello.txt"), []byte("hi"), 0o644); err != nil {
		t.Fatalf("could not write test file %v", err)
	}

	ws, err := TempCopy{Source: src}.New(context.Background())
	if err != nil {
		t.Fatalf("New returned an error: %s", err)
	}
	dir := ws.Dir()

	data, err := os.ReadFile(filepath.Join(dir, "hello.txt"))

	if err != nil {
		t.Fatalf("copied file not found %v", err)
	}

	if string(data) != "hi" {
		t.Errorf("copied file contains %q, want %q", data, "hi")
	}

	if err := ws.Close(); err != nil {
		t.Fatalf("Close returned an error: %s", err)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("folder %q still exists after close", dir)
	}
}
