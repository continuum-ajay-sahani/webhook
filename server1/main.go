package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/webhook/util"
)

const (
	destURL = "http://localhost:9090"
)

func main() {
	fmt.Println("I am server1")
	listener()
}

func listener() {
	r := gin.Default()
	r.PUT("/callback/record/:id", callbackHandler)
	r.POST("/action", performLoadTest)
	r.Run(":9080")
}

func callbackHandler(c *gin.Context) {
	bs, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("output=", string(bs))
}

func performLoadTest(c *gin.Context) {
	uid := uuid.New()
	url := fmt.Sprintf("%s/handler/record/%v", destURL, uid)
	util.Call("POST", url, nil)
}
