package main

// Topology represents the complete container network topology at a point in time.
type Topology struct {
	Networks   []Network   `json:"networks"`
	Containers []Container `json:"containers"`
	Edges      []Edge      `json:"edges"`
	Timestamp  string      `json:"timestamp"`
}

// Network represents a Docker network.
type Network struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Driver  string `json:"driver"`
	Subnet  string `json:"subnet"`
	Gateway string `json:"gateway"`
	Scope   string `json:"scope"`
}

// Container represents a Docker container and its network attachments.
type Container struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Image    string            `json:"image"`
	State    string            `json:"state"`
	Networks map[string]NetInfo `json:"networks"`
	Ports    []PortMapping     `json:"ports"`
}

// NetInfo holds a container's connection details for a specific network.
type NetInfo struct {
	NetworkID  string   `json:"networkId"`
	IPAddress  string   `json:"ipAddress"`
	MacAddress string   `json:"macAddress"`
	Aliases    []string `json:"aliases"`
}

// PortMapping represents a container port binding to the host.
type PortMapping struct {
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	Protocol      string `json:"protocol"`
}

// Edge represents a connection between a container and a network.
type Edge struct {
	Source    string `json:"source"`
	Target   string `json:"target"`
	IPAddress string `json:"ipAddress"`
}
