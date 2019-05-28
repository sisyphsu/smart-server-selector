package selector

import (
	"os/user"
)

const sidebarWidth = 23

var configFile = ".sss"

func init() {
	u, _ := user.Current()
	configFile = u.HomeDir + "/" + configFile
}

var testServers = []server{
	{env: "test", host: "172.10.10.130", desc: "admin,Front,user"},
	{env: "test", host: "172.10.10.131", desc: "trade,user,mysql"},
	{env: "test", host: "172.10.10.132", desc: "trade,redis,zookeeper"},
	{env: "pre", host: "172.10.40.45", desc: "admin"},
	{env: "pre", host: "172.10.40.46", desc: "user, trade"},
	{env: "pre", host: "172.10.40.47", desc: "trade, search"},
	{env: "prod", host: "172.10.40.203", desc: "admin"},
	{env: "prod", host: "172.10.40.204", desc: "user"},
	{env: "prod", host: "172.10.40.205", desc: "user, search"},
	{env: "prod", host: "172.10.40.206", desc: "trade, search"},
	{env: "prod", host: "172.10.40.207", desc: "trade"},
}

func loadServers() []server {
	return testServers
}
