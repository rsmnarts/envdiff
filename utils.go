package main

import (
	"os"
	"strings"
)

func parsePath(path string) (out [][]string, err error) {
	sourceData, err := os.ReadFile(path)
	if err != nil {
		return out, err
	}

	lines := strings.Split(strings.ReplaceAll(string(sourceData), "\r\n", "\n"), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			out = append(out, strings.Split(line, "="))
		}
	}

	return
}

func checkDiff(data ...[][]string) (out [][]string) {
	out = make([][]string, len(data))
	for _, d := range data[0] {
		for i, d2 := range data[1] {
			if d[0] == d2[0] {
				if d[1] != d2[1] {
					out[1] = append(out[1], strings.Join(d2, "="))
				}
				break
			}

			if d[0] != d2[0] && i == len(data[1])-1 {
				out[0] = append(out[0], d[0])
			}
		}
	}
	return
}
