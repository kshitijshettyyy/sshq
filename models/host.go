package models

type Host struct {
	User    string `json:"user"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
