package generator

import (
	"bytes"
	"go/format"
	"io"

	"github.com/facebook/ent/entc/gen"
)

type ResolversData struct {
	Imports map[string]string
}

func (g *Generator) Resolvers(wr io.Writer) error {
	imports := map[string]string{}

	buf := &bytes.Buffer{}
	pkg := g.Graph.Config.Package

	for _, t := range g.Graph.Nodes {
		name := t.Package()
		imports[pkg+"/"+name] = ""
		// TODO import other inputs types
	}

	data := &ModelsData{
		Imports: imports,
	}
	if err := templates.ExecuteTemplate(buf, "template/base_resolvers_go.tmpl", data); err != nil {
		return err
	}

	for _, t := range g.Graph.Nodes {
		if err := nodeResolvers(buf, t); err != nil {
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

func nodeResolvers(wr io.Writer, t *gen.Type) error {
	return templates.ExecuteTemplate(wr, "template/node_resolvers_go.tmpl", t)
}
