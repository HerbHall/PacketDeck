<!--
  Scope: AGENTS.md guides the Copilot coding agent and Copilot Chat.
  For code completion and code review patterns, see .github/copilot-instructions.md
  For Claude Code, see CLAUDE.md
-->

# PacketDeck

Docker Desktop extension for visual container network inspection and packet capture.
Brings network visibility natively into Docker Desktop -- topology maps, connection
monitoring, and packet capture without leaving the IDE.

**Status**: Pre-development scaffold. No source code directories exist yet.

## Tech Stack

- **Frontend** (planned): React + Material UI + D3.js (topology visualization)
- **Backend** (planned): TBD -- Node.js or Go service running in Desktop VM
- **Storage** (planned): SQLite or JSON file on mounted volume
- **Build**: Docker Extensions CLI + Makefile
- **Platform**: Docker Desktop 4.8.0+ (Windows, Mac, Linux)

## Build and Test Commands

```bash
# Build the extension image
make build-extension

# Install into Docker Desktop
make install-extension

# Update after changes
make update-extension

# Debug mode
make debug-extension

# Push multi-arch image
make push-extension

# Remove extension
make clean
```

## Project Structure

```text
PacketDeck/
├── .github/             - CI workflows and Copilot config
├── docs/                - Feasibility research and architecture
│   ├── ARCHITECTURE.md  - Architecture design
│   └── FEASIBILITY.md   - Feasibility analysis
├── Dockerfile           - Extension image definition
├── Makefile             - Build targets
├── metadata.json        - Extension metadata (name, icon, description)
├── docker.svg           - Extension icon
├── CLAUDE.md            - Claude Code instructions
├── CONTRIBUTING.md      - Contribution guidelines
├── HANDOFF.md           - Next steps and implementation plan
├── VERSION              - Current version
└── CHANGELOG.md         - Version history
```

When source code directories are created, they will follow this layout:

```text
├── ui/                  - React frontend (D3.js topology)
└── backend/             - Backend service
```

## Workflow Rules

### Always Do

- Create a feature branch for every change (`feature/issue-NNN-description`)
- Use conventional commits: `feat:`, `fix:`, `refactor:`, `docs:`, `test:`, `chore:`
- Run build before opening a PR (`make build-extension`)
- Test extension installs cleanly (`make install-extension`)
- Fix every error you find, regardless of who introduced it

### Ask First

- Adding new dependencies
- Architectural changes (e.g., choosing Go vs Node.js backend)
- Changes to the tiered capability model (Tier 1-4 features)
- Changes to CI/CD workflows
- Docker Desktop API version bumps

### Never Do

- Commit directly to `main` -- always use feature branches
- Skip tests or lint checks -- even for "small changes"
- Use `--no-verify` or `--force` flags
- Commit secrets, credentials, or API keys
- Add TODO comments without a linked issue number
- Mark work as complete when build failures remain

## Core Principles

These are unconditional -- no optimization or time pressure overrides them:

1. **Quality**: Once found, always fix, never leave. There is no "pre-existing" error.
2. **Verification**: Build must pass before any commit.
3. **Safety**: Never force-push `main`. Never skip hooks. Never commit secrets.
4. **Honesty**: Never mark work as complete when it is not.

## Docker Desktop Extension Notes

- Extension UI reinitializes on every tab switch -- all state must come from backend
- DD VM runs containers alongside extension backend -- direct namespace access
- MVP (Tier 1) needs zero elevated privileges -- topology from `docker network inspect`
- Higher tiers require capabilities: `SYS_PTRACE` (Tier 2), `NET_ADMIN`/`NET_RAW` (Tier 3)
- Multi-arch images required: `linux/amd64` + `linux/arm64`
- Use `@docker/docker-mui-theme` for consistent Docker Desktop theming (pins MUI v5)

## Commit Format

```text
feat: add network topology visualization

Implements D3.js force-directed graph for container network layout.

Closes #42
Co-Authored-By: GitHub Copilot <copilot@github.com>
```

Types: `feat` (new feature), `fix` (bug fix), `refactor` (no behavior change),
`docs` (documentation only), `test` (tests only), `chore` (build/tooling).
