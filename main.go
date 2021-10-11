package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	sourceEnv := flag.String("source", ".env", "flag -source for your source env. ex: -source:example.env")
	targetEnv := flag.String("target", "", "flag -target for your target env. ex: -target:example-another.env")
	flag.Parse()

	if *sourceEnv == "" || *targetEnv == "" {
		flag.Usage()
		return
	}

	sourceData, err := parsePath(*sourceEnv)
	if err != nil {
		log.Println(err)
	}

	targetData, err := parsePath(*targetEnv)
	if err != nil {
		log.Println(err)
	}

	output := checkDiff(sourceData, targetData)
	fmt.Println("Difference of Key Environment:\n", strings.Join(output[0], "\n"))
	fmt.Println("Difference of Key-Value Environment:\n", strings.Join(output[1], "\n"))
}
