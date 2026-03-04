/** Complete container network topology at a point in time. */
export interface Topology {
  networks: Network[];
  containers: Container[];
  edges: Edge[];
  timestamp: string;
}

/** Docker network. */
export interface Network {
  id: string;
  name: string;
  driver: string;
  subnet: string;
  gateway: string;
  scope: string;
}

/** Docker container and its network attachments. */
export interface Container {
  id: string;
  name: string;
  image: string;
  state: string;
  networks: Record<string, NetInfo>;
  ports: PortMapping[];
}

/** Container's connection details for a specific network. */
export interface NetInfo {
  networkId: string;
  ipAddress: string;
  macAddress: string;
  aliases: string[];
}

/** Container port binding to the host. */
export interface PortMapping {
  containerPort: number;
  hostPort: number;
  protocol: string;
}

/** Connection between a container and a network. */
export interface Edge {
  source: string;
  target: string;
  ipAddress: string;
}
