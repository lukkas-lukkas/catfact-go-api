package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewGetFactService(url string) *CatFactService {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(cxt context.Context) (*CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close();

	catFact := &CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(catFact); err != nil {
		return nil, err
	}

	return catFact, nil
}
