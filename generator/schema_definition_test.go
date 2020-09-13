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
	generator := &Generator{Graph: graph}
	if err != nil {
		t.Error(err)
	}
	if err := generator.SchemaDefinition(os.Stdout); err != nil {
		t.Error(t)
	}
	if err := generator.Models(os.Stdout); err != nil {
		t.Error(t)
	}
	if err := generator.Resolvers(os.Stdout); err != nil {
		t.Error(t)
	}
}
