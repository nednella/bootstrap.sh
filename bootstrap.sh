#!/usr/bin/env bash
set -euo pipefail

REPO="nednella/bootstrap.sh"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="bootstrap"
BINARY_PATH="$INSTALL_DIR/$BINARY_NAME"

ASSET="bootstrap-darwin-arm64"
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$ASSET"

if [ -t 1 ]; then
    GREEN=$'\033[32m'
    RED=$'\033[31m'
    ORANGE=$'\033[91m'
    DIM=$'\033[2m'
    RESET=$'\033[0m'
else
    GREEN="" RED="" ORANGE="" DIM="" RESET=""
fi

echo "Setting up $BINARY_NAME..."

TMP_FILE=$(mktemp)
trap 'rm -f "$TMP_FILE"' EXIT

if ! curl -fsSL "$DOWNLOAD_URL" -o "$TMP_FILE"; then
    echo "${RED}✖ Could not download the latest release from $DOWNLOAD_URL${RESET}" >&2
    exit 1
fi

chmod +x "$TMP_FILE"

echo "${DIM}Installing to $INSTALL_DIR requires sudo.${RESET}"
sudo mkdir -p "$INSTALL_DIR"
sudo install -m 755 "$TMP_FILE" "$BINARY_PATH"

VERSION=$("$BINARY_PATH" --version | awk '{print $NF}')

echo
echo "${GREEN}✔ $BINARY_NAME successfully installed!${RESET}"
echo
echo "  ${DIM}Version:${RESET} ${ORANGE}$VERSION${RESET}"
echo
echo "  ${DIM}Location:${RESET} $BINARY_PATH"
echo
echo "  ${DIM}Run ${RESET}${ORANGE}$BINARY_NAME${RESET}${DIM} to get started${RESET}"
echo
