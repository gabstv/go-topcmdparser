// Parses the linux top command output into structured data.
package topcmdparser

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

type Top struct {
	Uptime    string
	Users     int
	LoadAvg5  float64
	LoadAvg10 float64
	LoadAvg15 float64
	Tasks     TasksSummary
}

type TasksSummary struct {
	Total    int
	Running  int
	Sleeping int
	Stopped  int
	Zombie   int
}

func DoString(input string) (top Top) {
	buf := bytes.NewBufferString(input)
	reader := bufio.NewReader(buf)
	insideTopCmds := false
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		if !insideTopCmds {
			if strings.HasPrefix(line, "top") {
				// row 0
				doRow0(&top, line)
			} else if strings.HasPrefix(line, "Tasks:") {
				doRow1(&top, line)
			}
		}
	}
	return
}

func doRow0(top *Top, line string) {

	rexp := regexp.MustCompile(`top[\s\-0-9:]+up([:0-9A-z,\s]+?)([0-9] user.*)`)
	matches := rexp.FindStringSubmatch(line)
	if len(matches) != 3 {
		return
	}

	top.Uptime = strings.Trim(matches[1], " ,")

	ls := strings.Split(matches[2], ",")
	if len(ls) != 4 {
		return
	}
	// users
	ls[0] = strings.TrimSpace(ls[0])
	top.Users, _ = strconv.Atoi(strings.Split(ls[0], " ")[0])
	// load avgs
	// 5
	if i := strings.Index(ls[1], "load average:"); i > -1 {
		ls[1] = strings.TrimSpace(ls[1][i+13:])
		top.LoadAvg5, _ = strconv.ParseFloat(ls[1], 64)
	}
	// 10
	top.LoadAvg10, _ = strconv.ParseFloat(strings.TrimSpace(ls[2]), 64)
	// 15
	top.LoadAvg15, _ = strconv.ParseFloat(strings.TrimSpace(ls[3]), 64)
}

func doRow1(top *Top, line string) {
	line = stripchrs(line, "0123456789,")
	ls := strings.Split(line, ",")
	if len(ls) != 5 {
		return
	}
	top.Tasks.Total, _ = strconv.Atoi(ls[0])
	top.Tasks.Running, _ = strconv.Atoi(ls[1])
	top.Tasks.Sleeping, _ = strconv.Atoi(ls[2])
	top.Tasks.Stopped, _ = strconv.Atoi(ls[3])
	top.Tasks.Zombie, _ = strconv.Atoi(ls[4])
}

func stripchrs(str, valid string) string {
	b := new(bytes.Buffer)
	validmap := make(map[rune]bool)
	for _, run := range valid {
		validmap[run] = true
	}
	for _, run := range str {
		if validmap[run] {
			b.WriteRune(run)
		}
	}
	return b.String()
}
