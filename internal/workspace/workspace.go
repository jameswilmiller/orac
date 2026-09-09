package workspace

import (
	"context"
	"fmt"
	"os"
)

// Workspace is a directory a single trial may freely dirty
type Workspace interface {
	Dir() string
	Close() error
}

// Factory produces fresh workspaces, one per trial
type Factory interface {
	New(ctx context.Context) (Workspace, error)
}

// InPlace the room is the project folder
type InPlace struct {
	Source string
}

var _ Factory = InPlace{}

func (f InPlace) New(_ context.Context) (Workspace, error) {
	return inPlaceWS{dir: f.Source}, nil
}

type inPlaceWS struct {
	dir string
}

func (w inPlaceWS) Dir() string  { return w.dir }
func (w inPlaceWS) Close() error { return nil }

// TempCopy copies the project into a temp directory per trial, so
// one trial can never leak state into the next. the safe default.
type TempCopy struct {
	Source string
}

var _ Factory = TempCopy{}

func (f TempCopy) New(ctx context.Context) (Workspace, error) {
	dir, err := os.MkdirTemp("", "orac--")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}

	if err := os.CopyFS(dir, os.DirFS(f.Source)); err != nil {
		_ = os.RemoveAll(dir)
		return nil, fmt.Errorf("copy %s into temp dir: %w", f.Source, err)
	}

	return tempWS{dir: dir}, nil
}

type tempWS struct {
	dir string
}

func (w tempWS) Dir() string  { return w.dir }
func (w tempWS) Close() error { return os.RemoveAll(w.dir) }
