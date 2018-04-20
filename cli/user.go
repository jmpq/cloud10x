package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmpq/cloud10x/v1"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
)

func NewCmdUser(in io.Reader, out, err io.Writer) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "User commands",
	}

	userCmd.Flags().StringVar(&gRemoteUrl, "url", gRemoteUrl, "Management server's url")
	userCmd.Flags().StringVar(&gTenantName, "tenant", gTenantName, "The tenant's name")
	userCmd.Flags().StringVar(&gTenantSecret, "secret", gTenantSecret, "Tenant's secret")

	var (
		org        = ""
		department = ""
		phoneNum   = ""
		email      = ""
		password   string
	)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Please specify the user name\n")
				os.Exit(-1)
			}
			userCreate(cmd, args, org, department, phoneNum, email, password)
		},
	}

	createCmd.Flags().StringVar(&org, "org", org, "Organization of the user")
	createCmd.Flags().StringVar(&department, "department", department, "Department of the user")
	createCmd.Flags().StringVar(&phoneNum, "phone", phoneNum, "Phone number of the user")
	createCmd.Flags().StringVar(&email, "email", email, "Email of the user")
	createCmd.Flags().StringVar(&password, "password", password, "Password of the user")

	userCmd.AddCommand(createCmd)

	return userCmd
}

func userCreate(cmd *cobra.Command, args []string, org, department, phoneNum, email, password string) {
	url := gRemoteUrl
	tenant := gTenantName
	secret := gTenantSecret

	userName := args[0]
	fmt.Printf("Creating user %v\n", userName)

	req := v1.UserCreateReq{
		Tenant:     tenant,
		Name:       userName,
		Department: department,
		PhoneNum:   phoneNum,
		Email:      email,
		Password:   password,
	}

	data, _ := json.Marshal(req)

	resp, err := doPostRequest(url+"/v1/Users", string(data), tenant, secret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "response %v", resp)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	r := &v1.UserCreateResp{}
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), r)
	fmt.Printf("User %s was created, secret: %s, please save your secret in a safe place.", userName, r.Secret)
}
