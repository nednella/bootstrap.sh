# bootstrap.sh

Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

## Install

On a fresh Mac, with nothing pre-installed but `curl`:

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/bootstrap.sh | bash
```

This downloads the latest `bootstrap` binary and drops it on `$PATH` (`/usr/local/bin`). The first time you run a job, `bootstrap` installs Homebrew (if missing) and clones this repo to `~/.bootstrap.sh/` — symlinks point into the clone, so edits to your config files write through to the repo and can be committed back upstream.

## Usage

```sh
bootstrap             # run all jobs (install, dotfiles, macos)
bootstrap install     # install packages from Brewfile
bootstrap dotfiles    # symlink dotfiles into $HOME / XDG
bootstrap macos       # apply macOS preferences
bootstrap update      # update binary + pull latest content
```

Preview without making changes:

```sh
bootstrap --dry-run
```
