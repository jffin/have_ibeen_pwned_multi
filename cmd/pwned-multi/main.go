package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"

	"github.com/jffin/have_ibeen_pwned_multi/pkg/client"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/constants"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/files"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/structs"
)

func main() {

	inputFile := flag.String("input", constants.INPUT_FILE_NAME, "input file with targets")
	outputFile := flag.String("output", constants.RESULT_FILE_NAME, "output file write results to")
	apiKey := flag.String("key", "api-key", "hibp-api-key")
	flag.Parse()

	var targetsArray []string = files.ReadInputFile(*inputFile)

	rateLimiter := rate.NewLimiter(rate.Every(1501*time.Millisecond), 1) // 1 request every 1500 milliseconds
	client := client.NewClient(rateLimiter)

	channel := make(chan structs.Response)
	for _, target := range targetsArray {
		go checkEmail(target, *apiKey, client, channel)
	}

	results := make([]structs.Response, len(targetsArray))
	for index, _ := range results {
		results[index] = <-channel
	}
	files.WriteOutputFile(*outputFile, results)
}

func checkEmail(target string, apiKey string, client *client.RLHTTPClient, channel chan structs.Response) {
	endpoint := fmt.Sprintf("%s/%s?%s", constants.REQUEST_URL, url.QueryEscape(target), "truncateResponse=false")
	request, err := http.NewRequest(constants.DEFAULT_REQUEST_METHOD, endpoint, nil)
	errors.Check("new request constraining", err)

	request.Header.Set("User-agent", constants.DEFAULT_USER_AGENT)
	request.Header.Set("hibp-api-key", apiKey)

	response, err := client.Do(request)
	errors.Check("request issuing", err)

	defer response.Body.Close()
	channel <- readResponse(target, response)
}

func readResponse(target string, response *http.Response) structs.Response {
	var b bytes.Buffer
	if _, err := io.Copy(&b, response.Body); err != nil {
		errors.Check("reading response body", err)
	}

	var responseData []structs.ResponseData
	json.Unmarshal([]byte(b.String()), &responseData)

	return structs.Response{Target: target, Data: responseData}
}
