package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	admapi "github.com/jmpq/cloud10x/adm/api"
	"github.com/jmpq/cloud10x/adm/certs"
	"github.com/jmpq/cloud10x/v1"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
)

func NewCmdCluster(in io.Reader, out, err io.Writer) *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "Cluster commands",
	}

	clusterCmd.Flags().StringVar(&gRemoteUrl, "url", gRemoteUrl, "Management server's url")
	clusterCmd.Flags().StringVar(&gTenantName, "tenant", gTenantName, "The tenant's name")
	clusterCmd.Flags().StringVar(&gTenantSecret, "secret", gTenantSecret, "Tenant's secret")

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify the cluster's name\n")
				os.Exit(-1)
			}
			clusterCreate(cmd, args)
		},
	}

	clusterCmd.AddCommand(createCmd)

	var (
		cluster string
	)

	joinCmd := &cobra.Command{
		Use:   "join",
		Short: "join the host to a cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify cluster's name")
				os.Exit(-1)
			}
			clusterJoin(cmd, args, cluster)
		},
	}

	joinCmd.Flags().StringVar(&cluster, "cluster", cluster, "Cluster name to join")

	clusterCmd.AddCommand(joinCmd)

	return clusterCmd
}

func clusterCreate(cmd *cobra.Command, args []string) {
	url := gRemoteUrl
	tenant := gTenantName
	secret := gTenantSecret

	clusterName := args[0]
	fmt.Printf("Creating cluster %v\n", clusterName)
	fmt.Printf("Step 1. create certificates\n")

	// create cert files
	cfg := &admapi.MasterConfiguration{CertificatesDir: "~/.cloud10x/certs"}

	err := certs.CreatePKIAssets(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating cert files %v", err)
		os.Exit(-1)
	}

	req := v1.ClusterCreateReq{
		Name: clusterName,
	}

	data, _ := json.Marshal(req)

	resp, err := doPostRequest(url+"/v1/Clusters", string(data), tenant, secret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "response %v", resp)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	r := &v1.ClusterCreateResp{}
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), r)
	fmt.Printf("Cluster created, token: %s", r.Token)
}

func clusterJoin(cmd *cobra.Command, args []string, clusterName string) {
	ipAddress := args[1]
	fmt.Printf("Joined cluster %s for %s.\n", clusterName, ipAddress)
}