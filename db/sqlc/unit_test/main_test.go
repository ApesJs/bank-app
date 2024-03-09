package unit_test

import (
	"context"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	//dbDriver = "postgres"
	dbSource = "postgresql://root:apesjs123@localhost:5432/app-bank?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal("cannot close connection to DB:", err)
		}
	}(conn, context.Background())

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
