fizzbuzzd
=======

FizzBuzz server for generating FizzBuzz terms written in GO. 

## Architecture

fizzbuzzd is a REST server with two endpoints.

## Requirements 

Install :

* Docker version 17.12.0+
* Go 1.9.2+

## Installation 

### Binaries

`go get -u github.com/fkarakas/fizzbuzzd/cmd/fizzbuzzd`

### Docker container

* Signed version 

```
docker run -d -p 8080:8080 fkarakas/fizzbuzzd:signed
```

* Unsigned version (automatically build from github push) 

```
docker run -d -p 8080:8080 fkarakas/fizzbuzzd:latest
```

## Build

* Set $GOPATH variable to your GO workspace : `export GOPATH=$HOME/go`

* Checkout the project : 
```
git clone git@github.com:fkarakas/fizzbuzzd.git $GOPATH/src/github.com/fkarakas/fizzbuzzd
```

* Install dep : `go get -u github.com/golang/dep/cmd/dep`

* Install gomegalinter (for lint checks) : `go get -u gopkg.in/alecthomas/gometalinter.v2`

* Install linters : `$GOPATH/bin/gometalinter.v2 --install`

* run `make` in `$GOPATH/src/github.com/fkarakas/fizzbuzzd`

## Build the docker image

run :
```
make image
```


## fizzbuzzd Configuration

```
Usage of fizzbuzzd:
    -port : server binding port (default 8080)
```

example :

```
fizzbuzzd -port=8888     
```

## Endpoints Documentation

fizzbuzzd responds directly to the following endpoints.

### Health check :

* `/` - returns a 200 OK response with the version number (used for health checks)

### Get an array of FizzBuzz terms :

* `GET /api/v1/fizzbuzz/numbers/:number1/:number2/terms/:term1/:term2`  - return an 200 OK response : returns an array of strings

### Parameters

* number1 : first match number
* number2 : second match number
* term1 : first replacement string
* term2 : second replacement string

* limit (query parameter ..?limit=200) : optional limit parameter for the number of returned terms (default 100)

## curl example

```
curl -v -XGET http://localhost:8080/api/v1/fizzbuzz/numbers/2/3/terms/bob/dylan\?limit=123
```

Response : HTTP CODE 200

`{"result" : [...]}`

## Tests

Several kind of tests are available :

* Unit tests

* Blackbox tests


### Unit tests

Unit tests are used to verify the wanted behaviour of the server.

To run the unit tests : `make test`

### Blackbox tests

Blackbox tests are used to verify the docker image built.
The Blackbox tests simulate external call from a program using fizzbuzzd. 

To run the tests : `make blackbox-test VERSION=1.0.0`

where `VERSION` is the container version of fizzbuzzd