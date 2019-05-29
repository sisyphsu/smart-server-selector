package selector

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type server struct {
	env   string
	host  string
	port  string
	user  string
	desc  string
	score int
}

type serverArray []server

func (a serverArray) Len() int      { return len(a) }
func (a serverArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a serverArray) Less(i, j int) bool {
	if a[i].score != a[j].score {
		return a[i].score < a[j].score
	}
	if a[i].env != a[j].env {
		return a[i].env < a[j].env
	}
	if a[i].host != a[j].host {
		return a[i].host < a[j].host
	}
	return i < j
}

// load servers from config file.
func loadServers() (arr []server) {
	arr = make([]server, 0)
	fs, _ := ioutil.ReadFile(SssFile)
	if len(fs) == 0 {
		return
	}
	body := string(fs)
	var errs []string
	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' || line[0] == '/' {
			continue
		}
		s := parseServerFull(line)
		if s == nil {
			s = parseServerSimp(line)
		}
		if s != nil {
			arr = append(arr, *s)
		} else {
			errs = append(errs, line)
		}
	}
	if len(errs) > 0 {
		fmt.Printf("some invalid config in file[%v]: \n", SssFile)
		for _, e := range errs {
			println("> " + e)
		}
		println("press any key to continue")
		getchar()
	}
	return
}

var fullPtn = regexp.MustCompile("^(\\w+)\\s+([\\w.]+)\\s+(\\d+)\\s+([\\w.]+)\\s+(.*)$")

// parse server by full pattern
func parseServerFull(s string) *server {
	sm := fullPtn.FindStringSubmatch(s)
	if len(sm) == 0 || len(sm) != 6 {
		return nil
	}
	return &server{
		env:  sm[1],
		host: sm[2],
		port: sm[3],
		user: sm[4],
		desc: sm[5],
	}
}

var simpPtn = regexp.MustCompile("^(\\w+)\\s+([\\w.]+)\\s+(.*)$")

// parse server by simple pattern
func parseServerSimp(s string) *server {
	sm := simpPtn.FindStringSubmatch(s)
	if len(sm) == 0 || len(sm) != 4 {
		return nil
	}
	return &server{
		env:  sm[1],
		host: sm[2],
		port: "",
		user: "",
		desc: sm[3],
	}
}
