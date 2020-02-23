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

	// Remove NOPs.
	{
		updated := []string{}
		for _, line := range lines {
			if line != "NOP" {
				updated = append(updated, line)
			}
		}
		lines = updated
	}

	for {
		updated := false
		for i := 0; i < len(lines)-1; i++ {
			if strings.HasSuffix(lines[i], ":") && strings.HasSuffix(lines[i+1], ":") {
				l1, l2 := strings.TrimSuffix(lines[i], ":"), strings.TrimSuffix(lines[i+1], ":")
				updatedLines := []string{}

				for j, line := range lines {
					if j == i {
						continue
					}

					if strings.HasSuffix(line, l1) {
						updatedLines = append(updatedLines, strings.TrimSuffix(line, l1)+l2)
					} else {
						updatedLines = append(updatedLines, line)
					}
				}

				lines = updatedLines
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}
