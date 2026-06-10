# bootstrap.sh

A Go CLI that bootstraps a fresh macOS install — Homebrew + Brewfile, dotfiles, macOS preferences — in one command, on a machine that has only `curl`. Apple Silicon only.

---

## How it works

Monorepo: this repo holds both the CLI's Go source (`cli/`) and the dotfiles it manages (`dotfiles/`). The binary is distributed as a prebuilt artefact on GitHub Releases (`bootstrap-darwin-arm64`).

There is **no self-location logic** — the binary always operates against a fixed **install location**, `~/.bootstrap.sh` (the embedded `install_path`), regardless of where the binary or the dev clone actually live. So the dev repo can sit anywhere (e.g. `~/dev/bootstrap.sh`); the install location is reserved for the clone the tool manages.

### Install (fresh Mac)

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/bootstrap.sh | bash
```

`bootstrap.sh` is **light**: it downloads the latest `bootstrap-darwin-arm64` (via the `releases/latest/download/` redirect) and drops it at `/usr/local/bin/bootstrap` (sudo only if that dir isn't writable). It does **not** install Homebrew or clone the repo — that's `preflight`'s job.

### Runtime

Every **job command** runs `preflight` first (root `PersistentPreRun`, gated to job commands by Cobra `GroupID`). Preflight ensures — silently, acting only on what's missing — Homebrew → git → the repo clone at `~/.bootstrap.sh`. Then the job runs against that clone:

- `bootstrap install` — `brew bundle` against `<clone>/Brewfile`.
- `bootstrap dotfiles` — symlink `<clone>/dotfiles/` into `$HOME` / `$XDG_CONFIG_HOME` (existing files backed up to `~/.dotfiles-backup/<timestamp>/` first).
- `bootstrap macos` — apply macOS `defaults` read from `<clone>/macos/settings.yaml`.
- `bootstrap ssh-key` — generate an ed25519 SSH key (comment = `git config user.email`, which it requires) and copy the public key to the clipboard via `pbcopy`. Git/GitHub-specific, not a generic keygen.

Symlinks point into the clone, so editing a config file writes through to the repo — commit + push upstream with normal git from `~/.bootstrap.sh`.

### Update & sync

Two independent maintenance commands — different cadences, no shared state:

- `bootstrap update` — **binary only**. Fetch the latest release tag (GitHub API), `semver.Compare` it against the running binary's version; if newer, download + atomically replace `/usr/local/bin/bootstrap` (sudo, primed once via `PromptSudo`). If current, it's a no-op. `--tag <tag>` / `-t` installs a specific release instead (downgrade confirmed via `utils.Confirm`, but no newer-only guard); `--list` / `-l` prints the available release tags, newest first, marking the current one. The shared download+swap is `replaceBinary`; `/releases/latest` (stable) feeds the no-arg path, `/releases` (all) feeds `--list`.
- `bootstrap sync` — **content only**. `git pull --rebase --autostash` on the clone, so local dotfile edits are stashed, the pull replays on top, and they're restored — surviving the update.

Both are job commands too, so `preflight` runs first — it guarantees the clone exists before `sync` pulls, and is a harmless no-op for `update` (which only touches the binary).

### Versioning

The binary knows its own version via **ldflags injection** — `cli/internal/version.go` (`package internal; var Version = "dev"`). Release builds stamp the git tag (`-X …/internal.Version=<tag>`); local builds show `dev` (or a real `git describe` version via `make build`). Surfaced by `bootstrap --version` and the banner; consumed by `update`'s semver check. No version constant is committed — the git tag is the single source of truth.

### Release (release-please)

- `release-please.yml` (on push to `main`) maintains a release PR — changelog + next version derived from conventional commits — authored with the **`RELEASE_PLEASE_TOKEN`** PAT. Merging the PR tags + cuts a GitHub Release.
- `upload-binary-to-release.yml` (on `release: published`) runs `make release VERSION=<tag>` to cross-compile `bootstrap-darwin-arm64` (the `VERSION=` override stamps the exact release tag) and attaches it. The PAT is what lets the release event trigger this second workflow.
- To ship: push conventional commits → merge the release PR (rebase). `feat:`/`fix:` drive a release; **content-only changes** (dotfiles, Brewfile, macOS settings) use `chore:`/`docs:` so they propagate via `sync`'s `git pull` without cutting a needless new binary.

---

## Repository layout

```
bootstrap.sh/
├── bootstrap.sh         # the curl one-liner: downloads + installs the binary
├── Brewfile             # consumed by `bootstrap install`
├── Makefile             # build, run, release, local install-dev tasks
├── macos/               # settings.yaml — consumed by `bootstrap macos`
├── dotfiles/            # consumed by `bootstrap dotfiles` (ghostty, git, starship, zsh)
├── .github/workflows/   # release-please.yml + upload-binary-to-release.yml
├── CLAUDE.md  README.md  CHANGELOG.md
└── cli/                 # the Go tool (go.mod lives here — build from cli/)
    ├── main.go          # thin: wires Cobra, defers to cmd/
    ├── cmd/             # one file per command: parse → call internal/jobs  (interface layer)
    └── internal/
        ├── version.go   # package internal: var Version (ldflags-injected)
        ├── config/      # embedded default_config.yaml + loader
        ├── jobs/        # one file per job: preflight, install, dotfiles, macos, sshkey, update, sync  (logic layer)
        ├── ui/          # banner + styled logging
        └── utils/       # dry-run-aware shell-out runner, fs/symlink/exec/yaml helpers
```

`cmd/` is the *interface* layer (Cobra); `internal/jobs/` is the *logic* layer — same shape as API endpoint ↔ service.

---

## Configuration

`cli/internal/config/default_config.yaml` is embedded via `//go:embed` — the source of truth for `install_path` (`~/.bootstrap.sh`), `backup_path`, `repo_url`. No runtime config overlay; change a default in source and release a new binary.

---

## Dotfiles convention

Flat tree, **one directory per program**. Within a program directory the symlink destination follows a single rule:

- File starts with `.` → symlinked into `$HOME`.
- Otherwise → `${XDG_CONFIG_HOME:-$HOME/.config}/<program>/`.

| Source | Destination |
|---|---|
| `dotfiles/zsh/.zshrc` | `~/.zshrc` |
| `dotfiles/git/.gitconfig` | `~/.gitconfig` |
| `dotfiles/ghostty/config` | `~/.config/ghostty/config` |
| `dotfiles/starship/starship.toml` | `~/.config/starship/starship.toml` |

The dot prefix carries the intent — no metadata files. The walk skips non-directories, so a loose file in `dotfiles/` would be ignored (which is why the `Brewfile` lives at the repo root, not here).

---

## Command surface

```
bootstrap install             # install packages from Brewfile
bootstrap dotfiles            # symlink dotfiles into $HOME / XDG
bootstrap macos               # apply macOS preferences
bootstrap ssh-key             # generate an SSH key and copy it to the clipboard
bootstrap update              # update the binary to the latest release
bootstrap update --list       # list available releases
bootstrap update --tag <tag>  # install a specific release
bootstrap sync                # pull the latest changes from the remote repository
bootstrap --version           # print the version
--dry-run / -d                # global flag: print what would happen, change nothing
```

Bare `bootstrap` prints Cobra's help — **by design**. Every job is **atomic and order-independent** — run any one on its own, in any order — so there's intentionally no run-all entry point and no prescribed sequence; each job is invoked explicitly.

---

## Build & development

`go.mod` lives in `cli/`, so build/run from there — or use the Makefile, which stamps the version from `git describe` (a plain `go build`, or `make install-dev`, shows `dev`):

```sh
make build               # → bin/bootstrap, version stamped
make install-dev         # build (stamped `dev`) + install to /usr/local/bin/bootstrap
make uninstall-dev       # remove the local install
make run ARGS="macos -d" # run from source with the version stamped
cd cli && go build ./... # plain compile check
```

`--dry-run` previews any job without touching the system — everything destructive routes through `utils.Command`, which prints the command instead of running it when the flag is set.

---

## Ideas (future work)

Not built yet — a rough backlog, unordered within each group.

**Reversibility** (agreed direction)
- `bootstrap dotfiles --undo` / `-u` — reverse the symlinks and restore originals from the latest `~/.dotfiles-backup/<timestamp>/`. The backup mirrors each file's `$HOME`-relative path, so the restore is a clean inverse walk. Must run atomically, like every job. Also makes lifecycle testing trivial.

**UX polish** (vague, low priority)
- Replace some raw stdout (`git pull`, `brew bundle`, the binary download) with loaders / spinners / progress — no specifics yet, just "might be nicer." If pursued: gate behind `term.IsTerminal`, and don't let a spinner swallow the sudo prompt.

**Interactivity** (vague, low priority)
- Interactive selection where a flag currently takes an explicit value: pick a version from the `update --list` output instead of typing `--tag`, or select specific apps/dotfiles to install/symlink rather than doing the whole set. Same caveats as the spinners above — gate behind `term.IsTerminal` so non-TTY/CI runs stay non-interactive, and keep the explicit-flag path as the scriptable default (the prompt is a convenience layer over it, never the only way in). Atomicity still holds: the selection just feeds the same job.

---

## Conventions

### Terminology
The unit of work is always a **job**. Never "phase", "step", or "task".

### Language
British English (en-GB) in commits, comments, prose, and identifiers I control (`colour`, `behaviour`, `initialise`). Standard Go stdlib names stay American (`color.Color`, etc.) — not mine to control.

### Go style
Never write `if err := foo(); err != nil { ... }`. Split the assignment onto its own line:

```go
err := foo()
if err != nil {
    return err
}
```

### Declaration order
A function called by a neighbouring function — a helper in a local call chain — sits in execution (reading) order, caller above callee (see `cli/internal/jobs/preflight.go`, `cli/internal/jobs/dotfiles.go`). A function only ever called from elsewhere — an independent primitive or entry point — is ordered alphabetically (see `cli/internal/utils/os.go`).

### Casing
Standard Go: exported `PascalCase`, unexported `camelCase`, initialisms all-caps (`URL`, `ID`, `HTTP`). Never export a helper purely so another package can reach it — that breaks the casing consistency of its sibling helpers. Route the cross-package call through the package's public entry point instead.

### Comments
Code should be self-descriptive — reach for clearer names before a comment. Skip comments that restate what the code already says. Keep load-bearing comments that explain non-obvious behaviour, ordering, or *why* a decision was made — something no name can carry (see the `//go:embed` note in `cli/internal/config/config.go`). Compiler/tooling directives (`//go:embed`, release-please markers) are not comments and stay.

### Style var naming
Prefix style vars with the feature that consumes them: `headerArrowStyle`, not `arrowStyle`.

### Commit discipline
"and" / "&" in a commit subject means split into multiple commits — one logical change per commit. Subject-only; no body unless genuinely load-bearing. Before committing, check what's actually staged (`git diff --cached`) — `git mv` pre-stages renames, which is easy to sweep into the wrong commit.

### Command framing
Describe commands (Cobra `Short`/`Long`, README, command-surface lines) by the **primary action**. Prerequisites and defensive checks are implementation detail, not user-facing copy.
