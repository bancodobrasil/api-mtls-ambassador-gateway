package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"

	"github.com/sirupsen/logrus"
)

func main() {
	logLevel := "info"
	url := ""

	flag.StringVar(&url, "url", "http://localhost:9090", "Protected service URL")

	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic("Invalid loglevel")
	}
	formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	formatter.File = true
	logrus.SetFormatter(&formatter)
	logrus.SetLevel(l)

	logrus.Infof("Starting mTLS plain HTTP requester Demo...")

	method := "POST"
	var httpClient = &http.Client{Timeout: 30 * time.Second}
	payloadBuf := new(bytes.Buffer)
	payload := map[string]string{
		"requestMessage": "I'm behind the Ambassador",
	}
	err = json.NewEncoder(payloadBuf).Encode(payload)

	if err != nil {
		logrus.Errorf("Error encoding the request body to JSON. Error: %s", err)
		logrus.Exit(999)
	}
	req, err := http.NewRequest(method, url, payloadBuf)
	req.Header.Add("dummy-api-key", "dummy_key_2021")
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		logrus.Errorf("Error creatig request objetct. Error: %s", err)
		logrus.Exit(999)
	}

	logrus.Infof("Requesting to %s", url)

	resp, err := httpClient.Do(req)
	if err != nil {
		logrus.Errorf("Error doing the request. Error: %s", err)
		logrus.Exit(999)
	}
	defer resp.Body.Close()

	respMap := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&respMap)

	if resp.StatusCode > 299 {
		logrus.Errorf("Unexpected http status on the request. Details: %+v", resp)
		logrus.Exit(999)
	}

	logrus.Infof("Invocation return: `%s`", respMap["responseMessage"])

}
