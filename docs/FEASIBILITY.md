# PacketDeck — Docker Desktop Extension Feasibility Assessment

## Summary

**Verdict**: Technically ambitious but feasible for an MVP. Most technically interesting of the three extensions. Leverages Herb's networking background (Navy Seabee, ISP co-founder, Subnetree).

**Concept**: Visual container network inspection natively inside Docker Desktop. Tiered approach from topology visualization (zero-privilege) through connection monitoring to packet capture.

**Market Gap**: No existing tool provides visual container network inspection natively inside Docker Desktop with zero configuration. Edgeshark (Siemens) is closest but requires separate deployment and Wireshark installation.

## The Problem

Docker Desktop shows container names/ports/status but nothing about network traffic. Debugging multi-container networking requires: finding veth pairs, installing tcpdump/netshoot, capturing with filters, transferring pcap files, opening in Wireshark. Multi-step, CLI-heavy, error-prone.

## Competitive Landscape

| Tool | In Docker Desktop | Visual UI | No CLI Required | Zero Setup | Capture |
|------|-------------------|-----------|-----------------|------------|---------|
| Docker Desktop (native) | ✅ | Port list only | ✅ | ✅ | ❌ |
| Edgeshark (Siemens) | ❌ (standalone) | ✅ (web) | Mostly | ❌ | ✅ via Wireshark |
| netshoot | ❌ | ❌ | ❌ | ❌ | ✅ |
| Wireshark container | ❌ | ✅ (VNC) | ✅ | ❌ | ✅ |
| **PacketDeck** | **✅** | **✅** | **✅** | **✅** | **Tier 3** |

## Tiered Build Approach

### Tier 1: Network Topology Map (MVP)
- Visual diagram of container networks (bridge, overlay, host)
- Which containers on which networks, port mappings, inter-container connections
- Data: `docker network inspect`, `docker inspect` — **no special privileges needed**
- Effort: 3-5 days

### Tier 2: Connection Monitor
- Active TCP/UDP connections per container, states, byte counts, latency
- Data: `nsenter` + `ss` or `/proc/net/tcp`
- Requires: `SYS_PTRACE`/`SYS_ADMIN` capabilities
- Effort: +1-2 weeks

### Tier 3: Traffic Capture
- Select container → capture packets → display summary in-extension
- Export pcap, live stream with filtering (port, protocol, IP)
- Data: tcpdump via netshoot sidecar or namespace-aware capture
- Requires: `NET_ADMIN`, `NET_RAW`, `SYS_PTRACE`, `SYS_ADMIN`
- Effort: +3-4 weeks

### Tier 4: Application-Layer Inspection
- HTTP viewer, DNS logging, TLS analysis
- Major complexity — 6-8 weeks additional

## Technical Architecture

```
[React Frontend] ←socket→ [Backend Service (Go/Node)]
                                    ↓
                           [Docker CLI / Docker API]
                                    ↓
                    [Container network inspection]
                    [Sidecar tcpdump containers]
                    [nsenter for namespace access]
```

**DD VM advantage**: On Windows/Mac, extension backend runs inside the Linux VM alongside containers — direct access to container network namespaces. Easier than host-based tools.

## Risks

1. **Privilege escalation** — Packet capture requires NET_ADMIN/SYS_PTRACE. DD may restrict for extension backends. Needs hands-on testing.
2. **Performance** — Streaming packet data through SDK to React could be slow. Need buffering/sampling.
3. **Pcap rendering** — Browser-based packet display requires JS dissector (limited vs Wireshark).
4. **Scope creep** — Network tools are endlessly extensible. Ship MVP first.
5. **Cross-platform** — DD VM networking differs between Windows (Hyper-V/WSL2), Mac, Linux.

## Inspiration

- **Edgeshark**: Container topology discovery with click-to-capture (gold standard)
- **Chrome DevTools Network tab**: UX model for app-layer inspection
- **Cilium Hubble**: Flow visualization for Kubernetes
- **Wireshark IO Graphs**: Time-series traffic volume per protocol
