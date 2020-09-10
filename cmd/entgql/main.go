package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "gentc"}
	cmd.AddCommand(
		&cobra.Command{
			Use:     "generate",
			Short:   "generate go code and GraphQL schema for `ent` directory",
			Example: "entgql generate ./ent/schema",
			Args:    cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, paths []string) {
				entSchemaPath := paths[0]
				if err := generate(entSchemaPath, defaultGeneratedGraphqlSchemaPath, defaultGeneratedResolversPath); err != nil {
					log.Fatalln(err)
				}
			},
		},
	)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

const (
	defaultGeneratedGraphqlSchemaPath = "./graph/ent_schema.graphqls"
	defaultGeneratedResolversPath     = "./graph/ent_schema.resolvers.go"
)
