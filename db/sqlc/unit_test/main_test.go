package unit_test

import (
	"context"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
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
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer testDB.Close()

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
