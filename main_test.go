package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/urfave/cli/v2"
)

func TestMain(t *testing.T) {
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
	app.Setup()
	go func() {
		args := []string{os.Args[0], "daemon"}
		if err := app.Run(args); err != nil {
			panic(err)
		}
	}()

	const timeout = 15
	fmt.Printf("Running for %v seconds.\n", timeout)
	time.Sleep(timeout * time.Second)
	fmt.Println("Shutting down.")
}
