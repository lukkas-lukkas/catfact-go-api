package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) *LoggingService {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(cxt context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%v err=%v took=%v", fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(cxt)
}
