package release

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinned(t *testing.T) {
	client := New()
	pinned, err := client.Pinned()
	assert.NoError(t, err)

	t.Logf("pinnged release %+v", pinned)
	assert.True(t, pinned.Pinned)
}

func TestLatest(t *testing.T) {
	client := New()
	latest, err := client.Latest(true)
	assert.NoError(t, err)
	t.Logf("stable latest %+v", latest)

	latest, err = client.Latest(false)
	assert.NoError(t, err)
	t.Logf("unstable latest %+v", latest)
}

func TestList(t *testing.T) {
	client := New()
	releases, err := client.List(1)
	assert.NoError(t, err)
	for i, release := range releases {
		t.Logf("%d: %+v", i+1, release)
	}
}
