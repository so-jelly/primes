package primes

type Primes struct {
	// Number is the integer to check
	Number *int `json:"Number,omitempty"`
	// Primes is the list of primes for primeInt
	Primes []int `json:"primes"`
}

// GetPrimes adds primes to a Primes
func (p *Primes) GetPrimes() {
	primes := make([]int, 0)
	// 3 is the first prime number
	for j := 3; j <= *p.Number; j++ {
		if IsPrime(j) {
			primes = append(primes, j)
		}
	}

	p.Primes = primes
	// set empty value so response does not include
	p.Number = nil
}

// IsPrime checks if an integer is prime
func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
