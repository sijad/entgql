<p align="center">🚧🚧 this is pre-alpha software 🚧🚧</p>

<h1 align="center">entgql</h1>

<p align="center">
  <strong>Simple</strong>, <strong>Efficient</strong>, <strong>CRUD GraphQL API</strong> Generator for <a href="https://entgo.io/">ent.</a> and <a href="https://gqlgen.com/">gqlgen</a>
</p>

### Quick Start

it's recommended to install entgql using [Go modules](https://github.com/golang/go/wiki/Modules#quick-start).

1. create [a new gqlgen project](https://gqlgen.com/getting-started/)
2. [install ent.](https://entgo.io/docs/getting-started/) and generate your models
3. install entgql by running `go get github.com/sijad/entgql/cmd/entgql`
4. generate CRUD code for your ent schema using `go run github.com/sijad/entgql/cmd/entgql generate ./ent/schema`
5. edit `./graph/resolver.go` and add ent client to the resolver struct:
```diff
-type Resolver struct{}
+import "your.com/project/path/ent"
+
+type Resolver struct {
+	EntClient *ent.Client
+}
```
6. edit `./server.go` to [create a new client](https://entgo.io/docs/crud/#create-a-new-client) and pass it to resolvers:
```diff
+import "your.com/project/path/ent"
+import _ "your.com/dialect/driver"

func main() {
// ...
+	client, err := ent.Open("dialect-name", "dialet-options")
+	if err != nil {
+		log.Fatal(err)
+	}
+	defer client.Close()

-	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
+	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{EntClient: client}}))
// ...
```
7. start your GraphQL server `go run server.go`

### Credits

inspired by [Nexus](https://www.nexusjs.org/) and [PostGraphile](https://www.graphile.org/postgraphile/)
