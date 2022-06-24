package serve

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/so-jelly/primes/pkg/primes"
)

func primesHandler(w http.ResponseWriter, r *http.Request) {
	respBody := make([]byte, 0)
	if r.Body == nil {
		w.Write(respBody)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respBody)
		return
	}
	defer r.Body.Close()

	primeInt, err := strconv.Atoi(string(reqBody))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respBody)
		return
	}

	p := &primes.Primes{PrimeInt: primeInt}
	p.GetPrimes()
	respBody, err = json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respBody)
		return
	}
	w.Write(respBody)
}

// Serve primes handler on specified port
func Serve(portString string) {
	http.HandleFunc("/primes", primesHandler)
	err := http.ListenAndServe(":"+portString, nil)
	fmt.Println(err)
}

// IsServeSet checks if the serve flag is set
func IsServeSet() bool {
	var set bool
	flag.Visit(
		func(serveFlag *flag.Flag) {
			if serveFlag.Name == "serve" {
				set = true
			}
		})
	return set
}
