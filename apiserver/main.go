package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func newApiCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	var (
		listenAddr = ":8080"
		dburl      = "localhost:27017"
	)

	cmds := &cobra.Command{
		Use:   "cxserver",
		Short: "cxserver: cloud10x server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer(dburl, listenAddr)
		},
	}

	cmds.Flags().StringVar(&dburl, "dburl", dburl, "Database's URL")
	cmds.Flags().StringVar(&listenAddr, "listen", listenAddr, "Listen address")
	return cmds
}

func main() {
	cmd := newApiCommand(os.Stdin, os.Stdout, os.Stderr)
	cmd.Execute()

}

func startServer(dburl string, listenAddr string) {
	var err error

	fmt.Printf("Create database driver\n")

	gDB, err = newDBDriver("mongo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("Initialize database\n")

	err = gDB.Init(dburl, "c10x")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("Create API routes\n")
	app := newApiApp()

	fmt.Printf("Start server\n")
	app.Run(iris.Addr(listenAddr))

}
