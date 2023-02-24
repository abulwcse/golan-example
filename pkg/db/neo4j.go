package db

import (
	"github.com/abulwcse/golan-example/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetDB() neo4j.DriverWithContext {
	dbUri := config.Neo4jUrl
	instance, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		panic(err)
	}
	return instance
}
