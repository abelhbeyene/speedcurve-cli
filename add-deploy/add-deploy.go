package main

import (
	"fmt"

	"github.com/speedcurve/common"
)

func main() {
	var headers map[string]string = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	payload := "site_id=730439&note=Release-XX&detail=Details%20go%20here"
	response, _ := common.MakeRequest("/deploys", "POST", headers, payload)

	fmt.Println(response)
}
