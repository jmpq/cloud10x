package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmpq/cloud10x/v1"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
)

func NewCmdTenant(in io.Reader, out, err io.Writer) *cobra.Command {
	tenantCmd := &cobra.Command{
		Use:   "tenant",
		Short: "Tenant commands",
	}

	tenantCmd.Flags().StringVar(&gRemoteUrl, "url", gRemoteUrl, "Management server's url")
	tenantCmd.Flags().StringVar(&gTenantName, "tenant", gTenantName, "The tenant's name")
	tenantCmd.Flags().StringVar(&gTenantSecret, "secret", gTenantSecret, "Tenant's secret")

	var (
		org      = ""
		phoneNum = ""
		email    = ""
		vcode    = ""
		password string
	)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a tenant",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify the tenant's name\n")
				os.Exit(-1)
			}
			if phoneNum == "" {
				fmt.Fprintf(os.Stderr, "Mobile phone number must not be empty\n")
				os.Exit(-1)
			}
			if vcode == "" {
				fmt.Fprintf(os.Stderr, "Verification code must not be empty\n")
				os.Exit(-1)
			}
			if password == "" {
				fmt.Fprintf(os.Stderr, "Password must not be empty\n")
				os.Exit(-1)
			}
			tenantCreate(cmd, args, org, phoneNum, email, password, vcode)
		},
	}

	createCmd.Flags().StringVar(&org, "org", org, "Organization of the tenant")
	createCmd.Flags().StringVar(&phoneNum, "phone", phoneNum, "Phone number of the tenant")
	createCmd.Flags().StringVar(&email, "email", email, "Email of the tenant")
	createCmd.Flags().StringVar(&password, "password", vcode, "Password")
	createCmd.Flags().StringVar(&vcode, "vcode", vcode, "Verification code")

	tenantCmd.AddCommand(createCmd)

	requestVCodeCmd := &cobra.Command{
		Use:   "vcode",
		Short: "Request verification code",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify your mobile phone number\n")
				os.Exit(-1)
			}
			tenantRequestVCode(cmd, args)
		},
	}

	tenantCmd.AddCommand(requestVCodeCmd)
	return tenantCmd
}

func tenantRequestVCode(cmd *cobra.Command, args []string) {
	url := gRemoteUrl
	phoneNum := args[0]

	resp, err := doGetRequest(url+"/v1/VCode/"+phoneNum, "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "response %v", resp)
		os.Exit(-1)
	}

	fmt.Printf("Verification code was sent to %s.\n", phoneNum)
}

func tenantCreate(cmd *cobra.Command, args []string, org, phoneNum, email, password, vcode string) {
	url := gRemoteUrl
	tenantName := args[0]
	fmt.Printf("Creating tenant %v\n", tenantName)

	req := v1.TenantCreateReq{
		Name:     tenantName,
		PhoneNum: phoneNum,
		Email:    email,
		Password: password,
		VCode:    vcode,
	}

	data, _ := json.Marshal(req)

	resp, err := doPostRequest(url+"/v1/Tenants", string(data), "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "response %v", resp)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	r := &v1.TenantCreateResp{}
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), r)
	fmt.Printf("Tenant %s was created, secret: %s, please save your secret in a safe place.\n", tenantName, r.Secret)
}
