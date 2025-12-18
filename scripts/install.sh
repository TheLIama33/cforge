#!/bin/bash

REPO="TheLIama33/cforge"
BINARY_NAME="cforge"

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux)  OS_TYPE="linux" ;;
    Darwin) OS_TYPE="darwin" ;;
    *)      echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
    x86_64) ARCH_TYPE="amd64" ;;
    aarch64|arm64) ARCH_TYPE="arm64" ;;
    *)      echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

LATEST_URL="https://github.com/$REPO/releases/latest/download/${BINARY_NAME}_${OS_TYPE}_${ARCH_TYPE}.tar.gz"

echo "Downloading ${BINARY_NAME} for ${OS_TYPE} (${ARCH_TYPE})..."

TMP_DIR=$(mktemp -d)
curl -sL --fail "$LATEST_URL" | tar xz -C "$TMP_DIR"
if [ $? -ne 0 ]; then
    echo "Download failed. Could not find asset: $LATEST_URL"
    rm -rf "$TMP_DIR"
    exit 1
fi

if [ -w "/usr/local/bin" ]; then
    mv "$TMP_DIR/$BINARY_NAME" "/usr/local/bin/$BINARY_NAME"
else
    echo "🔑 Root privileges needed to install to /usr/local/bin"
    sudo mv "$TMP_DIR/$BINARY_NAME" "/usr/local/bin/$BINARY_NAME"
fi

rm -rf "$TMP_DIR"

echo "Installation complete! Try running: $BINARY_NAME --version"