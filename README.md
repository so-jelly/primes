# Primes

find primes for a given integer

## build 

local binary `go build`
as a container `docker compose build`

## run

### argument input

run with go `go run . --int 42`
run as a container `docker compose run primes --int 42` 

### file input

if no int argument is provided, primes will try to read the file `input.txt` different files can be provided with the `--file` argument

create a test file `echo 42 > ./input.txt` 
run with go `go run . --file ./input.txt`

we need to mount the file via docker compose, this is done with an environment variable
`export PRIMES_FILE=$PWD/input.txt`
then
`docker compose run primes --file $PRIMES_FILE`

# test

test all the subdirectories, tests have been written for the meaty part of this, more coverage could certainly be added

`go test -v ./...`
