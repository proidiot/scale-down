package main

import (
	"git.stuph.net/proidiot/scale-down"
	"fmt"
	"github.com/fitstar/falcore"
)

func main() {
	pipeline := falcore.NewPipeline()

	pipeline.Upstream.PushBack(scaledown.NewLandingFilter())

	server := falcore.NewServer(8080, pipeline)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Could not start server", err)
	}
}

