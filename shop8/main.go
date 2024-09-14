package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Store struct {
	name  []string
	price []float64
	count []int
}

func NewStore() *Store {
	return &Store{
		name:  []string{},
		price: []float64{},
		count: []int{},
	}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	og := name
	name = strings.ToLower(name)
	for _, v := range s.name {
		v = strings.ToLower(v)
		if v == name {
			return fmt.Errorf("%s already exists", og)
		}
	}
	if price <= 0 {
		return errors.New("price should be positive")
	}
	if count <= 0 {
		return errors.New("count should be positive")
	}
	s.name = append(s.name, name)
	s.count = append(s.count, count)
	s.price = append(s.price, price)
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	name = strings.ToLower(name)
	for i := 0; i < len(s.name); i++ {
		if s.name[i] == name {
			return s.count[i], nil
		}
	}
	return 0, errors.New("invalid product name")

}

func (s *Store) GetProductPrice(name string) (float64, error) {
	name = strings.ToLower(name)
	for i := 0; i < len(s.name); i++ {
		if s.name[i] == name {
			return s.price[i], nil
		}
	}
	return 0, errors.New("invalid product name")

}

func (s *Store) Order(name string, count int) error {
	name = strings.ToLower(name)
	if count <= 0 {
		return errors.New("count should be positive")
	}
	var selected int
	var b bool
	for i := 0; i < len(s.name); i++ {
		if s.name[i] != name {
			b = false
		} else {
			b = true
			selected = i
			break
		}
	}
	if !b {
		return errors.New("invalid product name")
	} else if b && s.count[selected] == 0 {
		return fmt.Errorf("there is no %s in the store", s.name[selected])
	} else if b && count > s.count[selected] {
		return fmt.Errorf("not enough %s in the store. there are %d left", s.name[selected], s.count[selected])
	} else {
		s.count[selected] -= count
		return nil
	}

}

func (s *Store) ProductsList() ([]string, error) {

	if len(s.name) == 0 {
		return nil, errors.New("store is empty")
	}
	var b bool
	for _, v := range s.count {
		if v > 0 {
			b = true
			break
		} else {
			b = false
		}
	}
	if !b {
		return nil, errors.New("store is empty")
	}
	nz := make([]string, 0)
	for i := 0; i < len(s.count); i++ {
		if s.count[i] > 0 {
			nz = append(nz, s.name[i])
		}
	}
	sort.Strings(nz)
	return nz, nil
}
