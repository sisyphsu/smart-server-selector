package selector

import "testing"

func TestLoad(t *testing.T) {
	t.Log(parseServerFull("test    172.10.10.132   trade,redis,zookeeper"))

	t.Log(parseServerFull("test    172.10.10.132   22 admin  trade,redis,zookeeper"))

	t.Log(parseServerSimp("test    172.10.10.132   trade,redis,zookeeper"))

	loadServers()
}
