package utils

import (
	"testing"
)

func TestWorker_NextId(t *testing.T) {
	w, err := NewWorker(0)
	if err != nil {
		t.Fatal("Create new worker error")
	}

	t.Logf("Created a new ID: %d", w.NextId())
}
