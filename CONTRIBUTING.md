# Contributing to PacketDeck

Thanks for your interest in contributing to PacketDeck!

## Getting Started

1. Fork the repository
2. Clone your fork locally
3. Create a feature branch from `main`
4. Make your changes
5. Submit a pull request

## Development Setup

Prerequisites:

- Docker Desktop 4.8.0+
- Node.js 18+ (for frontend development)
- Go 1.21+ (if working on backend)

Build and install locally:

```bash
make build-extension
make install-extension
```

## Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` — New feature
- `fix:` — Bug fix
- `docs:` — Documentation only
- `chore:` — Maintenance, build, CI
- `refactor:` — Code change that neither fixes a bug nor adds a feature

## Pull Requests

- Reference the issue number in your PR description
- Keep PRs focused — one feature or fix per PR
- Update documentation if behavior changes

## Code Style

- Frontend: Follow existing React/Material UI patterns
- Backend: `gofmt` for Go, Prettier for JavaScript/TypeScript
- Use `.editorconfig` settings (your editor should pick these up)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
