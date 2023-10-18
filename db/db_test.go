package db_test

import (
	"testing"

	"github.com/DaveSaah/voting-system/db"
)

func TestDatabaseConnection(t *testing.T) {
	conn, err := db.Init()
	if err != nil {
		t.Fatalf("Cannot create the db connection: %s", err)
	}

	defer conn.Close()

	if err := conn.Ping(); err != nil {
		t.Fatalf("Ping Error: %s", err)
	}

	t.Log("Database connected!")
}
