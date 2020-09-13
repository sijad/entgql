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

	g := &generator.Generator{Graph: graph}

	if err := generateSchema(g, generatedGraphqlPath); err != nil {
		return fmt.Errorf("generate schema: %w", err)
	}

	if err := generatedModels(g, generatedModelsPath); err != nil {
		return fmt.Errorf("generate models: %w", err)
	}

	if err := generatedResolvers(g, generatedResolversPath); err != nil {
		return fmt.Errorf("generate resolvers: %w", err)
	}

	return nil
}

func generateSchema(g *generator.Generator, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return g.SchemaDefinition(file)
}

func generatedModels(g *generator.Generator, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return g.Models(file)
}

func generatedResolvers(g *generator.Generator, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return g.Resolvers(file)
}
