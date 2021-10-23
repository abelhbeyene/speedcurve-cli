package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	godotenv "github.com/joho/godotenv"
)

const API_HOST string = "API_HOST"
const API_VERSION string = "API_VERSION"
const API_KEY string = "API_KEY"
const ERR_NO_API_KEY string = "Please provide API_KEY env var when running. E.g. API_KEY=[YOUR KEY HERE] go run ."

var ConfigMap = make(map[string]string)

func init() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ConfigMap[API_HOST] = os.Getenv(API_HOST)
	ConfigMap[API_KEY] = os.Getenv(API_KEY)
	ConfigMap[API_VERSION] = os.Getenv(API_VERSION)
}

func getEncodedKey() (string, error) {
	var apiKey string = os.Getenv(API_KEY)
	if apiKey == "" {
		return "", errors.New(ERR_NO_API_KEY)
	}
	encodedKey := base64.StdEncoding.EncodeToString([]byte(apiKey))
	return encodedKey, nil
}

func MakeRequest(path string, method string, headers map[string]string, rawPayload string) (string, error) {
	fmt.Println("rawPayload", rawPayload)
	encodedKey, apiKeyErr := getEncodedKey()
	if apiKeyErr != nil {
		fmt.Println(apiKeyErr)
		return "", nil
	}
	url := ConfigMap[API_HOST] + ConfigMap[API_VERSION] + path

	payload := strings.NewReader(rawPayload)
	fmt.Println("payload", payload)

	req, reqErr := http.NewRequest(method, url, payload)

	if reqErr != nil {
		fmt.Println("reqErr", reqErr)
		return "", reqErr
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+encodedKey)

	for key, h := range headers {
		req.Header.Add(key, h)
	}

	res, resErr := http.DefaultClient.Do(req)

	if resErr != nil {
		fmt.Println("resErr", resErr)
		return "", resErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	return bodyStr, nil
}

func main() {

	fmt.Println("Usage = common.GetEncodedKey")
}
