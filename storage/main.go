package main

import (
	"context"
	"database/sql"
  "log"
  "net"
  "os"

	pb "github.com/nibba228/kv-storage/proto/storage"

	"google.golang.org/grpc"
	"github.com/go-sql-driver/mysql"
)

type StorageManager struct {
  pb.UnimplementedStorageServer
  db *sql.DB
}

func newServer() *StorageManager {
  conf := mysql.Config{
    User: os.Getenv("STORAGE_DB_USER"),
    Passwd: os.Getenv("STORAGE_DB_PASSWORD"),
    Net: "tcp",
    Addr: os.Getenv("STORAGE_DB_HOST"),
    DBName: os.Getenv("STORAGE_DB"),
  }

  log.Printf("dsn connection string: %s", conf.FormatDSN())

  db, err := sql.Open("mysql", conf.FormatDSN())
  if err != nil {
    log.Fatal(err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  return &StorageManager{
    db: db,
  }
}

func (sm *StorageManager) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
  user_id := ctx.Value("user_id").(string)
  value, err := Get(ctx, sm.db, user_id, key.Key)
  return &pb.Value{Value: value}, err
}

func main() {
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  pb.RegisterStorageServer(grpcServer, newServer())

  lis, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
