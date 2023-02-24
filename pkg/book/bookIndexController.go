package book

import (
	"encoding/json"
	"fmt"
	"github.com/abulwcse/golan-example/pkg/db"
	"net/http"
)

type BookIndexController struct {
}

func (b BookIndexController) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	repo := BookNeo4jRepository{
		Driver: db.GetDB(),
	}
	book, _ := repo.GetBook(1)
	response, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(response)
}
