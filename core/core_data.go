package core

var cxt context

type server struct {
	env  string
	host string
	desc string
}

type context struct {
	keyword        string
	serverSelected string
	serverVisible  []server
	serverAll      []server
}

// build table by servers and keyword
func (c *context) buildTableData() (data [][]string) {
	for _, server := range c.serverAll {
		data = append(data, []string{server.env, server.host, server.desc})
	}
	return
}

// reload all servers from configuration file.
func (c *context) reload() {

}

func init() {
	cxt = context{
		serverAll: []server{
			{"test", "172.10.10.130", "admin,front,user"},
			{"test", "172.10.10.131", "trade,user,mysql"},
			{"test", "172.10.10.132", "trade,redis,zookeeper"},
			{"pre", "172.10.40.45", "admin"},
			{"pre", "172.10.40.46", "user, trade"},
			{"pre", "172.10.40.47", "trade, search"},
			{"prod", "172.10.40.203", "admin"},
			{"prod", "172.10.40.204", "user"},
			{"prod", "172.10.40.205", "user, search"},
			{"prod", "172.10.40.206", "trade, search"},
			{"prod", "172.10.40.207", "trade"},
		},
	}
	cxt.reload()
}
