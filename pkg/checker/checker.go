package checker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jffin/have_ibeen_pwned_multi/pkg/client"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
)

const (
	REQUEST_URL            string = "https://haveibeenpwned.com/api/v3/breachedaccount"
	DEFAULT_REQUEST_METHOD string = "GET"
	DEFAULT_USER_AGENT     string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0"
)

type ResponseData struct {
	Name         string
	Title        string
	Domain       string
	BreachDate   string
	AddedDate    string
	ModifiedDate string
	PwnCount     string
	Description  string
	LogoPath     string
	DataClasses  []string
	IsVerified   bool
	IsFabricated bool
	IsSensitive  bool
	IsRetired    bool
	IsSpamList   bool
}

type Response struct {
	Target string
	Data   []ResponseData
}

func StartCheck(targetsArray []string, apiKey string) []Response {
	client := client.CreateNewClient()

	channel := make(chan Response)
	for _, target := range targetsArray {
		go checkEmail(target, apiKey, client, channel)
	}

	return getResults(channel, len(targetsArray))
}

func getResults(channel chan Response, resultsSize int) []Response {
	results := make([]Response, resultsSize)
	for index, _ := range results {
		results[index] = <-channel
	}
	return results
}

func checkEmail(target, apiKey string, client *client.RLHTTPClient, channel chan Response) {
	endpoint := fmt.Sprintf("%s/%s?%s", REQUEST_URL, url.QueryEscape(target), "truncateResponse=false")
	request, err := http.NewRequest(DEFAULT_REQUEST_METHOD, endpoint, nil)
	errors.Check("new request constraining", err)

	request.Header.Set("User-agent", DEFAULT_USER_AGENT)
	request.Header.Set("hibp-api-key", apiKey)

	response, err := client.Do(request)
	errors.Check("request issuing", err)

	defer response.Body.Close()
	channel <- readResponse(target, response)
}

func readResponse(target string, response *http.Response) Response {
	var b bytes.Buffer
	if _, err := io.Copy(&b, response.Body); err != nil {
		errors.Check("reading response body", err)
	}

	var responseData []ResponseData
	json.Unmarshal([]byte(b.String()), &responseData)

	return Response{Target: target, Data: responseData}
}
