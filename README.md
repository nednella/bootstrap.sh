# bootstrap.sh

Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

## Features

- ⚡ **Nothing to install** — one `curl` and `bootstrap` is on your `$PATH`. No Go toolchain, no faff.
- 🍺 **Homebrew on autopilot** — grabs Homebrew if it's missing, then pours every package from your `Brewfile`.
- 🔗 **Dotfiles that write home** — your configs symlink out of the repo, so tweak one and it lands straight back in git. Commit, push, done.
- 💻 **A Mac that feels like yours** — bends a fresh machine's settings to your will in a single job.
- 🔄 **Ages like fine wine** — one command fetches the newest binary and pulls the latest config.
- 👀 **No nasty surprises** — `--dry-run` shows you exactly what'll happen before anything changes.

## Install

On a fresh Mac, with nothing but `curl`:

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/bootstrap.sh | bash
```

## Usage

```sh
bootstrap             # show help and all available commands
bootstrap dotfiles    # symlink dotfiles into $HOME / XDG_CONFIG_HOME
bootstrap install     # install packages via Brewfile
bootstrap macos       # apply macOS preferences
bootstrap update      # update the binary + pull the latest content
```

```sh
-d, --dry-run         # preview any command without changing anything
-h, --help            # show help for any command
    --version         # print the version
```

## Build

The Go module lives in `cli/`:

```sh
make build    # → bin/bootstrap (version stamped from git)
```
