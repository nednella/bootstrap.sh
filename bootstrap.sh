#!/usr/bin/env bash
set -euo pipefail

REPO="nednella/bootstrap.sh"
INSTALL_DIR="/usr/local/bin"
BINARY_PATH="$INSTALL_DIR/bootstrap"

ASSET="bootstrap-darwin-arm64"
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$ASSET"

echo "==> Downloading latest release..."

TMP_FILE=$(mktemp)
if ! curl -fL "$DOWNLOAD_URL" -o "$TMP_FILE"; then
    echo "Error: could not find latest release at $DOWNLOAD_URL" >&2
    exit 1
fi

chmod +x "$TMP_FILE"
xattr -d com.apple.quarantine "$TMP_FILE" 2>/dev/null || true

if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_FILE" "$BINARY_PATH"
else
    echo "==> Installing to $INSTALL_DIR (requires sudo)"
    sudo mv "$TMP_FILE" "$BINARY_PATH"
fi

echo "Installation complete! Run 'bootstrap' to get started."
