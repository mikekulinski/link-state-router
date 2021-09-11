package network

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	createNetwork()
}

func createNetwork() (Network, []*Switch) {
	network := make(Network)

	// Init switches
	switches := make([]*Switch, 6)
	for i := 0; i < 6; i++ {
		// Convert index to letter in alphabet and set as name
		s := NewSwitch(string(rune('A' + i)))
		s.Graph = network

		switches[i] = s
		network.InsertSwitch(switches[i])
	}

	network.Print()

	// Add the links between the switches
	insertLinks(network, switches, 0, []int{1, 2, 3}, []int{1, 2, 6})
	insertLinks(network, switches, 1, []int{0, 3}, []int{1, 4})
	insertLinks(network, switches, 2, []int{0, 4}, []int{2, 7})
	insertLinks(network, switches, 3, []int{0, 1, 4, 5}, []int{6, 4, 3, 8})
	insertLinks(network, switches, 4, []int{2, 3, 5}, []int{7, 3, 2})
	insertLinks(network, switches, 5, []int{3, 4}, []int{8, 2})

	return network, switches
}

func insertLinks(n Network, switches []*Switch, fromIndex int, toIndices []int, distances []int) {
	for i, to := range toIndices {
		n.InsertLink(switches[fromIndex], switches[to], distances[i])
	}
}
