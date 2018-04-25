package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"github.com/jmpq/cloud10x/adm/certs"
	"github.com/jmpq/cloud10x/v1"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	kubeadmapiext "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"os"
)

var (
	cfg *kubeadmapiext.MasterConfiguration
)

func NewCmdCluster(in io.Reader, out, err io.Writer) *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "Cluster commands",
	}

	clusterCmd.Flags().StringVar(&gRemoteUrl, "url", gRemoteUrl, "Management server's url")
	clusterCmd.Flags().StringVar(&gTenantName, "tenant", gTenantName, "The tenant's name")
	clusterCmd.Flags().StringVar(&gTenantSecret, "secret", gTenantSecret, "Tenant's secret")

	// Cluster create

	//FIXME, why can't find this line in kubeadm but it works ?
	kubeadmapiext.RegisterDefaults(legacyscheme.Scheme)
	cfg = &kubeadmapiext.MasterConfiguration{}

	// Default values for the cobra help text
	legacyscheme.Scheme.Default(cfg)

	var (
		cfgPath string
	)
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify the cluster's name\n")
				os.Exit(-1)
			}
			clusterCreate(cmd, args, &cfgPath, cfg)
		},
	}

	// Add flags to the command
	createCmd.Flags().StringVar(&cfgPath, "config", cfgPath, "Path to config file (WARNING: Usage of a configuration file is experimental)")
	createCmd.Flags().StringVar(&cfg.CertificatesDir, "cert-dir", cfg.CertificatesDir, "The path where to save the certificates")
	createCmd.Flags().StringVar(&cfg.Networking.DNSDomain, "service-dns-domain", cfg.Networking.DNSDomain, "Alternative domain for services, to use for the API server serving cert")
	createCmd.Flags().StringVar(&cfg.Networking.ServiceSubnet, "service-cidr", cfg.Networking.ServiceSubnet, "Alternative range of IP address for service VIPs, from which derives the internal API server VIP that will be added to the API Server serving cert")
	createCmd.Flags().StringSliceVar(&cfg.APIServerCertSANs, "apiserver-cert-extra-sans", []string{}, "Optional extra altnames to use for the API server serving cert. Can be both IP addresses and dns names")
	createCmd.Flags().StringVar(&cfg.API.AdvertiseAddress, "apiserver-advertise-address", cfg.API.AdvertiseAddress, "The IP address the API server is accessible on, to use for the API server serving cert")

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

func clusterCreate(cmd *cobra.Command, args []string, cfgPath *string, cfg *kubeadmapiext.MasterConfiguration) {
	url := gRemoteUrl
	tenant := gTenantName
	secret := gTenantSecret

	clusterName := args[0]
	fmt.Printf("Creating cluster %v\n", clusterName)
	fmt.Printf("Step 1. create certificates\n")

	// create cert files
	err := createPKIAssets(cmd, args, cfgPath, cfg)
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
