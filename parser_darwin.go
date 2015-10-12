package topcmdparser

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func DoStringDarwin(input string) (top Top) {
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
			if strings.HasPrefix(line, "Load Avg:") {
				// row 2
				doRow2darwin(&top, line)
			}
		}
	}
	return
}

func doRow2darwin(top *Top, line string) {
	line = strings.TrimSpace(line[9:])
	ls := strings.Split(line, ",")
	top.LoadAvg5, _ = strconv.ParseFloat(strings.TrimSpace(ls[0]), 64)
	if len(ls) > 1 {
		top.LoadAvg10, _ = strconv.ParseFloat(strings.TrimSpace(ls[1]), 64)
	}
	if len(ls) > 2 {
		top.LoadAvg15, _ = strconv.ParseFloat(strings.TrimSpace(ls[2]), 64)
	}
}
