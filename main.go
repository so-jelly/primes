package main

import (
	"flag"

	"github.com/so-jelly/primes/pkg/cmd"
	"github.com/so-jelly/primes/pkg/serve"
)

func main() {
	i := flag.Int("int", 0, "integer to check for primes")
	f := flag.String("file", cmd.DefaultFile, "filename with integer to check for primes")
	port := flag.String("serve", "", "Start HTTP Server to listen on /primes, integer inputs will output primes")
	flag.Parse()

	// If serve is specified, start http server
	if serve.IsServeSet() {
		serve.Serve(*port)
		return
	}
	// initialize integer to use
	pInt := *i
	if !cmd.IsIntSet() {
		pInt = cmd.ReadIntFile(*f)
	}
	cmd.PrimesCmd(pInt)
}
