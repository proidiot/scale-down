package scaledown

import (
	"github.com/fitstar/falcore"
)

type Authenticator interface {
	Authenticating(req *falcore.Request) bool
	AuthenticationChallenge() falcore.RequestFilter
	Authenticate(req *falcore.Request) bool
}

