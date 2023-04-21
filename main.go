package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	kilo = 1024
	mega = kilo * 1024
	giga = mega * 1024
)

func getSize(s float64) string {
	f := ""
	if s < kilo {
		f = "%.2f B"
	} else if s < mega {
		s = s / kilo
		f = "%.2f KiB"
	} else if s < giga {
		s = s / mega
		f = "%.2f MiB"
	} else {
		s = s / giga
		f = "%.2f GiB"
	}
	return fmt.Sprintf(f, s)
}

func main() {
	buf := make([]byte, 8192)
	read := int64(0)
	startTime := time.Now()
	it := 0

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
		if it%7000 == 0 {
			dur := time.Now().Sub(startTime)
			fmt.Fprintf(os.Stderr, "\033[2K\033[1G[prog] Transferred %s, avg %s/s", getSize(float64(read)), getSize(float64(read)/dur.Seconds()))
		}
		_, err = os.Stdout.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	dur := time.Now().Sub(startTime)
	fmt.Fprintf(os.Stderr, "\033[2K\033[1G[prog] Transferred %s, avg %s/s\n", getSize(float64(read)), getSize(float64(read)/dur.Seconds()))
}
