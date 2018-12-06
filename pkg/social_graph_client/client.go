package social_graph_client

import (
	"bytes"
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	om "github.com/the-gigi/delinkcious/pkg/object_model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func NewClient(baseURL string) (om.SocialGraphManager, error) {
	// Quickly sanitize the instance string.
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	followEndpoint := httptransport.NewClient(
			"POST",
			copyURL(u, "/follow"),
			encodeHTTPGenericRequest,
		    decodeSimpleResponse,
			nil).Endpoint()

	unfollowEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/unfollow"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse,
		nil).Endpoint()

	getFollowingEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/following"),
		encodeHTTPGenericRequest,
		decodeGetFollowingResponse,
		nil).Endpoint()

	getFollowersEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/followers"),
		encodeHTTPGenericRequest,
		decodeGetFollowersResponse,
		nil).Endpoint()


	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return EndpointSet{
		FollowEndpoint: followEndpoint,
		UnfollowEndpoint: unfollowEndpoint,
		GetFollowingEndpoint: getFollowingEndpoint,
		GetFollowersEndpoint: getFollowersEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}


// encodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}