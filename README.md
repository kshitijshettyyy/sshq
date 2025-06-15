# sshq — SSH Simplified for Devs

**sshq** is a macOS-native command-line utility to manage SSH connections. It lets developers save, connect, list, and delete remote VM SSH configs from a single CLI, improving daily productivity for DevOps and backend engineers.

---

## ✨ Features

- 🔐 Add and save SSH host credentials
- ⚡ One-click SSH to remote machines via alias
- 🔑 Automatically installs your SSH public key using `ssh-copy-id`
- 📋 List all saved SSH connections
- 🗑️ Delete unused or outdated hosts
- 💾 Stores config data in `~/.sshq_hosts.json`
- 🧩 Written in Go — fast, lightweight, cross-platform friendly

---

## 💻 Prerequisites

Before using **sshq**, ensure your system has the following:

| Tool          | How to Install                                |
|---------------|------------------------------------------------|
| Go (>=1.18)   | [golang.org](https://golang.org/dl/) or `brew install go` |
| OpenSSH       | Pre-installed on macOS. If missing: `brew install openssh` |

To verify installation:

```bash
go version
ssh -V
