package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func NewCxiCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "cxi",
		Short: "cxi: cloud10x cli",
	}

	cmds.AddCommand(NewCmdCluster(in, out, err))
	return cmds
}

func main() {
	cmd := NewCxiCommand(os.Stdin, os.Stdout, os.Stderr)
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
}
