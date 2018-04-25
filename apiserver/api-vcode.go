package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/jmpq/cloud10x/apiserver/db"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"os"
	"time"
)

func apiVCodeRequest(ctx iris.Context) {
	phoneNum := ctx.Params().Get("phone")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//db := gDB.(*db.Mongo)
	//err := db.Db.C("vcode").Insert(&VCode{VCode: vcode, PhoneNum: phoneNum})

	err := gDB.Upsert("vcode",
		bson.M{"phonenum": phoneNum},
		&VCode{VCode: vcode, PhoneNum: phoneNum, Time: time.Now()})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to request vcode, error: %v\n", err)
		return
	}

	fmt.Printf("vcode %s was sent to phone %s.\n", vcode, phoneNum)
}
