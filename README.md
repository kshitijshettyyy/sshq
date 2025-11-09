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

### Using Homebrew

1. Using brew tap into my repo and then brew install the sshq utility
```bash
brew tap kshitijshettyyy/sshq
brew install sshq
```

### Using git clone

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

## Prerequisites

1. Start the ssh-agent in the background
```bash
eval "$(ssh-agent -s)"
```

2. Add your private key (e.g., ~/.ssh/id_rsa)
```bash
ssh-add ~/.ssh/id_rsa
```

## Usage

1. Add a host (Provide values as prompted)
```bash
sshq add 
Alias name:
Host:
Username:
Port:
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

## Roadmap

- [x] **Homebrew install** (`brew install sshq`) 🍺
- [ ] **Encrypted password fallback** (optional) 🔐
- [ ] **Autocomplete shell integration** ⌨️
- [ ] **GUI Notch App: SSH Management** 🖥️
- [ ] **GUI Notch App: Clipboard History** 📋
- [ ] **GUI Notch App: Mac System Stats** 📊

---


## Author

Made with ❤️ by [**Kshitij Shetty**](https://github.com/kshitijshettyyy)  

