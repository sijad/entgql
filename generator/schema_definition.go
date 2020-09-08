package generator

import (
	"io"

	"github.com/facebook/ent/entc/gen"
)

func SchemaDefinition(wr io.Writer, g *gen.Graph) error {
	for _, t := range g.Nodes {
		if err := templates.ExecuteTemplate(wr, "template/node_sdl.tmpl", t); err != nil {
			return err
		}
	}
	return nil
}
