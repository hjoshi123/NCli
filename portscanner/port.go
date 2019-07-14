package portscanner

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ttacon/chalk"
	"golang.org/x/sync/semaphore"
)

// PortScanner is an abstraction containing the IP and no of goroutines to be allowed to run
type PortScanner struct {
	IP   string
	Lock *semaphore.Weighted
}

// Ulimit calculates the limit for the Lock function
func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()

	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))
	limit, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return limit
}

// ScanPort scans the port with the given IP and the port
func ScanPort(IP string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", IP, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open connections") {
			time.Sleep(timeout)
			ScanPort(IP, port, timeout)
		} else {
			fmt.Println(chalk.Cyan, port, chalk.Reset, " closed")
		}
		return
	}
	lime := chalk.Green.NewStyle().
		WithTextStyle(chalk.Bold)

	conn.Close()
	fmt.Printf("%s%d%s%s\n", lime, port, " open", chalk.Reset)
}

// Start starts the actual connection attempts
func (ps *PortScanner) Start(firstPort, lastPort int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := firstPort; port <= lastPort; port++ {
		wg.Add(1)
		ps.Lock.Acquire(context.TODO(), 1)

		// anonymous go routine defer is used to check the output of ScanPort
		go func(port int) {
			defer ps.Lock.Release(1)
			defer wg.Done()
			ScanPort(ps.IP, port, timeout)
		}(port)
	}
}
