package yeast

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test if yeast prepends an iterated seed when previous id is the same.
func TestPrependID(t *testing.T) {
	waitUntilNextMillisecond()
	y := New()

	id1 := y.Yeast()
	id2 := y.Yeast()
	id3 := y.Yeast()
	assert.NotContains(t, id1, ".")
	assert.Contains(t, id2, ".0")
	assert.Contains(t, id3, ".1")
}

// Test if yeast resets the seed.
func TestResetSeed(t *testing.T) {
	waitUntilNextMillisecond()
	y := New()

	id1 := y.Yeast()
	id2 := y.Yeast()
	id3 := y.Yeast()
	assert.NotContains(t, id1, ".")
	assert.Contains(t, id2, ".0")
	assert.Contains(t, id3, ".1")

	waitUntilNextMillisecond()

	id1 = y.Yeast()
	id2 = y.Yeast()
	id3 = y.Yeast()
	assert.NotContains(t, id1, ".")
	assert.Contains(t, id2, ".0")
	assert.Contains(t, id3, ".1")
}

// Test if yeast does not collide.
func TestAgainstCollision(t *testing.T) {
	waitUntilNextMillisecond()
	y := New()
	a := make([]string, 30000)

	for i := 0; i < len(a); i++ {
		a[i] = y.Yeast()
	}
	sort.Strings(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			t.Fatalf("found a duplicate entry, index: %d - %d, elements: '%s' and '%s'", i, i+1, a[i], a[i+1])
		}
	}
}

// Test if id can be converted to a timestamp.
func TestConvertIDToATimestamp(t *testing.T) {
	waitUntilNextMillisecond()
	y := New()
	now := time.Now().Unix()

	id := y.Yeast()
	assert.Equal(t, id, y.Encode(now))

	d, err := y.Decode(id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, now, d)
}

func waitUntilNextMillisecond() {
	now := time.Now().Unix()
	for now == time.Now().Unix() {
		// Do nothing
	}
}
