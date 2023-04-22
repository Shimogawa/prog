package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	kibi = 1024
	mebi = kibi * 1024
	gibi = mebi * 1024

	kilo = 1000
	mega = kilo * 1000
	giga = mega * 1000

	interval = 30 * time.Millisecond
)

func getSize(s float64) string {
	f := ""
	if s < kilo {
		f = "%.2f B"
	} else if s < mega {
		s = s / kibi
		f = "%.2f KiB"
	} else if s < giga {
		s = s / mebi
		f = "%.2f MiB"
	} else {
		s = s / gibi
		f = "%.2f GiB"
	}
	return fmt.Sprintf(f, s)
}

func main() {
	buf := make([]byte, 8192)
	read := int64(0)
	startTime := time.Now()
	it := 0
	lastPrint := time.Now()

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		it++
		read += int64(n)
		now := time.Now()
		if now.Sub(lastPrint) > interval {
			lastPrint = now
			dur := now.Sub(startTime)
			fmt.Fprintf(
				os.Stderr,
				"\033[2K\033[1G[prog] Elapsed %.1fs, transferred %s, avg %s/s",
				dur.Seconds(),
				getSize(float64(read)),
				getSize(float64(read)/dur.Seconds()),
			)
		}
		_, err = os.Stdout.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	dur := time.Now().Sub(startTime)
	fmt.Fprintf(os.Stderr, "\033[2K\033[1G[prog] Elapsed %s, transferred %s, avg %s/s\n", dur.String(), getSize(float64(read)), getSize(float64(read)/dur.Seconds()))
}
