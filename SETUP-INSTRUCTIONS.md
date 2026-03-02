# PacketDeck — Agent Setup Instructions

> **Purpose**: Complete project setup before any development begins.
> **Prerequisites**: `gh` CLI installed and authenticated as HerbHall. Run all commands in CMD (not PowerShell) to avoid bracket escaping issues.
> **Shell**: CMD (`cmd /c` wrapper if calling from PowerShell)
> **Working directory**: `D:\devspace\PacketDeck`

---

## Pre-flight Checks

Before running any commands, verify:

1. `gh auth status` returns authenticated as HerbHall
2. `git --version` returns a version
3. Check whether `D:\devspace\PacketDeck\.git` exists — if NOT, Step 1 includes `git init`

---

## Step 1 — Initialize Git and Create GitHub Repository

PacketDeck does NOT have a `.git` directory yet. Initialize before creating the remote.

```cmd
cd /d D:\devspace\PacketDeck
git init
gh repo create HerbHall/PacketDeck --public --source=. --remote=origin --description "Docker Desktop extension — visual container network inspection and packet capture"
git add -A
git commit -m "chore: initial project scaffold"
git push -u origin main
```

**Verify**: `gh repo view HerbHall/PacketDeck` shows the repo. `git log --oneline -1` shows the commit.

---

## Step 2 — Create Issue Labels

These labels must exist before the issue batch file runs.

```cmd
cd /d D:\devspace\PacketDeck
gh label create mvp --color 0E8A16 --description "Minimum viable product" --repo HerbHall/PacketDeck
gh label create feat --color 1D76DB --description "New feature" --repo HerbHall/PacketDeck
gh label create enhancement --color A2EEEF --description "Enhancement" --repo HerbHall/PacketDeck
gh label create chore --color FBCA04 --description "Maintenance task" --repo HerbHall/PacketDeck
gh label create docs --color 0075CA --description "Documentation" --repo HerbHall/PacketDeck
```

**Verify**: `gh label list --repo HerbHall/PacketDeck` shows all five labels.

---

## Step 3 — Create GitHub Issues

```cmd
cd /d D:\devspace\PacketDeck
create-issues.bat
```

This creates 10 issues covering the full backlog (4 MVP, 4 enhancement, 1 chore, 1 docs).

**Verify**: `gh issue list --repo HerbHall/PacketDeck` shows 10 open issues.

---

## Step 4 — Clean Up Batch Artifacts

The batch file creates a temporary `.body` file. Delete it if present and make sure it stays in `.gitignore` (it already is).

```cmd
cd /d D:\devspace\PacketDeck
if exist .body del .body
if exist .issuebody del .issuebody
```

---

## Step 5 — Final Commit

If any files were created or modified during setup:

```cmd
cd /d D:\devspace\PacketDeck
git add -A
git status
git commit -m "chore: complete project setup" --allow-empty
git push
```

---

## Completion Checklist

- [ ] GitHub repo `HerbHall/PacketDeck` exists and is public
- [ ] All scaffold files pushed to `main` branch
- [ ] 5 labels created (mvp, feat, enhancement, chore, docs)
- [ ] 10 issues created with correct labels
- [ ] No temp files (.body, .issuebody) left in working directory
- [ ] `git status` is clean

---

## What NOT To Do

- Do NOT create `ui/` or `backend/` directories — those get scaffolded when development begins
- Do NOT modify CLAUDE.md, HANDOFF.md, or docs/ content — research phase is complete
- Do NOT run `docker extension init` — the scaffold is already in place and customized
- Do NOT use PowerShell directly for `gh` commands — use CMD or `cmd /c` wrapper
