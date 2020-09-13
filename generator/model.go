package generator

import (
	"bytes"
	"go/format"
	"io"

	"github.com/facebook/ent/entc/gen"
)

type ModelsData struct {
	Imports map[string]string
}

func (g *Generator) Models(wr io.Writer) error {
	imports := map[string]string{
		g.Graph.Package: "ent",
	}

	data := &ModelsData{
		Imports: imports,
	}

	buf := &bytes.Buffer{}

	if err := templates.ExecuteTemplate(buf, "template/base_models_go.tmpl", data); err != nil {
		return err
	}

	for _, t := range g.Graph.Nodes {
		if err := nodeModels(buf, t); err != nil {
			return err
		}
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// write to buffer for debuging
		if _, err := wr.Write(buf.Bytes()); err != nil {
			return err
		}
		return err
	}

	if _, err := wr.Write(formatted); err != nil {
		return err
	}

	return nil
}

func nodeModels(wr io.Writer, t *gen.Type) error {
	return templates.ExecuteTemplate(wr, "template/node_models_go.tmpl", t)
}
