package models


/*
// this is the format in which the host configs are stored in the json.
example:
{
User root
Address 9.46.111.218
Port 22
}

The prot for our ssh connections is 22 by default. Need to default it in future.
*/
type Host struct {
	User    string `json:"user"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
