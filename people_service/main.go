package main

import (
	"fmt"
	"people_service/service/store"
)

func main() {
	s := store.NewStore("postgres://postgres:databaza1@localhost:5432/people_service")

	people, _ := s.ListPeople()
	fmt.Println(people)

	person, _ := s.GetPeopleByID(1)

	fmt.Println(person)
}
