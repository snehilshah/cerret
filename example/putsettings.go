package main

import (
	"context"
	"errors"
	"math/rand"
)

type person struct {
	Name string
}

func (p person) PutSettings(ctx context.Context, in string) (string, error) {
	if rand.Int()%2 == 0 {
		return "false", errors.New("Even Number")
	}
	err := errors.New("Some Internal Error")
	return "ok", err
}
