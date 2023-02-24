package app

import (
	"github.com/abulwcse/golan-example/config"
	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

type Routes struct {
	Path    string
	Method  string
	Handler string
}

func (r Routes) GatName() string {
	return r.Path
}

type App struct {
	Router *mux.Router
}

func (app *App) Initialise() {
	app.Router = mux.NewRouter()
	app.addRoutes()
	err := http.ListenAndServe(config.URL+":"+config.Port, app.Router)
	if err != nil {
		panic(err)
	}
}

func (app *App) checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func (app *App) addRoutes() {
	currentPath, _ := os.Getwd()
	configFile := filepath.Join(currentPath, "/config/routes.yml")
	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_CREATE|os.O_RDONLY, 0600)
	app.checkErrors(err)
	decoder := yaml.NewDecoder(f)
	data := make(map[string]Routes)
	err = decoder.Decode(data)
	app.checkErrors(err)
	var handler RouteHandler
	for _, v := range data {
		handler = handlerMap[v.Handler]
		app.Router.Handle(v.Path, handler).Methods(v.Method)
	}

}
