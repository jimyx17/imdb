package imdb

import (
	"net/http"
	"time"
)

// client is used by tests to perform cached requests.
// If cache directory exists it is used as a persistent cache.
// Otherwise a volatile memory cache is used.
var client *http.Client

var ttl = 24 * time.Hour

func init() {

	client = http.DefaultClient
}
