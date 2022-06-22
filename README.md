# Primes

find primes for a given integer

## build 

local binary `go build`

as a container `docker compose build`


## run

run directly `go run . --int 42`

run as a container `docker compose run --int 42` 


## Requirements

A prime number is an integer greater than 1, divisible only by one and itself. 

A prime number sieve is an algorithm or program that, when given an integer, returns a list of 
prime numbers greater than 1 and less than or equal to that integer.

For example, a sieve, given the number 10, would return: 2, 3, 5, and 7.

Given the number 21, it would return: 2, 3, 5, 7, 11, 13, 17, and 19.

For this exercise, write a prime number sieve in any language you want. Performance of the sieve does
