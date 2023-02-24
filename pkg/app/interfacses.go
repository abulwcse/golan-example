package app

import (
	"net/http"
)

type Model interface {
	GatName() string
}

type RouteHandler interface {
	http.Handler
}
