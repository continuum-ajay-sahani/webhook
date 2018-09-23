package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/webhook/util"
)

const destURL = "http://localhost:9080"

func main() {
	fmt.Println("I am in server2")
	listner()
}

func listner() {
	r := gin.Default()
	r.POST("/handler/record/:id", recordHandler)
	r.Run(":9090")
}

func recordHandler(c *gin.Context) {
	rid := c.Param("id")
	go performAction(rid)
}

func performAction(rid string) {
	randNo := time.Duration(rand.Intn(10))
	res := util.Response{
		RID:          rid,
		RandDuration: randNo,
	}
	fmt.Println("Sleep duration:=", randNo)
	time.Sleep(randNo * time.Second)
	url := fmt.Sprintf("%s/callback/record/%v", destURL, rid)
	util.Call("PUT", url, &res)
}
