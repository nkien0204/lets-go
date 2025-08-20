package generator_test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectTreeMap(t *testing.T) {
	treeMap := generator.GetProjectTreeMap()
	
	// Verify the tree map is not nil
	assert.NotNil(t, treeMap)
	
	// Verify it contains expected top-level keys based on the architecture
	assert.Contains(t, treeMap, "cmd")
	assert.Contains(t, treeMap, "internal")
	assert.Contains(t, treeMap, "Makefile.tmpl")
	assert.Contains(t, treeMap, "main.go.tmpl")
	assert.Contains(t, treeMap, "go.mod.tmpl")
	
	// Verify internal structure exists
	internal, ok := treeMap["internal"].(map[string]any)
	assert.True(t, ok, "internal should be a map")
	assert.Contains(t, internal, "delivery")
	assert.Contains(t, internal, "usecase")
	assert.Contains(t, internal, "repository")
	assert.Contains(t, internal, "domain")
	
	// Verify domain structure exists
	domain, ok := internal["domain"].(map[string]any)
	assert.True(t, ok, "domain should be a map")
	assert.Contains(t, domain, "entity")
	
	// Verify the structure is consistent
	assert.IsType(t, map[string]any{}, treeMap)
}