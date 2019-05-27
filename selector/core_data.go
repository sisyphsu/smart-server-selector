package selector

var cxt context

type context struct {
	keyword        string
	serverSelected string
	serverVisible  []server
	serverAll      []server
}

// build table by serverTable and keyword
func (c *context) buildTableData() (data [][]string) {
	for _, server := range c.serverAll {
		data = append(data, []string{server.env, server.host, server.desc})
	}
	return
}

// reload all serverTable from configuration file.
func (c *context) reload() {

}

func loadServers() []server {
	return testServers
}
