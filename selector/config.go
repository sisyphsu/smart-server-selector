package selector

import "os/user"

var SssFile = ".servers"
var SshOptions = "-o ConnectTimeout=10"
var SidebarWidth = 23

func init() {
	u, _ := user.Current()
	SssFile = u.HomeDir + "/" + SssFile
}
