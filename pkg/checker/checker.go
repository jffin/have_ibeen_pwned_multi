package checker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jffin/have_ibeen_pwned_multi/pkg/client"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/constants"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/structs"
)

func StartCheck(targetsArray []string, apiKey string) []structs.Response {
	client := client.CreateNewClient()

	channel := make(chan structs.Response)
	for _, target := range targetsArray {
		go checkEmail(target, apiKey, client, channel)
	}

	return getResults(channel, len(targetsArray))
}

func getResults(channel chan structs.Response, resultsSize int) []structs.Response {
	results := make([]structs.Response, resultsSize)
	for index, _ := range results {
		results[index] = <-channel
	}
	return results
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
