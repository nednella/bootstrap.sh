# bootstrap.sh

A Go CLI that bootstraps a fresh macOS install: Homebrew + Brewfile, dotfiles, macOS preferences. Installable in one command on a truly fresh machine (only `curl` required).

---

## Architecture

Monorepo. This repo contains both the CLI tool's Go source and the dotfiles it manages.

The binary is distributed as a prebuilt artefact via GitHub Releases. `install.sh` clones this repo to a known location on disk (default `~/.bootstrap.sh/`) and drops the binary on `$PATH`. From then on, the binary uses the clone as the live source for symlinks. Edits to symlinked files write through to the clone, so changes can be committed and pushed back upstream — the dev workflow is normal git.

The Go source ends up on disk alongside the dotfiles inside the clone. It is unused at runtime (the binary on `$PATH` is already built) — accepted as harmless dead weight in exchange for a single-curl install with no Go required on the user's machine.

### Install flow (fresh Mac)

User runs:

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/install.sh | bash
```

`install.sh` (idempotent throughout):

1. Installs Homebrew if missing (which pulls in `xcode-select`, giving us `git`).
2. Clones this repo into the install path if absent (default `~/.bootstrap.sh/`).
3. Downloads the latest prebuilt `bootstrap` binary from GitHub Releases.
4. Drops it on `$PATH` at `$(brew --prefix)/bin/bootstrap`.

Re-running install.sh is safe: brew is skipped if installed, the clone is skipped if it exists, the binary is refreshed.

### Runtime flow

After `install.sh`, the clone exists and the binary is on `$PATH`. Subcommands operate against the clone:

- `bootstrap install` — runs `brew bundle` against `<install_path>/Brewfile`. Defensively installs Homebrew first if somehow missing.
- `bootstrap dotfiles` — symlinks `<install_path>/dotfiles/` into `$HOME` and `$XDG_CONFIG_HOME` per the convention below. Existing files get backed up to `~/.dotfiles-backup/<timestamp>/` first.
- `bootstrap macos` — runs the macOS settings job.

### Update flow

`bootstrap update` is a single command that updates both binary and content:

1. Checks GH releases for a newer binary; if newer, downloads and atomically replaces self.
2. Stash-aware `git pull` on the clone:
   ```
   git status --short            # any uncommitted changes?
   git stash -u                  # if dirty, stash including untracked
   git pull
   git stash pop                 # if we stashed, restore
   ```

Local edits to symlinked dotfiles survive the pull. To push local edits upstream: `cd ~/.bootstrap.sh && git commit && git push`. Standard git.

---

## Configuration

A YAML config (`internal/config/default_config.yaml`) is embedded into the binary via `//go:embed`. To change defaults, edit the YAML in source and release a new binary.

CLAUDE.md does not duplicate the defaults — see the YAML file for what's configurable.

There is no runtime config-file overlay (yet). Single-user tool; the embedded YAML is the source of truth.

---

## Dotfiles convention

The `dotfiles/` tree is flat: one directory per program. Within each program directory, the symlink destination is determined by a single rule:

- Files starting with `.` are symlinked relative to `$HOME`.
- Files not starting with `.` are symlinked relative to `${XDG_CONFIG_HOME:-$HOME/.config}/<program>/`.

| Source | Destination |
|---|---|
| `dotfiles/zsh/.zshrc` | `~/.zshrc` |
| `dotfiles/git/.gitconfig` | `~/.gitconfig` |
| `dotfiles/ghostty/config` | `~/.config/ghostty/config` |
| `dotfiles/starship/starship.toml` | `~/.config/starship/starship.toml` |

No metadata files needed — the dot prefix carries the intent.

---

## Command surface

```
bootstrap                 # run all jobs in order (install, dotfiles, macos)
bootstrap install         # install packages from Brewfile
bootstrap dotfiles        # symlink dotfiles into $HOME / XDG
bootstrap macos           # apply macOS preferences
bootstrap update          # update binary + pull latest content

--dry-run                 # global flag; print what would happen, change nothing
```

CLI is built with **Cobra**. `main.go` is thin (wires the root command, defers everything else to `cmd/`).

---

## Repository layout

```
bootstrap.sh/
├── CLAUDE.md           # this file
├── README.md           # user-facing docs
├── install.sh          # the curl one-liner: installs Homebrew + clones repo + downloads binary
├── Brewfile            # consumed by `bootstrap install` from the clone at runtime
├── dotfiles/           # consumed by `bootstrap dotfiles` from the clone at runtime
│   ├── ghostty/
│   ├── git/
│   ├── zsh/
│   └── ...
├── go.mod
├── main.go             # Cobra root + subcommand wiring; thin
├── cmd/                # one file per subcommand: parse → call into internal/jobs
└── internal/
    ├── config/         # embedded YAML + loader
    ├── ui/             # styled output, status messages, prompts
    ├── jobs/           # one file per job: install, dotfiles, macos, update — actual work
    └── utils/          # shared primitives added just-in-time (shell-out runner, fs, symlinks)
```

`cmd/` is the *interface* layer (Cobra). `internal/jobs/` is the *logic* layer. Same shape as API endpoint ↔ service.

---

## Conventions

### Terminology

The unit of work is always a **job**. Never "phase", "step", or "task".

### Language

British English (en-GB) in commits, comments, prose, and identifiers I control (`colour`, `behaviour`, `initialise`). Standard Go stdlib names stay American (`color.Color`, etc.) — those are not mine to control.

### Go style

Never write `if err := foo(); err != nil { ... }`. Split assignment onto its own line:

```go
err := foo()
if err != nil {
    return err
}
```

### Declaration order

A function called by a neighbouring function — a helper in a local call chain — sits in execution (reading) order, caller above callee (see `internal/utils/preflight.go`, `internal/jobs/dotfiles.go`). A function only ever called from elsewhere — an independent primitive or entry point — is ordered alphabetically (see `internal/utils/run.go`).

### Casing

Standard Go: exported identifiers `PascalCase`, unexported `camelCase`, initialisms all-caps (`URL`, `ID`, `HTTP`). Never export a helper purely so another package can reach it — that breaks the casing consistency of its sibling helpers. Route the cross-package call through the package's public entry point instead.

### Style var naming

Prefix style vars with the feature that consumes them: `headerArrowStyle`, not `arrowStyle`.

### Commit discipline

"and" / "&" in a commit subject means the commit should be split into multiple commits. One logical change per commit. Subject-only — no commit body unless genuinely load-bearing.

### Command framing

Describe commands (Cobra `Short`/`Long`, README usage, command-surface lines) by the **primary action**. Prerequisites and defensive checks are implementation detail, not user-facing copy.

---

## Milestones

- [ ] **M1** — Cobra CLI skeleton: root command, subcommand dispatch, help text (no job logic yet)
- [ ] **M2** — `internal/config`: embedded YAML loader (`go:embed default_config.yaml`)
- [ ] **M3** — `internal/ui`: styled output, status messages, prompts
- [ ] **M4** — dry-run plumbing: shell-out runner that honours `--dry-run`; everything destructive routes through it
- [ ] **M5** — `install` job: `brew bundle` against clone's Brewfile (defensive Homebrew install first)
- [ ] **M6** — `dotfiles` job: walk `dotfiles/<program>/`, symlink per convention; back up existing first
- [ ] **M7** — `macos` job: macOS settings
- [ ] **M8** — `update` job: download newer release binary if available; stash-aware `git pull` on clone
- [ ] **M9** — `install.sh` (the curl one-liner): idempotent Homebrew install + clone repo + binary download + PATH placement
- [ ] **M10** — GitHub Actions release workflow: cross-compile arm64 + amd64, attach to release
- [ ] **M11** — README polish, optional Homebrew tap (`brew install nednella/tap/bootstrap.sh`)

`internal/utils/` doesn't get its own milestone — fs/symlink helpers land in the same commit as the job that first needs them.
