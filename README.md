# ContextForge (cforge)

<div align="center">

![GitHub release (latest by date)](https://img.shields.io/github/v/release/TheLIama33/cforge?style=for-the-badge)
![Go Version](https://img.shields.io/github/go-mod/go-version/TheLIama33/cforge?style=for-the-badge)
![License](https://img.shields.io/github/license/TheLIama33/cforge?style=for-the-badge)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20windows%20%7C%20macos-lightgrey?style=for-the-badge)

<p align="center">
  <strong>The ultimate CLI tool to bridge your codebase and Large Language Models.</strong>
</p>

</div>

---

## 📖 Overview

**ContextForge** (`cforge`) is a high-performance, system-wide CLI tool designed to streamline the workflow between software development and AI assistance.

Instead of manually copying content from dozens of files, `cforge` scans your project, intelligently filters irrelevant noise (like `node_modules` or build artifacts), and delivers a formatted, token-optimized context string directly to your clipboard.

### Why ContextForge?

- **⚡ Lightning Fast:** Built in Go, it scans thousands of files in milliseconds.
- **📋 Clipboard First:** No intermediate files. Output goes straight to RAM/Clipboard.
- **🧠 Context Aware:** Respects `.gitignore` and allows for custom profiles (e.g., "Backend Only").
- **🐧 Cross-Platform:** Native binaries for Linux, Windows, and macOS.

---

## 📥 Installation

Choose your operating system to install **ContextForge** in seconds.

<table>
  <tr>
    <td align="center" width="120">
      <img src="https://upload.wikimedia.org/wikipedia/commons/8/87/Windows_logo_-_2021.svg" width="45"><br>
      <b>Windows</b>
    </td>
    <td align="left">
      Run in <b>PowerShell</b>:<br>
      <code>iwr https://raw.githubusercontent.com/TheLIama33/cforge/main/scripts/install.ps1 -useb | iex</code>
    </td>
  </tr>
  <tr>
    <td align="center" width="120">
      <img src="https://upload.wikimedia.org/wikipedia/commons/1/1b/Apple_logo_grey.svg" width="30"><br>
      <b>macOS</b>
    </td>
    <td align="left">
      Run in <b>Terminal</b>:<br>
      <code>curl -sL https://raw.githubusercontent.com/TheLIama33/cforge/main/scripts/install.sh | bash</code>
    </td>
  </tr>
  <tr>
    <td align="center" width="120">
      <img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" width="35"><br>
      <b>Linux</b>
    </td>
    <td align="left">
      Run in <b>Terminal</b>:<br>
      <code>curl -sL https://raw.githubusercontent.com/TheLIama33/cforge/main/scripts/install.sh | bash</code>
    </td>
  </tr>
</table>

### Alternative Methods

**Go Install:**

```bash
go install github.com/TheLIama33/cforge@latest
```

**Manual Binary:**
Download the latest pre-compiled binary from the [Releases Page](https://github.com/TheLIama33/cforge/releases).

---

## 🚀 Quick Start

Once installed, `cforge` is available globally in your terminal.

### 1. Initialize (Optional)

Create a default configuration file in your project root.

```bash
cforge init
```

### 2. Run (Default)

Scans the current directory using standard exclusions (git, node_modules, etc.) and copies the result to your clipboard.

```bash
cforge
```

_> Output: [+] Copied 42 files to system clipboard._

### 3. Run with Profile

Use a specific profile defined in your config (e.g., specific to frontend code).

```bash
cforge --profile frontend
```

### 4. Dry Run / Pipe

Print to stdout instead of clipboard (useful for piping to files or other tools).

```bash
cforge --stdout > context.md
```

### 5. Update

ContextForge includes a self-update mechanism.

```bash
cforge update
```

---

## ⚙️ Configuration

`cforge` looks for a `.cforge.json` file in the current directory (or your home directory as a fallback).

```json
{
  "global": {
    "copyToClipboard": true,
    "showTokenCount": true,
    "defaultProfile": "default",
    "useGitIgnore": true,
    "formatting": "markdown"
  },
  "profiles": {
    "default": {
      "includePatterns": ["*.go", "*.js", "*.ts", "*.md", "Dockerfile"],
      "excludePatterns": ["*_test.go", "*.spec.ts"],
      "excludePaths": ["node_modules", "dist", "vendor", ".git"]
    },
    "backend": {
      "includePaths": ["cmd/", "internal/", "pkg/"],
      "excludeFiles": ["internal/secrets.go"]
    }
  }
}
```

| Field             | Description                                                     |
| :---------------- | :-------------------------------------------------------------- |
| `includePatterns` | Whitelist files by glob pattern (e.g. `*.ts`).                  |
| `includePaths`    | Whitelist specific directories. If set, only these are scanned. |
| `excludePaths`    | Blacklist directories (performance optimization).               |
| `useGitIgnore`    | If true, automatically skips files listed in `.gitignore`.      |

---

## 🛠️ Development

Requirements: Go 1.21+

1.  **Clone the repo**

    ```bash
    git clone https://github.com/TheLIama33/cforge.git
    cd cforge
    ```

2.  **Build**

    ```bash
    go build -o cforge ./cmd/cforge
    ```

3.  **Test**
    ```bash
    ./cforge --version
    ```

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---

<div align="center">
  <sub>Built with ❤️ in Go.</sub>
</div>
