GOFLAGS+=-tags=clientQueryAskAPI

lotus-query-ask-api-daemon:
	go build .

clean:
	rm -f rm -f lotus-query-ask-api-daemon

why-ffi:
	go mod why github.com/filecoin-project/filecoin-ffi

list:
	go list -e -json -compiled=true -test=true -deps=true .

checkout:
	mkdir -p extern
	cd extern; git clone -b jim-query-ask-api-daemon-modified git@github.com:jimpick/lotus.git lotus-modified
	cd extern; git clone -b jim-query-ask-api-daemon-modified git@github.com:jimpick/go-fil-markets.git go-fil-markets-modified
	cd extern; git clone -b fil-blst-v0.1.1 https://github.com/filecoin-project/fil-blst.git
