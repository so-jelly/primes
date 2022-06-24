# Primes

find primes for a given integer

- [Primes](#primes)
  - [environment](#environment)
  - [build](#build)
  - [run](#run)
    - [argument input](#argument-input)
    - [file input](#file-input)
  - [test](#test)
  - [push](#push)
  - [k8s](#k8s)

## environment

using `.envrc` to provide some parameters for this project
`direnv` is a nice tool for populating the environment

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

## test

test all the subdirectories, tests have been written for the meaty part of this, more coverage could certainly be added

`go test -v ./...`

## push

push docker image to remote registry

`docker compose push`

## k8s

using [kind](https://kind.sigs.k8s.io/) as a local development environment

expose a nodePort with cluster creation 
<!-- TODO make some other ingress work / load balancer -->
`kind create cluster --config=deploy/kind.yml`

using `envsubst` to populate variables and apply the output

`envsubst <deploy/deployment.yml | kubectl apply -f -`

`kind.yml` has a static exposed port `8080`

test with `curl localhost:8080 --data 99999` or some such