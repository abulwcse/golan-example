package app

import (
	"github.com/abulwcse/golan-example/pkg/book"
)

var handlerMap = map[string]RouteHandler{
	"BookIndexController": book.BookIndexController{},
}
