package generator

import (
	"io"

	"github.com/facebook/ent/entc/gen"
)

func SchemaDefinition(wr io.Writer, g *gen.Graph) error {
	if err := templates.ExecuteTemplate(wr, "template/base_sdl.tmpl", nil); err != nil {
		return err
	}
	for _, t := range g.Nodes {
		if err := NodeSchemaDefinition(wr, t); err != nil {
			return err
		}
	}
	return nil
}

func NodeSchemaDefinition(wr io.Writer, t *gen.Type) error {
	return templates.ExecuteTemplate(wr, "template/node_sdl.tmpl", t)
}
