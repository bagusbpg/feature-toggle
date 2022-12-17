package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unleash/unleash-client-go/v3"
)

func main() {
	featureToggle, err := unleash.NewClient(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithRefreshInterval(2*time.Second),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"default:development.unleash-insecure-api-token"}}),
	)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		featureEnabled := featureToggle.IsEnabled("toggle")
		if featureEnabled {
			fmt.Fprintf(w, "feature is enabled")
		} else {
			fmt.Fprintf(w, "feature is disabled")
		}
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
