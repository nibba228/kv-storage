package main

import (
  "context"
  "database/sql"

  pb "github.com/nibba228/kv-storage/storage"

  _ "github.com/go-sql-driver/mysql"
)

type StorageManager struct {
  pb.UnimplementedStorageServer
  db *sql.DB
}


func (sm *StorageManager) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
  rows, err := sm.db.QueryContext(ctx, "SELECT FROM storage WHERE user_id = ? AND key = ?", interface{}{}, key.key) 

}
