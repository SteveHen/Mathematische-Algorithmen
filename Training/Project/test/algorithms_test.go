package test

import (
	"github.com/stretchr/testify/assert"

	"github.com/falo2/ma/input"

	"testing"
)

// Test the MooreBellmanFord algorithm with the Wege1.txt
// graph and the source and target nodes 0 and 4 respectively.
func TestMooreBellmanFord(t *testing.T) {
	g, _, _ := input.Read("Wege1.txt", true, false)
	_, pathWeight, _ := g.MooreBellmanFordLegacy(2, 0, false)

	assert.Equal(t, 6.0, pathWeight)
}
