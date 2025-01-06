package main

import (
  "context"
  "database/sql"

  pb "github.com/nibba228/kv-storage/storage"

  _ "github.com/go-sql-driver/mysql"
)

type StorageManager struct {
  pb.UnimplementedStorageServer
}

func (sm *StorageManager) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {

}
