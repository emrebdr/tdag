package dag_test

import (
	"testing"

	tdag "github.com/emrebdr/tdag"
)

func TestCreateDag(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}
}

func TestAddVertex(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex, _ := dag.AddVertex("test vertex")
	if vertex == "" {
		t.Errorf("Couldn't add vertex")
	}
}

func TestAddEdge(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge, err := dag.AddEdge(vertex1, vertex2)
	if edge == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}
}

func TestAddEdgeCyclicError(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex3, _ := dag.AddVertex("test vertex 3")
	if vertex3 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge1, err := dag.AddEdge(vertex1, vertex2)
	if edge1 == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	edge2, err := dag.AddEdge(vertex2, vertex3)
	if edge2 == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	edge3, err := dag.AddEdge(vertex3, vertex1)
	if edge3 != "" {
		t.Errorf("Couldn't add edge")
	}

	if err == nil {
		t.Errorf("Couldn't add edge")
	}
}

func TestDeleteVertex(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	err := dag.DeleteVertex(vertex1)
	if err != nil {
		t.Errorf("Couldn't delete vertex")
	}
}

func TestGetVertex(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	getVertex := dag.GetVertex(vertex1)
	if getVertex == nil {
		t.Errorf("Couldn't get vertex")
	}
}

func TestGetEdge(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge, err := dag.AddEdge(vertex1, vertex2)
	if edge == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	getEdge := dag.GetEdge(edge)
	if getEdge == nil {
		t.Errorf("Couldn't get edge")
	}
}

func TestDeleteEdge(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge, err := dag.AddEdge(vertex1, vertex2)
	if edge == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	getVertex := dag.GetVertex(vertex1)

	err = dag.DeleteEdge(getVertex, edge)
	if err != nil {
		t.Errorf("Couldn't delete edge")
	}
}

func TestGetAllVertices(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertices := dag.GetAllVertices()
	if len(vertices) != 2 {
		t.Errorf("Couldn't get all vertices")
	}
}

func TestGetAcyclicGraphs(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex3, _ := dag.AddVertex("test vertex 3")
	if vertex3 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge1, err := dag.AddEdge(vertex1, vertex2)
	if edge1 == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	edge2, err := dag.AddEdge(vertex2, vertex3)
	if edge2 == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	acyclicGraphs, err := dag.GetAcyclicGraphs()
	if len(acyclicGraphs) != 1 {
		t.Errorf("Couldn't get acyclic graphs")
	}

	if err != nil {
		t.Errorf("Couldn't get acyclic graphs")
	}
}

func TestIsEdge(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	vertex2, _ := dag.AddVertex("test vertex 2")
	if vertex2 == "" {
		t.Errorf("Couldn't add vertex")
	}

	edge, err := dag.AddEdge(vertex1, vertex2)
	if edge == "" {
		t.Errorf("Couldn't add edge")
	}

	if err != nil {
		t.Errorf("Couldn't add edge")
	}

	isEdge, err := dag.IsEdge(edge)
	if isEdge == false {
		t.Errorf("Couldn't found edge")
	}

	if err != nil {
		t.Errorf("Couldn't found edge")
	}
}

func TestIsVertex(t *testing.T) {
	dag := tdag.NewDag()

	if dag == nil {
		t.Errorf("Couldn't create graph")
	}

	vertex1, _ := dag.AddVertex("test vertex 1")
	if vertex1 == "" {
		t.Errorf("Couldn't add vertex")
	}

	isVertex, err := dag.IsVertex(vertex1)
	if isVertex == false {
		t.Errorf("Couldn't found vertex")
	}

	if err != nil {
		t.Errorf("Couldn't found vertex")
	}
}
