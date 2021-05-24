GOPKG ?=	moul.io/u

include rules.mk

generate:
	GO111MODULE=off go get github.com/campoy/embedmd
	GO111MODULE=off go get github.com/cespare/prettybench
	mkdir -p .tmp
	go doc -all > .tmp/godoc.txt
	go test -v -run=NONE -bench=. | prettybench -no-passthrough > .tmp/bench.txt
	embedmd -w README.md
	rm -rf .tmp

	$(GO) run github.com/tailscale/depaware --update .
