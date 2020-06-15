package db

import (
	"fmt"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type Database struct {
	session sqlbuilder.Database
}

func NewDatabaseConnection(dbURL string) (Database, error) {
	session, err := connectToDatabase(dbURL)
	if err != nil {
		return Database{}, err
	}

	return Database{session: session}, nil
}

func connectToDatabase(dbURL string) (sqlbuilder.Database, error) {
	url, err := mysql.ParseURL(dbURL)
	if err != nil {
		return nil, err
	}

	session, err := mysql.Open(url)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (d Database) InsertEvent(event Event) {
	_, err := d.session.
		InsertInto("events").
		Columns("id", "entity", "payload").
		Values(event.ID, event.Entity, event.Payload).
		Exec()

	if err != nil {
		fmt.Println("Failed to insert event to database", err)
	}
}
