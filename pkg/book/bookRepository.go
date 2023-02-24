package book

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Book struct {
	ID       int64
	Title    string
	ISBN     string
	Language string
	Author   string
}

type BookRepository interface {
	GetBook(id int) Book
	GetBooks(ids []int)
}

type BookNeo4jRepository struct {
	Driver neo4j.DriverWithContext
}

func (b BookNeo4jRepository) GetBook(id int) (Book, error) {
	books, err := b.GetBooks(id)
	fmt.Println(books)
	if err != nil {
		fmt.Println(err)
		return Book{}, err
	}
	return books[1], nil
}

func (b BookNeo4jRepository) GetBooks(id int) ([]Book, error) {
	ctx := context.Background()
	session := b.Driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	var books, err = session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(
			ctx,
			"MATCH(b:Book {ID:$id}) RETURN b.ID as id, b.name as name, b.isbn as isbn, b.language as language", map[string]any{
				"id": id,
			})

		if err != nil {
			fmt.Println(err)
		}

		response := make([]Book, 1)
		var record *neo4j.Record
		for result.NextRecord(ctx, &record) {
			id, _ := record.Get("id")
			if id.(int64) <= 0 {
				continue
			}
			name, _ := record.Get("name")
			isbn, _ := record.Get("isbn")
			language, _ := record.Get("language")

			response = append(response, Book{
				ID:       id.(int64),
				ISBN:     isbn.(string),
				Title:    name.(string),
				Language: language.(string),
			})
		}

		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return books.([]Book), nil

}
