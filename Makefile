NAME = fizzbuzzd

GOPATH ?= ${HOME}/go

LDFLAGS=-ldflags "-X main.version=${VERSION}"

default: build

build: clean fmt dep lint test
	go build -i -v ${LDFLAGS} -o ${NAME} ./cmd/fizzbuzzd

dep:
	if [ -f "Gopkg.toml" ] ; then dep ensure ; else dep init ; fi

clean:
	if [ -f "${NAME}" ] ; then rm ${NAME} ; fi

lint:
	${GOPATH}/bin/gometalinter.v2 go --vendor --tests --errors --concurrency=2 --deadline=60s ./...

fmt:
	go fmt ./...

test:
	go test -v ./...

blackbox-test: image
	VERSION=${VERSION} docker-compose -f ./test/docker-compose.yml up -d --force-recreate
	docker run --rm --net=fizzbuzzd-network -v ${GOPATH}:/go -i golang go test -v github.com/fkarakas/fizzbuzzd/test/... -tags=blackbox
	VERSION=${VERSION} docker-compose -f ./test/docker-compose.yml down

image: check-version
	docker build --rm . -t fkarakas/${NAME}:${VERSION} --build-arg VERSION=${VERSION}

check-version:
ifndef VERSION
	$(error VERSION is undefined)
endif

.PHONY: test image