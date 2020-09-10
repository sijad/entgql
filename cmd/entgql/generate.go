package main

import (
	"fmt"
	"os"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/sijad/entgql/generator"
)

func generate(entSchemaPath, generatedGraphqlPath, generatedResolversPath string) error {
	graph, err := entc.LoadGraph(entSchemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("generate code: %w", err)
	}

	file, err := os.Create(generatedGraphqlPath)
	if err != nil {
		return fmt.Errorf("generate code: %w", err)
	}
	defer file.Close()

	if err := generator.SchemaDefinition(file, graph); err != nil {
		return fmt.Errorf("generate code: %w", err)
	}

	// TODO generate code

	return nil
}
