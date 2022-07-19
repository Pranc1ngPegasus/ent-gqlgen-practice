package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Pranc1ngPegasus/ent-gqlgen-practice/adapter/resolver"
	"github.com/Pranc1ngPegasus/ent-gqlgen-practice/ent"
	"github.com/google/uuid"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
)

const (
	port string = "3000"
)

func main() {
	ctx := context.Background()

	log.Print("open DB connection")
	client, err := ent.Open("sqlite3", "file:sqlite.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	log.Print("migrate DB")
	if err := client.Debug().Schema.Create(ctx); err != nil {
		panic(err)
	}

	if _, err := client.Debug().Account.Create().SetID(uuid.New().String()).SetEmail("foobar@example.com").Save(ctx); err != nil {
		panic(err)
	}

	log.Print("initialize server")
	srv := handler.NewDefaultServer(
		resolver.NewSchema(client),
	)
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("Practice", "/query"))
	http.Handle("/query", srv)

	log.Print("start server")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
