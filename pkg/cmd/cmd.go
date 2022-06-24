package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/so-jelly/primes/pkg/primes"
)

var DefaultFile string = "input.txt"

func PrimesCmd(i int) {
	p := &primes.Primes{Number: &i}
	p.GetPrimes()
	out, _ := json.Marshal(p)
	fmt.Print(string(out))
}

// IsIntSet checks if the int flag is set
func IsIntSet() bool {
	var set bool
	flag.Visit(
		func(intFlag *flag.Flag) {
			if intFlag.Name == "int" {
				set = true
			}
		})
	return set
}

func ReadIntFile(f string) int {
	fileData, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("int not set and can't read input file")
	}
	fileInt, err := strconv.Atoi(strings.TrimSpace(string(fileData)))
	if err != nil {
		log.Fatalln("int not set and file data is not a single integer")
	}
	return fileInt
}
