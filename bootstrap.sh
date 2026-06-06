#!/usr/bin/env bash
set -euo pipefail

REPO="nednella/bootstrap.sh"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="bootstrap"
BINARY_PATH="$INSTALL_DIR/$BINARY_NAME"

ASSET="bootstrap-darwin-arm64"
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$ASSET"

echo "==> Downloading latest release ..."

TMP_FILE=$(mktemp)
trap 'rm -f "$TMP_FILE"' EXIT

if ! curl -fL "$DOWNLOAD_URL" -o "$TMP_FILE"; then
    echo "Error: could not find latest release at $DOWNLOAD_URL" >&2
    exit 1
fi

chmod +x "$TMP_FILE"
xattr -d com.apple.quarantine "$TMP_FILE" 2>/dev/null || true

echo "==> Installing $BINARY_NAME to $INSTALL_DIR (requires sudo) ..."

sudo mkdir -p "$INSTALL_DIR"
sudo install -m 755 "$TMP_FILE" "$BINARY_PATH"

echo "Installation complete! Run '$BINARY_NAME' to get started."
