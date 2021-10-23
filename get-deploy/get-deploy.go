package main

import (
	"fmt"

	common "github.com/speedcurve/common"
)

func main() {
	var headers map[string]string = map[string]string{}
	response, err := common.MakeRequest("/deploys", "GET", headers, "")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
