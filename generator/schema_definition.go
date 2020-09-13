package generator

import (
	"io"

	"github.com/facebook/ent/entc/gen"
)

func (g *Generator) SchemaDefinition(wr io.Writer) error {
	if err := templates.ExecuteTemplate(wr, "template/base_sdl.tmpl", nil); err != nil {
		return err
	}
	for _, t := range g.Graph.Nodes {
		if err := nodeSchemaDefinition(wr, t); err != nil {
			return err
		}
	}
	return nil
}

func nodeSchemaDefinition(wr io.Writer, t *gen.Type) error {
	return templates.ExecuteTemplate(wr, "template/node_sdl.tmpl", t)
}
