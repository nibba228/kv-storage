package main

import (
	"context"
	"database/sql"
  "log"

	_ "github.com/go-sql-driver/mysql"
)

func Get(ctx context.Context, db *sql.DB, user_id, key string) (string, error) {
  query := "select value from kv where user_id=? and key_=?"
 
  var value string
  err := db.QueryRowContext(ctx, query, user_id, key).Scan(&value)

  if err == sql.ErrNoRows {
    log.Printf("No value for user %s and key %s: %v", user_id, key, err)
    return "", err
  }

  if err != nil {
    log.Printf("Unknown error in db.Get: %v", err)
    return "", err
  }

  return value, nil
}
