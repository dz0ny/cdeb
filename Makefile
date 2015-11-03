all: prepare build upload

prepare:
	go get -d
	go get github.com/aktau/github-release

build:
	env GOOS=linux GOARCH=amd64 go build -o cdeb-linux-amd64

upload:
	github-release upload \
	    --user dz0ny \
	    --repo cdeb \
	    --tag v0.1.0 \
	    --name "cdeb-linux-amd64" \
	    --file cdeb-linux-amd64
