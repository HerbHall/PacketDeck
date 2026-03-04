# PacketDeck -- Copilot Instructions

Docker Desktop extension for visual container network inspection and packet capture.

**Status**: Pre-development scaffold. Source code directories (`ui/`, `backend/`)
do not exist yet.

## Tech Stack

- **Frontend** (planned): React + Material UI + D3.js
- **Backend** (planned): TBD -- Node.js or Go
- **Build**: Docker Extensions CLI + Makefile
- **Platform**: Docker Desktop 4.8.0+

## Project Structure

```text
PacketDeck/
├── docs/                - Feasibility research and architecture
├── Dockerfile           - Extension image definition
├── Makefile             - Build targets
├── metadata.json        - Extension metadata
├── docker.svg           - Extension icon
├── CLAUDE.md            - Claude Code instructions
├── AGENTS.md            - Copilot coding agent instructions
└── HANDOFF.md           - Next steps and implementation plan
```

## Code Style

- Conventional commits: `feat:`, `fix:`, `refactor:`, `docs:`, `test:`, `chore:`
- Co-author tag: `Co-Authored-By: GitHub Copilot <noreply@github.com>`
- All lint checks must pass before committing

## Coding Guidelines

- Fix errors immediately -- never classify them as pre-existing
- Build must pass before any commit
- Never skip hooks (`--no-verify`) or force-push main
- Remove unused code completely; no backwards-compatibility hacks

## Available Resources

```bash
make build-extension     # Build the extension image
make install-extension   # Install into Docker Desktop
make update-extension    # Update after changes
make debug-extension     # Enable debug mode
make push-extension      # Push multi-arch image
make clean               # Remove extension
```

## Docker Desktop Extension Context

- Extension UI reinitializes on every tab switch -- state must come from backend
- Use `@docker/docker-mui-theme` for Docker Desktop theming (MUI v5)
- Multi-arch images required: `linux/amd64` + `linux/arm64`
- MVP uses zero elevated privileges -- topology from `docker network inspect`

## Do NOT

- Commit generated files without regenerating them first
- Add dependencies without updating the lock file
- Store secrets, tokens, or credentials in code or config files
- Mark work as complete when known errors remain
- Use `any` in TypeScript or suppress TypeScript errors with `as unknown`
- Add `//nolint` directives without fixing the root cause first
