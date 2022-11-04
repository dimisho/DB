package store

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	conn *sqlx.DB
}

type People struct {
	ID   int
	Name string
}

func NewStore(connString string) *Store {
	conn, err := sqlx.Connect("pgx", connString)

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	people := make([]People, 0)
	var id int
	var name string

	rows, err := s.conn.Query("SELECT * FROM people")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		people = append(people, People{
			ID:   id,
			Name: name,
		})
	}

	return people, nil
}

func (s *Store) GetPeopleByID(id int) (People, error) {
	var name string
	var stringId = strconv.Itoa(id)

	row := s.conn.QueryRow("SELECT * FROM people WHERE id = " + stringId)

	err := row.Scan(&id, &name)
	if err != nil {
		return People{}, err
	}

	return People{ID: id, Name: name}, nil
}
