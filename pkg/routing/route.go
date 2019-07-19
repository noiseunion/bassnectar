package routing

import "net/http"

// Route that we want to configure for our server
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
