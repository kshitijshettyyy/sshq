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

var vaultPath = filepath.Join(os.Getenv("HOME"), ".sshq", "config.json")

func ensureDir() error {
	return os.MkdirAll(filepath.Dir(vaultPath), 0700)
}

func Save(e Entry) error {
	if err := ensureDir(); err != nil {
		return err
	}

	all, _ := List()
	all = append(all, e)
	data, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(vaultPath, Encrypt(data), 0600)
}

func List() ([]Entry, error) {
	content, err := os.ReadFile(vaultPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Entry{}, nil
		}
		return nil, err
	}

	decrypted := Decrypt(content)
	var entries []Entry
	if err := json.Unmarshal(decrypted, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func Load(alias string) (*Entry, error) {
	entries, err := List()
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.Alias == alias {
			return &e, nil
		}
	}
	return nil, errors.New("alias not found")
}

func SaveAll(entries []Entry) error {
	if err := ensureDir(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(vaultPath, Encrypt(data), 0600)
}
