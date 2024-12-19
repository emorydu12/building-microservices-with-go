package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsKittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}

	kittens := store.Search("Garfield")

	assert.Equal(t, 1, len(kittens))
}

func TestReturnsKittenWhenSearchUnknown(t *testing.T) {
	store := MemoryStore{}

	kittens := store.Search("Unknown")

	assert.Equal(t, 0, len(kittens))
}
