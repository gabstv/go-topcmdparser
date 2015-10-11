package topcmdparser

import (
	"testing"
)

func TestParser(t *testing.T) {
	str := `
top - 19:00:03 up  5:57,  1 user,  load average: 1.06, 1.17, 1.31
Tasks:  99 total,   3 running,  96 sleeping,   0 stopped,   0 zombie
%Cpu(s): 35.6 us,  0.5 sy,  0.0 ni,  0.0 id,  0.0 wa,  0.0 hi,  0.0 si, 63.8 st
KiB Mem:   2048484 total,  1273464 used,   775020 free,    41988 buffers
KiB Swap:  4194300 total,        0 used,  4194300 free.  1018108 cached Mem

  PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
 1810 root      20   0  181820 101624  30476 R 99.8  5.0 123:41.49 app
    1 root      20   0   33644   3008   1496 S  0.0  0.1   0:02.25 init
    2 root      20   0       0      0      0 S  0.0  0.0   0:00.00 kthreadd
    3 root      20   0       0      0      0 S  0.0  0.0   0:02.02 ksoftirqd/0`
	expected := Top{}
	expected.Users = 1
	expected.Uptime = "5:57"
	expected.LoadAvg5 = 1.06
	expected.LoadAvg10 = 1.17
	expected.LoadAvg15 = 1.31
	expected.Tasks.Total = 99
	expected.Tasks.Running = 3
	expected.Tasks.Sleeping = 96
	expected.Tasks.Stopped = 0
	expected.Tasks.Zombie = 0
	got := DoString(str)
	if expected != got {
		t.Fatalf("Expected:\n %v \n\nbut got: %v\n", expected, got)
	}
}
