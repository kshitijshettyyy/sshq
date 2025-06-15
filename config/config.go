package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sshq/models"
)

func getConfigPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".sshq_hosts.json")
}

func LoadHosts() map[string]models.Host {
	path := getConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return make(map[string]models.Host)
	}

	var hosts map[string]models.Host
	json.Unmarshal(data, &hosts)
	return hosts
}

func SaveHost(name string, host models.Host) {
	hosts := LoadHosts()
	if hosts == nil {
		hosts = make(map[string]models.Host)
	}
	hosts[name] = host
	saveToFile(hosts)
}

func DeleteHost(name string) {
	hosts := LoadHosts()
	if _, exists := hosts[name]; exists {
		delete(hosts, name)
		saveToFile(hosts)
		fmt.Printf("🗑️  Host '%s' deleted.\n", name)
	} else {
		fmt.Printf("⚠️  Host '%s' not found.\n", name)
	}
}

func saveToFile(hosts map[string]models.Host) {
	data, _ := json.MarshalIndent(hosts, "", "  ")
	os.WriteFile(getConfigPath(), data, 0600)
}
