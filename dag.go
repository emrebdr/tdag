package dag

import (
	"errors"
	"fmt"
	"sync"

	"github.com/emrebdr/tdag/src/models"
	"github.com/emrebdr/tdag/src/utils"

	"github.com/google/uuid"
)

type Dag struct {
	Id       string
	mutex    sync.RWMutex
	vertices map[string]*models.Vertex
	edges    map[string]*models.Edge
}

func NewDag() *Dag {
	return &Dag{
		Id:       uuid.NewString(),
		vertices: map[string]*models.Vertex{},
		edges:    map[string]*models.Edge{},
	}
}

func (d *Dag) AddVertex(vertex interface{}) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.addVertex(vertex)
}

func (d *Dag) addVertex(vertex interface{}) (string, error) {
	id := uuid.NewString()

	newVertex := &models.Vertex{
		Id:    id,
		Node:  vertex,
		Edges: []string{},
	}

	d.vertices[id] = newVertex

	return id, nil
}

func (d *Dag) DeleteVertex(id string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.deleteVertex(id)
}

func (d *Dag) deleteVertex(id string) error {
	vertex := d.GetVertex(id)
	if vertex == nil {
		return errors.New("vertex not found")
	}

	for _, edge := range vertex.Edges {
		err := d.deleteEdge(vertex, edge)
		if err != nil {
			return errors.New("occurs error while deleting connected edges")
		}
	}

	delete(d.vertices, id)
	return nil
}

func (d *Dag) AddEdge(src, dst string) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.addEdge(src, dst)
}

func (d *Dag) addEdge(src, dst string) (string, error) {
	id := uuid.NewString()

	srcVertex := d.GetVertex(src)
	if srcVertex == nil {
		return "", errors.New("source vertex not found")
	}

	dstVertex := d.GetVertex(dst)
	if dstVertex == nil {
		return "", errors.New("destination vertex not found")
	}

	isEdgeExist := d.isEdgeExist(srcVertex, dstVertex)
	if isEdgeExist {
		return "", errors.New("the edge is already exist")
	}

	newEdge := &models.Edge{
		Id:   id,
		Tail: dstVertex,
	}

	d.edges[id] = newEdge
	srcVertex.Edges = append(srcVertex.Edges, id)

	cyclic := d.checkCyclic()
	if cyclic != nil {
		err := d.DeleteEdge(srcVertex, id)
		if err != nil {
			return "", errors.New("error occurs while deleting edge")
		}

		return "", errors.New("error cyclic dependency")
	}

	return id, nil
}

func (d *Dag) DeleteEdge(srcVertex *models.Vertex, id string) error {
	return d.deleteEdge(srcVertex, id)
}

func (d *Dag) deleteEdge(srcVertex *models.Vertex, id string) error {
	srcVertex.Edges = utils.RemoveFromArray(id, srcVertex.Edges)
	delete(d.edges, id)

	return nil
}

func (d *Dag) checkCyclic() error {
	for _, vertex := range d.vertices {
		var vertices []*models.Vertex
		_, cyclic := d.walk(vertex, vertices)
		if cyclic != nil {
			return errors.New("error cyclic dependency")
		}
	}

	return nil
}

func (d *Dag) walk(vertex *models.Vertex, vertices []*models.Vertex) ([]*models.Vertex, error) {
	for _, edgeId := range vertex.Edges {
		edge := d.GetEdge(edgeId)
		vertices = append(vertices, vertex)
		if validation := d.isValid(vertices); !validation {
			return nil, errors.New("error cyclic dependency")
		}

		return d.walk(edge.Tail, vertices)
	}

	return vertices, nil
}

func (d *Dag) isEdgeExist(src, dst *models.Vertex) bool {
	for _, id := range src.Edges {
		edge := d.GetEdge(id)
		if edge.Tail.Id == dst.Id {
			return true
		}
	}

	return false
}

func (d *Dag) isValid(vertices []*models.Vertex) bool {
	if len(vertices) > 1 {
		if vertices[0].Id == vertices[len(vertices)-1].Id {
			return false
		}

		for i := 0; i < len(vertices); i++ {
			for j := i + 1; j < len(vertices); j++ {
				if vertices[i].Id == vertices[j].Id {
					return false
				}
			}
		}
	}

	return true
}

func (d *Dag) GetVertex(id string) *models.Vertex {
	vertex := d.vertices[id]
	if vertex != nil {
		return vertex
	}

	return nil
}

func (d *Dag) GetAllVertices() map[string]*models.Vertex {
	return d.vertices
}

func (d *Dag) GetAcyclicGraphs() ([]*models.Vertex, error) {
	graphs := make(map[string][]*models.Vertex)
	for _, vertex := range d.vertices {
		var vertices []*models.Vertex
		graph, err := d.walk(vertex, vertices)
		if err != nil {
			return nil, err
		}

		graphs[vertex.Id] = graph
	}

	var unique_keys []string
	for key, vertices := range graphs {
		if len(vertices) < 1 {
			continue
		}

		count := 0
		for k, v := range graphs {
			if k == key {
				continue
			}

			for _, vertex := range v {
				if key == vertex.Id {
					count += 1
					break
				}
			}
		}

		if count == 0 {
			unique_keys = append(unique_keys, key)
		}
	}

	var vertices []*models.Vertex
	for _, key := range unique_keys {
		vertices = append(vertices, d.vertices[key])
	}

	return vertices, nil
}

func (d *Dag) GetEdge(id string) *models.Edge {
	edge := d.edges[id]
	if edge != nil {
		return edge
	}

	return nil
}

func (d *Dag) IsEdge(id string) (bool, error) {
	if vertex := d.vertices[id]; vertex != nil {
		return false, nil
	} else if edge := d.edges[id]; edge != nil {
		return true, nil
	}

	return false, errors.New("id not found")
}

func (d *Dag) IsVertex(id string) (bool, error) {
	if vertex := d.vertices[id]; vertex != nil {
		return true, nil
	} else if edge := d.edges[id]; edge != nil {
		return false, nil
	}

	return false, errors.New("id not found")
}

func (d *Dag) Print() {
	for _, vertex := range d.vertices {
		var vertices []*models.Vertex
		depends, err := d.walk(vertex, vertices)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range depends {
			fmt.Print(v.Id, " -->")
		}

		fmt.Println()
	}
}
