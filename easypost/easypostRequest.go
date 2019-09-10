package easypost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const easypost_url = "https://api.easypost.com/v2"

var api_key string
var Request RequestControllerInterface

var debug = false

func SetDebug(enabled bool) {
	debug = enabled
}

type RequestControllerInterface interface {
	do(method string, objectType string, objectUrl string, payload interface{}) ([]byte, error)
}

func init() {
	Request = RequestControllerFake{}
}

func SetApiKey(key string) {
	api_key = key
}

type RequestController struct{}

func nest(k string, o interface{}) map[string]interface{} {
	return map[string]interface{}{
		k: o,
	}
}

func DebugLog(format string, args ...interface{}) {
	if debug {
		fmt.Printf("EasyPost: "+format+"\n", args...)
	}
}

//Request request an EasyPost API
func (rc RequestController) do(method string, objectType string, objectUrl string, payload interface{}) ([]byte, error) {
	url := getRequestURL(objectType, objectUrl)
	body, err := json.Marshal(payload)
	if err != nil {
		return body, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("cannot create EasyPost request")
	}

	DebugLog("easypost sending %v request to: %v\n", method, url)
	DebugLog("easypost api key: %v", api_key)
	DebugLog("easypost payload %v", string(body))

	req.SetBasicAuth(api_key, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//getRequestUrl returns the correct url for EasyPost API
func getRequestURL(objectType string, objectUrl string) string {
	url := fmt.Sprintf("%vs", objectType)
	if objectType == "address" {
		url = "addresses"
	}
	if objectType == "batch" {
		url = "batches"
	}

	if objectUrl != "" {
		url = fmt.Sprintf("%v/%v", url, objectUrl)
	}

	return fmt.Sprintf("%v/%v", easypost_url, url)
}
