package ent

import (
    "context"
    "log"

    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
    _ "github.com/lib/pq"
)

func NewClient(dsn string) *Client {
    drv, err := sql.Open(dialect.Postgres, dsn)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    return NewClient(Driver(drv))
}

func Migrate(client *Client) {
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
