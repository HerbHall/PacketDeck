# PacketDeck — Claude Code Handoff

## What This Is

PacketDeck is a Docker Desktop extension for visual container network inspection. It shows network topology, connection states, and (eventually) packet capture — all natively inside Docker Desktop with zero configuration.

**Owner**: Herb Hall (github.com/HerbHall)
**License**: MIT
**Status**: Pre-development — scaffold only, no code yet
**Build Order**: Third project (after RunNotes and DockPulse)

---

## What Already Exists (D:\devspace\PacketDeck)

- ✅ `CLAUDE.md` — Project context for Claude Code
- ✅ `HANDOFF.md` — This file
- ✅ `README.md` — Project overview
- ✅ `CONTRIBUTING.md` — Contribution guidelines
- ✅ `CHANGELOG.md` — Keep-a-Changelog format
- ✅ `LICENSE` — MIT
- ✅ `VERSION` — 0.1.0
- ✅ `.gitignore` — Comprehensive
- ✅ `.editorconfig` — Workspace standard
- ✅ `metadata.json` — Docker extension metadata
- ✅ `Dockerfile` — Extension image stub (labels set, stages TODO)
- ✅ `Makefile` — Build/install/push targets
- ✅ `docker.svg` — Placeholder icon
- ✅ `docs/FEASIBILITY.md` — Full feasibility assessment

---

## What Still Needs To Be Done

### 1. GitHub Repository (FIRST PRIORITY)

```powershell
cd /d D:\devspace\PacketDeck
cmd /c "gh repo create HerbHall/PacketDeck --public --source=. --remote=origin --description "Docker Desktop extension — visual container network inspection and packet capture""
git add -A
git commit -m "chore: initial project scaffold"
git push -u origin main
```

### 2. Create GitHub Issues

Suggested issue backlog:

1. **Network topology data model** (mvp) — Parse `docker network inspect` and `docker inspect` to build container-network graph
2. **React UI: topology map visualization** (mvp) — Interactive network diagram showing containers, networks, port mappings, connections (D3.js or similar)
3. **Backend service: Docker network queries** (mvp) — Backend that polls Docker API for network/container state and pushes updates to frontend
4. **Container detail panel** (mvp) — Click container in topology → see IP addresses, ports, connected networks, environment hints
5. **Connection monitor: active TCP/UDP** (enhancement) — Show live connections per container via nsenter + ss
6. **Connection state tracking** (enhancement) — ESTABLISHED/LISTEN/TIME_WAIT counts, byte counters
7. **Packet capture: basic tcpdump integration** (enhancement) — Spawn netshoot sidecar, stream capture, display summary
8. **Pcap export** (enhancement) — Download captures as .pcap for Wireshark analysis
9. **Capture filtering** (enhancement) — Filter by port, protocol, IP, container pair
10. **Docker Hub publishing** (chore) — Multi-arch build, marketplace listing, screenshots

### 3. Source Directories (Create When Development Begins)

```text
ui/           — React frontend (create when starting issue #2)
backend/      — Go or Node backend service (create when starting issue #3)
```

---

## Key Architecture Decisions

These are settled from the research phase:

- **Tiered approach**: MVP is topology only (zero privileges needed), capture comes later
- **Topology data** from `docker network inspect` + `docker inspect` — no elevated permissions
- **Connection monitoring** requires `nsenter` into container network namespaces + `ss` command
- **Packet capture** via netshoot sidecar: `docker run --network container:target --cap-add NET_ADMIN nicolaka/netshoot tcpdump -w -`
- **DD VM advantage**: Extension backend runs inside the Linux VM alongside containers — direct namespace access, easier than host-based tools
- **Storage in Docker volume** for capture history and settings
- **React frontend** with D3.js or similar for topology visualization
- **Socket communication** for streaming data (especially live captures)
- **Multi-arch required**: linux/amd64 + linux/arm64

## Privilege Escalation Notes (Critical for Tier 2+)

- Tier 1 (topology): NO special privileges — uses standard Docker CLI/API
- Tier 2 (connections): Needs `SYS_PTRACE` / `SYS_ADMIN` on backend container
- Tier 3 (capture): Needs `NET_ADMIN`, `NET_RAW`, `SYS_PTRACE`, `SYS_ADMIN`
- DD VM runs as root, extension backends can configure capabilities in Dockerfile/metadata
- **Needs hands-on testing before committing to Tier 2+ features**

## Existing Tool Reference (Don't Reinvent)

- **Edgeshark (Siemens)**: MIT licensed, SharkFest 2023 talk documents architecture. Ghostwire (discovery) + Packetflix (streaming). Good reference implementation.
- **nicolaka/netshoot**: Swiss army knife Docker image. Use as sidecar for capture, not as dependency.
- **pcap-parser (JS)**: Browser-side pcap parsing. Limited vs Wireshark but sufficient for summary display.

## Subnetree Synergy

This project overlaps significantly with Herb's Subnetree network monitoring tool:
- Topology visualization work feeds directly into Subnetree UI design
- Go backend experience transfers between projects
- Network namespace inspection techniques are shared knowledge
- Consider shared visualization components long-term

## Name Research

"PacketDeck" was chosen after conflict checks confirmed:
- No GitHub repository named PacketDeck
- No Docker Hub image named packetdeck
- No npm package named packetdeck
- No trademark conflicts
- Conveys "packet inspection" + "Docker Desktop deck/dashboard"

## Herb's Preferences (Important)

- **Green/earthy colors** for branding — dislikes blue
- **PowerShell is primary shell** on Windows
- **Use `cmd /c` wrapper** for `gh` commands in PowerShell
- **Conventional commits**: `feat:`, `fix:`, `docs:`, `chore:`
- **Co-authored commits**: `Co-Authored-By: Claude <noreply@anthropic.com>`
- **Executes steps immediately as he reads them** — put prerequisites BEFORE action steps

---

## Suggested First Session Plan

1. Create GitHub repo + push scaffold
2. Create labels + issues
3. Begin MVP: Docker network/container data parsing (issue #1)
4. Build topology visualization UI (issue #2) — this is the flagship feature
5. Wire backend data to frontend graph (issue #3)
