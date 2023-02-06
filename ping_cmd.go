package safedns

import (
	"os/exec"
	"time"
)

func PingCmd(ipAddr string) (avgTime time.Duration, err error) {
	t1 := time.Now()
	c := exec.Command("ping", "-c1", "-w1", ipAddr)
	c.Stdout = nil
	c.Stderr = nil
	err = c.Run()
	avgTime = time.Since(t1)
	return
}
