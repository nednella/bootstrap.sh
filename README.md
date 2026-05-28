# bootstrap.sh

Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

## Install

On a fresh Mac, with nothing pre-installed but `curl`:

```sh
curl -fsSL https://raw.githubusercontent.com/nednella/bootstrap.sh/main/install.sh | bash
```

This installs Homebrew (if missing), clones this repo to `~/.bootstrap.sh/`, and drops the `bootstrap` binary on `$PATH`. Symlinks point into the clone, so edits to your config files write through to the repo and can be committed back upstream.

## Usage

```sh
bootstrap             # run all jobs (install, dotfiles, macos)
bootstrap install     # Homebrew + Brewfile only
bootstrap dotfiles    # symlink dotfiles only
bootstrap macos       # apply macOS preferences only
bootstrap update      # update binary + pull latest content (stash-safe)
```

Preview without making changes:

```sh
bootstrap --dry-run
```
