package vault

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Entry struct {
	Alias    string `json:"alias"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Method   string `json:"method"`
	Password string `json:"password,omitempty"`
	PEMFile  string `json:"pem_file,omitempty"`
}

var vaultFile = filepath.Join(os.Getenv("HOME"), ".sshq_vault.json")

// LoadAll reads file and returns all entries as a map
func LoadAll() (map[string]Entry, error) {
	data := make(map[string]Entry)

	file, err := os.ReadFile(vaultFile)
	if err != nil {
		// Return empty map if file doesn't exist
		return data, nil
	}

	if len(file) == 0 {
		return data, nil
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SaveAll writes back full map to file
func SaveAll(data map[string]Entry) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(vaultFile, b, 0600)
}

// Load loads a single entry by alias
func Load(alias string) (Entry, error) {
	entries, err := LoadAll()
	if err != nil {
		return Entry{}, err
	}

	entry, exists := entries[alias]
	if !exists {
		return Entry{}, errors.New("alias not found")
	}

	return entry, nil
}

// Save inserts/updates a single entry
func Save(entry Entry) error {
	entries, err := LoadAll()
	if err != nil {
		entries = make(map[string]Entry)
	}
	// 👇 FIX: map is always initialized
	if entries == nil {
		entries = make(map[string]Entry)
	}

	entries[entry.Alias] = entry
	return SaveAll(entries)
}
