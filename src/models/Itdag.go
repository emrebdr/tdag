package models

type ITDag interface {
	AddVertex(vertex interface{}) (string, error)
	AddEdge(src, dst string) (string, error)
	DeleteVertex(id string) error
	DeleteEdge(srcVertex *Vertex, id string) error
	GetVertex(id string) *Vertex
	GetAllVertices() map[string]*Vertex
	GetAcyclicGraphs() ([]*Vertex, error)
	GetEdge(id string) *Edge
	IsEdge(id string) (bool, error)
	IsVertex(id string) (bool, error)
	Print()
}
