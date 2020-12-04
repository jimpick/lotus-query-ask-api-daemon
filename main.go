package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-jsonrpc"
	lcli "github.com/filecoin-project/lotus/cli/cmd"
	"github.com/filecoin-project/lotus/node/modules/moduleapi"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/jimpick/lotus-query-ask-api-daemon/api"
	"github.com/jimpick/lotus-query-ask-api-daemon/node"
	"github.com/jimpick/lotus-utils/fxnodesetup"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-daemon/p2pclient"
)

const flagQueryAskRepo = "query-ask-repo"

const listenAddr = "127.0.0.1:9301"

var daemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "run client query ask api daemon",
	Action: func(cctx *cli.Context) error {
		var queryAskAPI api.QueryAskAPI

		// remote libp2p node for non-wss
		// controlMaddr, _ := multiaddr.NewMultiaddr("/dns4/libp2p-caddy-p2pd.localhost/tcp/9059/wss")
		controlMaddr, _ := multiaddr.NewMultiaddr("/dns4/p2pd.v6z.me/tcp/9059/wss")
		listenMaddr, _ := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/0")
		p2pclientNode, err := p2pclient.NewClient(controlMaddr, listenMaddr)
		fmt.Printf("Jim p2pclientNode %v\n", p2pclientNode)
		nodeID, nodeAddrs, err := p2pclientNode.Identify()
		peerInfo := peer.AddrInfo{
			ID:    nodeID,
			Addrs: nodeAddrs,
		}
		fmt.Printf("Jim peerInfo %v\n", peerInfo)
		addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
		if err != nil {
			panic(err)
		}
		fmt.Println("p2pclient->p2pd node address:", addrs[0])

		nodeAPI, ncloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer ncloser()
		ctx := lcli.DaemonContext(cctx)

		r, err := repo.NewFS(cctx.String(flagQueryAskRepo))
		if err != nil {
			return xerrors.Errorf("opening fs repo: %w", err)
		}

		// Re-use repo.Worker type as it has no config defaults
		if err := r.Init(repo.Worker); err != nil && err != repo.ErrRepoExists {
			return xerrors.Errorf("repo init error: %w", err)
		}

		_, err = node.New(ctx,
			node.QueryAskAPI(&queryAskAPI),
			node.Repo(r),
			node.Online(),
			fxnodesetup.Override(new(*p2pclient.Client), p2pclientNode),
			fxnodesetup.Override(new(moduleapi.ChainModuleAPI), nodeAPI),
			fxnodesetup.Override(new(moduleapi.StateModuleAPI), nodeAPI),
		)
		if err != nil {
			return xerrors.Errorf("initializing node: %w", err)
		}
		rpcServer := jsonrpc.NewServer()
		rpcServer.Register("Filecoin", queryAskAPI)

		http.Handle("/rpc/v0", rpcServer)

		fmt.Printf("Listening on http://%s\n", listenAddr)
		return http.ListenAndServe(listenAddr, nil)
	},
}

func main() {
	app := &cli.App{
		Name: "lotus-client-query-ask-api-daemon",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    flagQueryAskRepo,
				EnvVars: []string{"LOTUS_QUERY_ASK_PATH"},
				Value:   "~/.lotus-client-query-ask", // TODO: Consider XDG_DATA_HOME
			},
		},
		Commands: []*cli.Command{
			daemonCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
