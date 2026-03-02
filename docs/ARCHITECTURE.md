# PacketDeck Architecture

## Overview

PacketDeck is a Docker Desktop extension for visualizing and inspecting container network traffic. It uses a tiered capability approach — each tier adds deeper inspection with progressively higher privilege requirements.

## Tiered Capability Model

### Tier 1: Network Topology Map (MVP)

**What it shows**: Visual graph of Docker networks with containers as nodes, networks as groups, port mappings as labels, and inter-container links.

**Data sources** (no special privileges):
- `docker network ls` — list all networks
- `docker network inspect <name>` — containers attached, subnet, gateway
- `docker inspect <container>` — ports, network settings, aliases

**Visualization**: D3.js force-directed graph or dagre layout. Containers as nodes, networks as colored groups, edges show connectivity.

**Privilege level**: None beyond standard extension access ✅

### Tier 2: Connection Monitor

**What it shows**: Active TCP/UDP connections per container — remote endpoints, states (ESTABLISHED, TIME_WAIT, etc.), byte counts.

**Data sources** (elevated privileges required):
- `nsenter --target <pid> --net ss -tunap` — socket stats from container's network namespace
- Or read `/proc/<pid>/net/tcp` and `/proc/<pid>/net/tcp6` directly

**Privilege level**: `SYS_PTRACE` + `SYS_ADMIN` capabilities on backend container

### Tier 3: Traffic Capture

**What it shows**: Packet-level capture for a selected container. Summary table (src, dst, protocol, size, timestamp). Filter by port/protocol/IP. Export to pcap.

**Data sources** (highest privileges):
- `nsenter --target <pid> --net tcpdump -i any -w -` — live capture from container namespace
- Or sidecar pattern: attach netshoot container to target's network namespace

**Privilege level**: `NET_ADMIN` + `NET_RAW` + `SYS_PTRACE` + `SYS_ADMIN`

### Tier 4: Application Layer (Future)

HTTP request/response viewer, DNS query log, TLS handshake analysis. Requires mitmproxy-style interception or eBPF. Major complexity — defer to post-1.0.

## Component Architecture

```
┌─────────────────────────────────────────────────┐
│  Docker Desktop                                 │
│  ┌───────────────────────────────────────────┐  │
│  │  PacketDeck Tab (React + MUI + D3.js)    │  │
│  │  ┌──────────────┐  ┌──────────────────┐  │  │
│  │  │  Topology    │  │  Connection      │  │  │
│  │  │  Graph       │  │  Table / Capture │  │  │
│  │  │  (D3.js)     │  │  Viewer          │  │  │
│  │  └──────┬───────┘  └────────┬─────────┘  │  │
│  │         │ socket             │ socket     │  │
│  │  ┌──────▼───────────────────▼──────────┐  │  │
│  │  │  Backend Service (VM)               │  │  │
│  │  │  ┌────────────┐ ┌───────────────┐  │  │  │
│  │  │  │ Topology   │ │ Capture       │  │  │  │
│  │  │  │ Scanner    │ │ Engine        │  │  │  │
│  │  │  │ (docker    │ │ (nsenter +    │  │  │  │
│  │  │  │  inspect)  │ │  tcpdump)     │  │  │  │
│  │  │  └────────────┘ └───────────────┘  │  │  │
│  │  │  ┌──────────────────────────────┐  │  │  │
│  │  │  │ SQLite (topology snapshots, │  │  │  │
│  │  │  │ capture history, prefs)     │  │  │  │
│  │  │  └──────────────────────────────┘  │  │  │
│  │  └────────────────────────────────────┘  │  │
│  └───────────────────────────────────────────┘  │
│  Docker Engine API + Container Namespaces       │
└─────────────────────────────────────────────────┘
```

## Key Technical Considerations

### DD VM Advantage

On Windows/Mac, Docker Desktop runs containers inside a Linux VM. The extension backend also runs in this VM, which means:

- Direct access to container network namespaces (no SSH needed)
- `nsenter` works natively against container PIDs
- Simpler than host-based tools that need to bridge into the VM

### Privilege Escalation

The backend container needs elevated capabilities for Tier 2+. Extensions can request capabilities in their Dockerfile or via compose configuration in metadata.json. This needs hands-on testing:

```dockerfile
# Needs validation — can DD extensions get these?
# Option 1: Dockerfile capability hints
# Option 2: compose.yaml in metadata.json with cap_add
```

**If DD restricts capabilities**: Tier 1 (topology) still works fully. Tier 2/3 may need a "helper container" pattern where the extension spawns a privileged sidecar for capture operations.

### Streaming Data

Packet capture generates high-volume data. Architecture must handle:

- Backend buffers captures, sends summaries to frontend via socket
- Frontend renders rolling table (newest first, capped at N rows)
- Full pcap stored to temp file for export
- Ring buffer pattern to prevent memory exhaustion

## Data Model

### Network Topology Snapshot

| Field | Type | Description |
|-------|------|-------------|
| id | INTEGER | Auto-increment PK |
| network_name | TEXT | Docker network name |
| network_id | TEXT | Docker network ID |
| driver | TEXT | bridge, overlay, host, macvlan |
| subnet | TEXT | CIDR notation |
| gateway | TEXT | Gateway IP |
| containers | TEXT | JSON array of attached container info |
| captured_at | DATETIME | Snapshot timestamp |

### Capture Session

| Field | Type | Description |
|-------|------|-------------|
| id | INTEGER | Auto-increment PK |
| container_name | TEXT | Target container |
| container_id | TEXT | Container ID |
| started_at | DATETIME | Capture start |
| ended_at | DATETIME | Capture end |
| packet_count | INTEGER | Total packets captured |
| pcap_path | TEXT | Path to pcap file (if saved) |
| filters | TEXT | Applied capture filters |

## Competitive Landscape

| Tool | In Docker Desktop | Visual Topology | Packet Capture | Setup Required |
|------|-------------------|----------------|----------------|----------------|
| Edgeshark (Siemens) | ❌ | Yes (web UI) | Yes → Wireshark | Docker Compose + Wireshark install |
| netshoot | ❌ | ❌ | CLI tools | Manual container attach |
| linuxserver/wireshark | ❌ | ❌ | Yes (web VNC) | Heavy container + host networking |
| Portainer | ❌ | Network list (no graph) | ❌ | Separate deployment |
| **PacketDeck** | **✅ Native tab** | **Yes (D3.js)** | **Yes (built-in)** | **Zero config** |

## Inspiration

- **Edgeshark**: Topology discovery + click-to-capture (gold standard for functionality)
- **Chrome DevTools Network tab**: UX model for application-layer inspection
- **Cilium Hubble UI**: Flow visualization for Kubernetes — clean graph aesthetics
- **Wireshark**: Packet dissection depth (we won't match this, but pcap export bridges the gap)
