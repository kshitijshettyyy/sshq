# sshq — SSH Simplified for Devs

**sshq** is a macOS-native command-line utility to manage SSH connections. It lets developers save, connect, list, and delete remote VM SSH configs from a single CLI, improving daily productivity for DevOps and backend engineers.

---

## Features

- 🔐 Add and save SSH host credentials
- ⚡ One-click SSH to remote machines via alias
- 🔑 Automatically installs your SSH public key using `ssh-copy-id`
- 📋 List all saved SSH connections
- 🗑️ Delete unused or outdated hosts
- 💾 Stores config data in `~/.sshq_hosts.json`
- 🧩 Written in Go — fast, lightweight, cross-platform friendly

---

## Prerequisites

Before using **sshq**, ensure your system has the following:

| Tool          | How to Install                                |
|---------------|------------------------------------------------|
| Go (>=1.18)   | [golang.org](https://golang.org/dl/) or `brew install go` |
| OpenSSH       | Pre-installed on macOS. If missing: `brew install openssh` |

To verify installation:

```bash
go version
ssh -V
```

## Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/sshq.git
cd sshq
```
2. Build the app (This creates a binary for our executions)
```bash
go build -o sshq
```
3. Make it globally accessible
```bash
sudo mv sshq /usr/local/bin/
```

## Usage

1. Add a host
```bash
sshq add <alias> <username@host> -p <port>
```
2. Connect to a host
```bash
sshq connect <alias>
```
3. List all hosts
```bash
sshq list
```
4. Delete a host
```bash
sshq delete <alias>
```


---

## Coming Soon

- 🔐 **Encrypted password fallback** (optional)
- ⌨️ **Autocomplete shell integration**
- 🍺 **Homebrew install** (`brew install sshq`)
- 🖥️ **GUI Notch App** with:
  - SSH management
  - Clipboard history
  - Mac system stats

---

## Author

Made with ❤️ by [**Kshitij Shetty**](https://github.com/kshitijshettyyy)  

