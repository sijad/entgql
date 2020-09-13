package generator

import (
	"os"
	"testing"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
)

func TestVarType(t *testing.T) {
	entSchemaPath := "./testdata/basic/ent/schema"
	graph, err := entc.LoadGraph(entSchemaPath, &gen.Config{})
	if err != nil {
		t.Error(err)
	}
	if err := SchemaDefinition(os.Stdout, graph); err != nil {
		t.Error(t)
	}
	if err := Models(os.Stdout, graph); err != nil {
		t.Error(t)
	}
}
