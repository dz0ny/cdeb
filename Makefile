VERSION := 1.0.0

all: setup build lint

setup:
	go get github.com/aktau/github-release
	go get github.com/alecthomas/gometalinter
	go get -v -d cdeb
	bin/gometalinter --install --update

clean:
	rm -f cdeb
	rm -rf pkg
	rm -rf bin
	find src/* -maxdepth 0 ! -name 'cdeb' -type d | xargs rm -rf

lint:
	bin/gometalinter --fast --disable=gotype --disable=gas --cyclo-over=15 src/cdeb/...
	find src/cdeb -name '*.go' | xargs gofmt -w -s

build:
	env GOOS=linux GOARCH=arm go build --ldflags '-w -X main.version=$(VERSION)' -o cdeb-Linux-armv7l cdeb
	env GOOS=linux GOARCH=amd64 go build --ldflags '-w -X main.version=$(VERSION)' -o cdeb-Linux-x86_64 cdeb
	env GOOS=darwin GOARCH=amd64 go build --ldflags '-w -X main.version=$(VERSION)' -o cdeb-Darwin-x86_64 cdeb
	env GOOS=windows GOARCH=amd64 go build --ldflags '-w -X main.version=$(VERSION)' -o cdeb-Windows-x86_64.exe cdeb

install:
	sudo mv cdeb-`uname -s`-`uname -m` /usr/local/bin/cdeb

upload:
	bin/github-release upload \
			--user dz0ny \
			--repo cdeb \
			--tag "v$(VERSION)" \
			--name "cdeb-Linux-armv6l" \
			--file cdeb-Linux-armv7l
	bin/github-release upload \
	    --user dz0ny \
	    --repo cdeb \
	    --tag "v$(VERSION)" \
	    --name "cdeb-Linux-armv7l" \
	    --file cdeb-Linux-armv7l
	bin/github-release upload \
	    --user dz0ny \
	    --repo cdeb \
	    --tag "v$(VERSION)" \
	    --name "cdeb-Linux-x86_64" \
	    --file cdeb-Linux-x86_64
	bin/github-release upload \
	    --user dz0ny \
	    --repo cdeb \
	    --tag "v$(VERSION)" \
	    --name "cdeb-Darwin-x86_64" \
	    --file cdeb-Darwin-x86_64
	bin/github-release upload \
	    --user dz0ny \
	    --repo cdeb \
	    --tag "v$(VERSION)" \
	    --name "cdeb-Windows-x86_64.exe" \
	    --file cdeb-Windows-x86_64.exe