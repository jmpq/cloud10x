package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func NewAdmCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "cxadm",
		Short: "cxadm: cloud10x cli",
	}

	cmds.AddCommand(NewCmdCluster(in, out, err))
	cmds.AddCommand(NewCmdTenant(in, out, err))
	cmds.AddCommand(NewCmdUser(in, out, err))
	return cmds
}
func main() {
	cmd := NewAdmCommand(os.Stdin, os.Stdout, os.Stderr)
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
}
