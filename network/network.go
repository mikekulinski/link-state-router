package network

import "fmt"

type Network map[*Switch][]Link

func (n Network) InsertSwitch(s *Switch) error {
	if _, ok := n[s]; ok {
		return fmt.Errorf("tried to insert switch that already exists: %v", *s)
	}

	// Insert the switch into the network with no links
	n[s] = make([]Link, 0)

	return nil
}

func (n Network) InsertLink(from *Switch, to *Switch, distance int) error {
	if _, ok := n[from]; !ok {
		return fmt.Errorf("tried to insert link for switch that doesn't exist: %s", from.Name)
	}

	// Insert the links into the network
	n[from] = append(n[from], Link{Dest: to, Distance: distance})

	// Update the neighbors for each of the switches that are updated
	from.Neighbors[to] = distance

	return nil
}

func (n Network) Print() {
	for s, links := range n {
		fmt.Printf("%s: [\n", s.Name)
		for _, link := range links {
			fmt.Printf("\t{%s: %v}\n", link.Dest.Name, link.Distance)
		}
		fmt.Printf("]\n")
	}
}
