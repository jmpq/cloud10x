package main

import (
	//"encoding/json"
	"fmt"
	"github.com/jmpq/cloud10x/v1"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func apiTenantList(ctx iris.Context) {
}

func apiTenantGet(ctx iris.Context) {
}

func apiTenantPost(ctx iris.Context) {
	var req v1.TenantCreateReq
	resp := v1.TenantCreateResp{}

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		resp.Message = err.Error()
		ctx.JSON(&resp)
		return
	}
	fmt.Printf("Create tenant %v\n", req)

	var vCode VCode
	err = gDB.One("vcode", bson.M{"phonenum": req.PhoneNum}, &vCode)
	if err != nil {
		//if err == errors.ErrNotFound {
		//}
		fmt.Fprintf(os.Stderr, "Failed to get vcode for phone num %s, Error: %v\n", req.PhoneNum, err)
		ctx.StatusCode(iris.StatusNotFound)
		resp.Message = err.Error()
		ctx.JSON(&resp)
		return
	}

	fmt.Printf("vcode %s vs %s\n", vCode.VCode, req.VCode)
	if vCode.VCode != req.VCode {
		ctx.StatusCode(iris.StatusNotFound)
		fmt.Fprintf(os.Stderr, "Verification code does not match for tenant %s\n", req.Name)
		resp.Message = "Verification code does not match"
		ctx.JSON(&resp)
		return
	}

	secret := getToken()
	err = gDB.Insert("tenants",
		&Tenant{Name: req.Name, Org: req.Org, PhoneNum: req.PhoneNum, Email: req.Email,
			Password: req.Password, Secret: secret, IsPremiumAccount: false})

	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		resp.Message = err.Error()
		ctx.JSON(&resp)
		return
	}

	resp.Secret = secret
	ctx.JSON(&resp)
}

func apiTenantDelete(ctx iris.Context) {
}
