package main

import (
	"fmt"
	"os"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/sijad/entgql/generator"
)

func generate(entSchemaPath, generatedGraphqlPath, generatedResolversPath, generatedHelpersPath, generatedModelsPath string) error {
	graph, err := entc.LoadGraph(entSchemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	if err := generateSchema(graph, generatedGraphqlPath); err != nil {
		return fmt.Errorf("generate schema: %w", err)
	}

	if err := generatedModels(graph, generatedModelsPath); err != nil {
		return fmt.Errorf("generate models: %w", err)
	}

	// TODO generate code

	return nil
}

func generateSchema(graph *gen.Graph, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return generator.SchemaDefinition(file, graph)
}

func generatedModels(graph *gen.Graph, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return generator.Models(file, graph)
}
