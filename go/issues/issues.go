package issues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.skia.org/infra/go/util"
)

const (
	API_QUERY_TEMPLATE = "%s.&fields=items/id,items/state,items/title&key=%s"
	API_URL_TEMPLATE   = "https://www.googleapis.com/projecthosting/v2/projects/skia/issues?q=%s"
)

// IssueTracker is a genric interface to an issue tracker that allows us
// to connect issues with items (identified by an id).
type IssueTracker interface {
	// FromQueury returns issue that match the given query string.
	FromQuery(q string) ([]Issue, error)
}

// CodesiteIssueTracker implements IssueTracker.
type CodesiteIssueTracker struct {
	apiKey string
	client *http.Client
}

func NewIssueTracker(apiKey string) IssueTracker {
	return &CodesiteIssueTracker{
		apiKey: apiKey,
		client: util.NewTimeoutClient(),
	}
}

// Issue is an individual issue returned from the project hosting response.
type Issue struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}

// IssueResponse is used to decode JSON responses from the project hosting API.
type IssueResponse struct {
	Items []Issue `json:"items"`
}

// FromQuery is part of the IssueTracker interface. See documentation there.
func (c *CodesiteIssueTracker) FromQuery(q string) ([]Issue, error) {
	qStr := fmt.Sprintf(API_QUERY_TEMPLATE, q, c.apiKey)
	url := fmt.Sprintf(API_URL_TEMPLATE, url.QueryEscape(qStr))

	//  This will return a JSON response of the form:
	//
	//  {
	//   "items": [
	//    {
	//     "id": 2874,
	//     "title": "this is a bug with..."
	//     "state": "open"
	//    }
	//   ]
	//  }
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer util.Close(resp.Body)

	issueResponse := &IssueResponse{
		Items: []Issue{},
	}
	if err := json.NewDecoder(resp.Body).Decode(&issueResponse); err != nil {
		return nil, err
	}

	return issueResponse.Items, err
}
