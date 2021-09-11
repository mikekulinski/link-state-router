package network

import (
	"fmt"
	"math"
)

type Switch struct {
	// The name of the switch for debugging
	Name string

	// The direct connections to its neighbors and their distances
	Neighbors map[*Switch]int

	// The topology of the network
	Graph Network

	// The network with distances
	routingTable map[*Switch][]*Switch
}

func NewSwitch(name string) *Switch {
	s := &Switch{}
	s.Name = name
	s.Neighbors = make(map[*Switch]int)
	s.routingTable = make(map[*Switch][]*Switch)
	return s
}

func (s *Switch) Route(dest *Switch) ([]*Switch, error) {
	// TODO: update this so that is reads from the routing table instead and instead
	// call dijkstra's when adding node to network
	return s.dijkstra(dest)
}

func (s *Switch) dijkstra(dest *Switch) ([]*Switch, error) {
	minDistances := initHeap(s.Graph, s)

	// Create a set for visited nodes and their final distance
	visited := make(map[*Switch]int)

	// Init routing table to just be the source
	s.routingTable[s] = []*Switch{s}
	visited[s] = 0

	currentNode := s
	for currentNode != dest {
		currentDist, ok := visited[currentNode]
		if !ok {
			return []*Switch{}, fmt.Errorf("error getting final distance for node: %v", currentNode.Name)
		}

		currentPath, ok := s.routingTable[currentNode]
		if !ok {
			return []*Switch{}, fmt.Errorf("error getting path from routing table for node: %+v", currentNode.Name)
		}

		for neighbor, dist := range currentNode.Neighbors {
			// Skip the nodes that we have already visited
			if _, ok := visited[neighbor]; ok {
				continue
			}

			prevMinDist, err := minDistances.Lookup(neighbor)
			if err != nil {
				return []*Switch{}, fmt.Errorf("error looking up node in minDistances: %w", err)
			}

			// Relax edge if we have a better value than we have previously calculated
			if prevMinDist == math.MaxInt64 || currentDist+dist < prevMinDist {
				err := minDistances.Update(neighbor, currentDist+dist)
				if err != nil {
					return []*Switch{}, fmt.Errorf("error updating relaxing edge: %w", err)
				}
				s.routingTable[neighbor] = append(currentPath, neighbor)
			}
		}

		link, err := minDistances.Extract()
		if err != nil {
			return []*Switch{}, fmt.Errorf("error extracting the next minimum node: %w", err)
		}

		// Mark the node as visited and update its final distance
		currentNode = link.Dest
		visited[currentNode] = link.Distance

	}

	shortestPath, ok := s.routingTable[dest]
	if !ok {
		return []*Switch{}, fmt.Errorf("error getting the shortest path for the destination")
	}

	return shortestPath, nil
}

func (s *Switch) calculateFinalDistance(node *Switch) {

}

func initHeap(network Network, source *Switch) *MinHeap {
	minHeap := NewMinHeap()
	for s := range network {
		if s == source {
			minHeap.Insert(s, 0)
		} else {
			minHeap.Insert(s, math.MaxInt64)
		}
	}
	return minHeap
}
