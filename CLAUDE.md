# PacketDeck

Docker Desktop extension — visual container network inspection and packet capture.

## Tech Stack

- **Frontend**: React + Material UI + D3.js (topology visualization)
- **Backend**: TBD — Node.js or Go service running in Desktop VM
- **Storage**: SQLite or JSON file on mounted volume
- **Build**: Docker Extensions CLI + Makefile
- **Platform**: Docker Desktop 4.8.0+ (Windows, Mac, Linux)

## Key Design Decisions

- Tiered approach: MVP is topology visualization only (zero privileges needed)
- Topology data from `docker network inspect` + `docker inspect`
- Connection monitoring via nsenter + ss (requires SYS_PTRACE)
- Packet capture via netshoot sidecar (requires NET_ADMIN, NET_RAW)
- Extension UI reinitializes on every tab switch — all state must come from backend
- DD VM runs containers alongside extension backend — direct namespace access

## Project Conventions

- Commit messages: conventional commits (`feat:`, `fix:`, `docs:`, `chore:`)
- Co-authored commits with Claude: `Co-Authored-By: Claude <noreply@anthropic.com>`
- Issues track all work; PRs reference issue numbers
- PowerShell is the primary scripting shell on Windows

## File Layout

```text
PacketDeck/
├── docs/                - Feasibility research
├── ui/                  - React frontend (D3.js topology)
├── backend/             - Backend service
├── metadata.json        - Extension metadata
├── Dockerfile           - Extension image
└── Makefile             - Build targets
```
