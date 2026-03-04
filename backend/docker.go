package main

import "time"

// getTopology returns the current container network topology.
// This is a stub that returns empty data. Real implementation comes in issue #4.
func getTopology() Topology {
	return Topology{
		Networks:   []Network{},
		Containers: []Container{},
		Edges:      []Edge{},
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}
