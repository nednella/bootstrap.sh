# bootstrap.sh

Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

## Features

- ⚡ **Nothing to install** — one `curl` and `bootstrap` is on your `$PATH`. No Go toolchain, no faff.
- 🍺 **Homebrew on autopilot** — grabs Homebrew if it's missing, then pours every package from your `Brewfile`.
- 🔗 **Dotfiles that write home** — your configs symlink out of the repo, so tweak one and it lands straight back in git. Commit, push, done.
- 💻 **A Mac that feels like yours** — bends a fresh machine's settings to your will in a single job.
- 🔑 **SSH in seconds** — `ssh-key` mints a git/GitHub SSH key and tells you how to grab the public half.
- 🔄 **Ages like fine wine** — `update` grabs the newest binary, `sync` pulls the latest changes.
- 👀 **No nasty surprises** — `--dry-run` shows you exactly what'll happen before anything changes.

## Install

On a fresh Mac, with nothing but `curl`:

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/bootstrap.sh | bash
```

## Usage

```sh
bootstrap                     # show help and all available commands
bootstrap dotfiles            # symlink dotfiles into $HOME / XDG
bootstrap dotfiles --undo     # unlink dotfiles and restore the latest backup
bootstrap install             # install packages from Brewfile
bootstrap macos               # apply macOS preferences
bootstrap ssh-key             # generate an SSH key
bootstrap sync                # pull the latest changes from the remote repository
bootstrap update              # update the binary to the latest release
bootstrap update --list       # list available releases
bootstrap update --tag <tag>  # install a specific release
```

```sh
-d, --dry-run         # preview any command without changing anything
-h, --help            # show help for any command
-v, --version         # print the version
```

## Build

The Go module lives in `cli/`:

```sh
make build         # → bin/bootstrap (version stamped from git)
make install-dev   # build + install a `dev` binary to /usr/local/bin
make uninstall-dev # remove the local install
```
