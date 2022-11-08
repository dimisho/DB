package service

import (
	"fmt"
	"people_service/service/store"
)

type storer interface {
	ListPeople() ([]store.People, error)
	GetPeopleByID(id int) (store.People, error)
}

type tax interface {
	GetTaxStatusByID(id int) (string, error)
}

type Service struct {
	Store storer
	Tax   tax
}

type PeopleWithTax struct {
	store.People
	TaxStatus string
}

func (s *Service) ListPeople() ([]PeopleWithTax, error) {
	list, err := s.Store.ListPeople()
	if err != nil {
		return nil, fmt.Errorf("list people: %w", err)
	}

	people := make([]PeopleWithTax, 0, len(list))
	for _, l := range list {
		st, err := s.Tax.GetTaxStatusByID(l.ID)
		if err != nil {
			return nil, fmt.Errorf("get tax status: %w", err)
		}

		people = append(people, PeopleWithTax{
			People:    l,
			TaxStatus: st,
		})
	}

	return people, nil
}

func (s *Service) GetPeopleByID(id int) ([]PeopleWithTax, error) {
	p, err := s.Store.GetPeopleByID(id)
	if err != nil {
		return nil, fmt.Errorf("get people by id: %w", err)
	}

	st, err := s.Tax.GetTaxStatusByID(p.ID)
	if err != nil {
		return nil, fmt.Errorf("get tax status: %w", err)
	}

	return []PeopleWithTax{{
		People:    p,
		TaxStatus: st,
	}}, nil
}
