package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	// Removes nops.
	{
		updated := []string{}
		for _, line := range lines {
			if line != "nop" {
				updated = append(updated, line)
			}
		}
		lines = updated
	}

loop:
	for i := 0; i < len(lines)-1; i++ {
		if strings.HasSuffix(lines[i], ":") && strings.HasSuffix(lines[i+1], ":") {
			l1, l2 := strings.TrimSuffix(lines[i], ":"), strings.TrimSuffix(lines[i+1], ":")
			updated := []string{}

			for j, line := range lines {
				if j == i {
					continue
				}

				if strings.HasSuffix(line, l1) {
					updated = append(updated, strings.TrimSuffix(line, l1)+l2)
				} else {
					updated = append(updated, line)
				}
			}

			lines = updated

			goto loop
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
