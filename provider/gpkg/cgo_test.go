// +build cgo

package gpkg

import (
	"github.com/go-spatial/tegola/dict"
	"testing"

	"github.com/DennisRutjes/tegola/provider"
)

// This is a test to just see that the init function is doing something.
func TestNewProviderStartup(t *testing.T) {
	_, err := NewTileProvider(dict.Dict{})
	if err == provider.ErrUnsupported {
		t.Fatalf("supported, expected any but unsupported got %v", err)
	}
}
