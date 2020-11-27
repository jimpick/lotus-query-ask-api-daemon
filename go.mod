module github.com/jimpick/lotus-query-ask-api-daemon

go 1.15

require (
	github.com/filecoin-project/go-address v0.0.4
	github.com/filecoin-project/go-fil-markets v1.0.0
	github.com/filecoin-project/go-jsonrpc v0.1.2
	github.com/filecoin-project/lotus v0.0.0-00010101000000-000000000000
	github.com/jimpick/lotus-utils v0.0.2
	github.com/libp2p/go-libp2p v0.11.0
	github.com/libp2p/go-libp2p-core v0.6.1
	github.com/libp2p/go-libp2p-daemon v0.2.2
	github.com/libp2p/go-libp2p-kad-dht v0.8.3
	github.com/libp2p/go-libp2p-peerstore v0.2.6
	github.com/libp2p/go-libp2p-record v0.1.3
	github.com/libp2p/go-ws-transport v0.3.1
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/fx v1.13.1
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/lotus => ./extern/lotus-modified

replace github.com/filecoin-project/go-fil-markets => ./extern/go-fil-markets-modified

replace github.com/supranational/blst => ./extern/fil-blst/blst

replace github.com/filecoin-project/fil-blst => ./extern/fil-blst

replace github.com/libp2p/go-libp2p => github.com/jimpick/go-libp2p v0.3.2-0.20201019200821-ea23aad1363a

replace github.com/libp2p/go-ws-transport => github.com/jimpick/go-ws-transport v0.1.1-0.20201019202657-b30eedb00d94

replace github.com/libp2p/go-reuseport-transport => github.com/jimpick/go-reuseport-transport v0.0.5-0.20201019202422-85fd62f8a44c

replace github.com/filecoin-project/go-jsonrpc => ../../../../browser-markets/go-jsonrpc-wasm

