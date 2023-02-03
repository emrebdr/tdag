# tdag
## Directed Acyclic Graph

This is a simple implementation of a directed acyclic graph (DAG) in Go. It is not meant to be a general purpose graph library, but rather a simple implementation of a DAG that can be used to build more complex graph libraries.

### Usage

```go
import "github.com/tosuninc/tdag"

// Create a new DAG
dag := tdag.New()

// Add a vertex to the DAG
vertex1, _ := dag.AddVertex("vertex1")

// Add another vertex to the DAG
vertex2, _ := dag.AddVertex("vertex2")

// Add an edge between the two vertices
edge, err := dag.AddEdge(vertex1, vertex2)
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

### Acknowledgments

* [Go](https://golang.org/)
* [Go Modules](
    github.com/google/uuid
)

### Authors

* **Tosun Inc.** - *Initial work* - [Tosun Inc.](
    github.com/tosuninc/tdag
)
